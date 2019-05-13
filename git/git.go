package git

import (
	"errors"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"

	"github.com/ditcraft/cli/helpers"
)

// GetRepository will return the name of the current repostiry
// minus the "https://github.com/" part
func GetRepository() (string, error) {
	// Retrieving the current repository url from git
	cmdName := "git"
	cmdArgs := []string{"config", "--get", "remote.origin.url"}
	cmdOut, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if err != nil {
		return "", errors.New("There was an error running git config - are you in a repository folder?")
	}

	cmdOutString := string(cmdOut)[0 : len(string(cmdOut))-1]

	// Removing unecessary stuff from the url
	cmdOutString = sanitizeURL(cmdOutString)

	configErr := checkGitSetup()
	if configErr != nil {
		helpers.PrintLine("You don't have a user.email and/or user.name set in your git config. dit won't work if this is not fixed!", 2)
		helpers.PrintLine("Please run 'git config user.name <GITHUB_USERNAME_HERE>' and 'git config user.email <GITHUB_EMAIL_HERE>'", 2)
		return "", errors.New("git is not configured correctly")
	}

	return cmdOutString, nil
}

// Clone wraps the git clone functionality and initialized the ditContract afterwards
func Clone(_repository string) (string, error) {
	// Only GitHub repositores are supported right now
	// if !strings.Contains(_repository, "github.com") {
	// 	return "", errors.New("Currently only github repositories are supported")
	// }

	// Calling git clone
	cmdArgs := []string{"clone", _repository, "--progress"}
	cmdOut, err := exec.Command("git", cmdArgs...).CombinedOutput()
	if strings.Contains(string(cmdOut), "Repository not found") {
		return "", errors.New("Repository not found")
	}
	if strings.Contains(string(cmdOut), "unable to update url base from redirection") {
		return "", errors.New("Repository URL invalid or not found")
	}
	if strings.Contains(string(cmdOut), "already exists and is not an empty directory.") {
		return "", errors.New("Destination path already exists")
	}
	if strings.Contains(string(cmdOut), "Invalid username or password") {
		helpers.PrintLine("Wrong username or password", 2)
		helpers.PrintLine("Note: If you have 2FA activated for GitHub, please go to: ", 1)
		helpers.PrintLine("https://github.blog/2013-09-03-two-factor-authentication/#how-does-it-work-for-command-line-git", 1)
		return "", errors.New("")
	}
	if err != nil {
		return "", errors.New("There was an error running git clone")
	}

	// Removing unecessary stuff from the url
	_repository = sanitizeURL(_repository)

	configErr := checkGitSetup()
	if configErr != nil {
		helpers.PrintLine("You don't have a user.email and/or user.name set in your git config. dit won't work if this is not fixed!", 1)
		helpers.PrintLine("Please run 'git config user.name <GITHUB_USERNAME_HERE>' and 'git config user.email <GITHUB_EMAIL_HERE>'", 1)
	}

	return _repository, nil
}

func checkGitSetup() error {
	// Verifying whether a git user email is set correctly
	cmdName := "git"
	cmdArgs := []string{"config", "user.email"}
	cmdOut, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if len(string(cmdOut)) <= 0 || err != nil {
		return errors.New("There is no git user.email set")
	}

	// Verifying whether a git user name is set correctly
	cmdName = "git"
	cmdArgs = []string{"config", "user.name"}
	cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if len(string(cmdOut)) <= 0 || err != nil {
		return errors.New("There is no git user.name set")
	}

	// Verifying whether the git password is being cached
	// If nothing is set, we will cache it for now
	cmdName = "git"
	cmdArgs = []string{"config", "credential.helper"}
	cmdOut, _ = exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if len(string(cmdOut)) <= 0 {
		cmdName = "git"
		cmdArgs = []string{"config", "credential.helper", "cache"}
		_, _ = exec.Command(cmdName, cmdArgs...).CombinedOutput()
	}

	return nil
}

// Validate will return an error when the connection to
// the git provider fails (validated through a git fetch)
func Validate() error {
	repoName, err := GetRepository()
	if err != nil {
		return err
	}
	repoName = strings.SplitN(repoName, "/", 2)[1]

	// Verifying whether a master branch (or any branch) exists
	cmdName := "git"
	cmdArgs := []string{"branch"}
	cmdOutBranch, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if err != nil {
		return errors.New("There was an error running git branch")
	}

	// Verifying whether the connection works correctly
	cmdName = "git"
	cmdArgs = []string{"fetch", "-v"}
	cmdOutFetch, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if (!strings.Contains(string(cmdOutFetch), repoName) && len(cmdOutBranch) > 0) || err != nil {
		return errors.New("The connection to the git provider failed - please ensure that the repository is configured correctly")
	}

	return nil
}

// CheckForChanges will return an error if there are no new files to commit
func CheckForChanges() error {
	// Verifying whether files have been added
	cmdName := "git"
	cmdArgs := []string{"ls-files", "--others", "--modified", "--exclude-standard"}
	cmdOut, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if err != nil {
		return errors.New("There was an error running the git ls-files command")
	}

	// If there are no files added, nothing can be committed
	if len(cmdOut) == 0 {
		return errors.New("There are no changed files to add for a commit")
	}

	return nil
}

// Commit will push the changes onto a new proposal branch
func Commit(_proposalID int, _commitMessage string) error {
	// Name of the branch that will be used
	branchName := "dit_proposal_" + strconv.Itoa(_proposalID)

	// Verifying whether a master branch (or any branch) exists
	cmdName := "git"
	cmdArgs := []string{"branch"}
	cmdOut, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if err != nil {
		return errors.New("There was an error running git branch")
	}

	// If not: an empty readme will be pushed to master in order to
	// prevent the dit_proposal branch to be the main branch through the first commit
	if len(cmdOut) == 0 {
		repositoryName, err := GetRepository()
		if err != nil {
			return err
		}

		// Create a README.md file
		err = ioutil.WriteFile("README.md", []byte("# "+repositoryName), 0644)
		if err != nil {
			return errors.New("Failed to write README.md file")
		}

		// Add it for a commit
		cmdName = "git"
		cmdArgs = []string{"add", "README.md"}
		cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput()
		if err != nil {
			return errors.New("There was an error adding the README.md")
		}

		// Commit it
		cmdName = "git"
		cmdArgs = []string{"commit", "-m", "Initial push to master"}
		cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput()
		if err != nil {
			return errors.New("There was an error commiting the initial README")
		}

		// Push it
		cmdName = "git"
		cmdArgs = []string{"push", "-v"}
		cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput()
		if strings.Contains(string(cmdOut), "Invalid username or password") {
			helpers.PrintLine("Wrong username or password", 2)
			helpers.PrintLine("Note: If you have 2FA activated for GitHub, please go to: ", 2)
			helpers.PrintLine("https://github.blog/2013-09-03-two-factor-authentication/#how-does-it-work-for-command-line-git", 2)
		}
		if err != nil {
			return errors.New("There was an error pushing the initial commit")
		}
	}

	// Switch to the dit_proposal branch
	cmdName = "git"
	cmdArgs = []string{"checkout", "-b", branchName}
	cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if err != nil {
		return errors.New("There was an error running git checkout")
	}

	// Verify whether we are on the dit_proposal branch
	cmdName = "git"
	cmdArgs = []string{"status"}
	cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if err != nil || !strings.Contains(string(cmdOut), "On branch "+branchName) {
		return errors.New("There was an error running git status")
	}

	// Adding all files that can be added
	cmdName = "git"
	cmdArgs = []string{"add", "."}
	cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if err != nil {
		return errors.New("There was an error running the git add command")
	}

	// Verifying whether the files have been added
	cmdName = "git"
	cmdArgs = []string{"diff", "--cached", "--name-only"}
	cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if err != nil {
		return errors.New("There was an error running the git diff command")
	}

	// If there are no files added, nothing can be committed
	if len(cmdOut) == 0 {
		return errors.New("There are no changed files to add for a commit")
	}

	// Commit the changes with the provided commit message
	cmdName = "git"
	cmdArgs = []string{"commit", "-m", _commitMessage, "-v"}
	cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if err != nil || !strings.Contains(string(cmdOut), _commitMessage) {
		return errors.New("There was an error running git commit")
	}

	// Push these changes to the new branch
	cmdName = "git"
	cmdArgs = []string{"push", "--set-upstream", "origin", branchName, "-v"}
	cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if strings.Contains(string(cmdOut), "Invalid username or password") {
		helpers.PrintLine("Wrong username or password", 2)
		helpers.PrintLine("Note: If you have 2FA activated for GitHub, please go to: ", 2)
		helpers.PrintLine("https://github.blog/2013-09-03-two-factor-authentication/#how-does-it-work-for-command-line-git", 2)
	}
	if err != nil || !strings.Contains(string(cmdOut), "set up to track remote branch") {
		return errors.New("There was an error running git push")
	}

	// Switch to the master branch
	cmdName = "git"
	cmdArgs = []string{"checkout", "master"}
	cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if err != nil {
		return errors.New("There was an error running the last git checkout")
	}

	// Verify whether we are on the master branch again
	cmdName = "git"
	cmdArgs = []string{"status"}
	cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if err != nil || !strings.Contains(string(cmdOut), "On branch master") {
		return errors.New("There was an error running the last git status")
	}

	return nil
}

// Merge will merge a dit_proposal into the master after a successfull vote
func Merge(_proposalID string) error {
	// Name of the branch that will be used
	branchName := "dit_proposal_" + _proposalID

	// Switching to the master branch
	cmdName := "git"
	cmdArgs := []string{"checkout", "master"}
	_, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if err != nil {
		return errors.New("There was an error running git checkout master")
	}

	// Merging the dit_proposal branch into the master
	cmdName = "git"
	cmdArgs = []string{"merge", branchName, "-v"}
	_, err = exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if err != nil {
		return errors.New("There was an error running git merge")
	}

	// Pushing the merged master branch
	cmdName = "git"
	cmdArgs = []string{"push", "-v"}
	cmdOut, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if strings.Contains(string(cmdOut), "Invalid username or password") {
		helpers.PrintLine("Wrong username or password", 2)
		helpers.PrintLine("Note: If you have 2FA activated for GitHub, please go to: ", 2)
		helpers.PrintLine("https://github.blog/2013-09-03-two-factor-authentication/#how-does-it-work-for-command-line-git", 2)
	}
	if err != nil {
		return errors.New("There was an error running git push")
	}

	// Deleting the now obsolete dit_proposal branch
	err = DeleteBranch(_proposalID)
	if err != nil {
		return errors.New("There was an error deleting the branch after a successful merge")
	}

	return nil
}

// DeleteBranch will delete a dit_proposal branch after it has been merged or after an unsuccessful vote
func DeleteBranch(_proposalID string) error {
	// Name of the branch that will be used
	branchName := "dit_proposal_" + _proposalID

	// Remove the branch on the remote side
	cmdName := "git"
	cmdArgs := []string{"push", "--delete", "origin", branchName}
	cmdOut, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if strings.Contains(string(cmdOut), "Invalid username or password") {
		helpers.PrintLine("Wrong username or password", 2)
		helpers.PrintLine("Note: If you have 2FA activated for GitHub, please go to: ", 2)
		helpers.PrintLine("https://github.blog/2013-09-03-two-factor-authentication/#how-does-it-work-for-command-line-git", 2)
	}
	if err != nil || !strings.Contains(string(cmdOut), "deleted") {
		return errors.New("There was an error deleting the branch remotely")
	}

	// Remove the branch on the local side
	cmdName = "git"
	cmdArgs = []string{"branch", "-D", branchName}
	cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if err != nil || !strings.Contains(string(cmdOut), "Deleted") {
		return errors.New("There was an error deleting the branch locally")
	}

	return nil
}

func sanitizeURL(_url string) string {
	_url = strings.Replace(_url, "http://", "", -1)
	_url = strings.Replace(_url, "https://", "", -1)
	_url = strings.Replace(_url, "git@", "", -1)
	_url = strings.Replace(_url, "github.com:", "github.com/", -1)
	_url = strings.Replace(_url, "gitlab.com:", "gitlab.com/", -1)
	_url = strings.Replace(_url, "bitbucket.org:", "bitbucket.com/", -1)
	_url = strings.Replace(_url, ".git", "", -1)

	return _url
}
