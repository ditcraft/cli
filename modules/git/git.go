package git

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// GetRepository will return the name of the current repostiry
// minus the "https://github.com/" part
func GetRepository() (string, error) {
	// Retrieving the current repository url from git
	cmdName := "git"
	cmdArgs := []string{"config", "--get", "remote.origin.url"}
	cmdOut, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if err != nil {
		return "", errors.New("There was an error running git config")
	}

	cmdOutString := string(cmdOut)[0 : len(string(cmdOut))-1]
	if !strings.Contains(cmdOutString, "github.com") {
		return "", errors.New("Invalid repository - currently only GitHub repositories are supported")
	}

	// Removing unecessary stuff from the url
	cmdOutString = strings.Replace(cmdOutString, "http://", "", -1)
	cmdOutString = strings.Replace(cmdOutString, "https://", "", -1)
	cmdOutString = strings.Replace(cmdOutString, "github.com/", "", -1)
	// Removing gitlab aswell for the check below
	cmdOutString = strings.Replace(cmdOutString, "gitlab.com/", "", -1)

	configErr := checkGitSetup()
	if configErr != nil {
		fmt.Println("******* WARNING *******")
		fmt.Println("You don't have a user.email and/or user.name set in your git config. dit won't work if this is not fixed!")
		fmt.Println("Please run 'git config --global user.name <GITHUB_USERNAME_HERE>' and 'git config --global user.email <GITHUB_EMAIL_HERE>'")
		fmt.Println("******* WARNING *******")
		return "", errors.New("git is not configured correctly")
	}

	return cmdOutString, nil
}

// Clone wraps the git clone functionality and initialized the ditContract afterwards
func Clone(_repository string) (string, error) {
	// Only GitHub repositores are supported right now
	if !strings.Contains(_repository, "github.com") {
		return "", errors.New("Currently only github repositories are supported")
	}

	// Calling git clone
	cmdArgs := []string{"clone", _repository, "--progress"}
	cmdOut, err := exec.Command("git", cmdArgs...).CombinedOutput()
	if strings.Contains(string(cmdOut), "Invalid username or password") {
		fmt.Println("Wrong username or password for GitHub")
		fmt.Println("Note: If you have 2FA activated for GitHub, please go to: ")
		fmt.Println("https://github.blog/2013-09-03-two-factor-authentication/#how-does-it-work-for-command-line-git")
	}
	if err != nil && !strings.Contains(string(cmdOut), "already exists") {
		return "", errors.New("There was an error running git clone")
	}

	// Removing unecessary stuff from the url
	_repository = strings.Replace(_repository, "http://", "", -1)
	_repository = strings.Replace(_repository, "https://", "", -1)
	_repository = strings.Replace(_repository, "github.com/", "", -1)

	// Only for testing purposes, will be removed
	_repository = strings.Replace(_repository, "gitlab.com/", "", -1)

	configErr := checkGitSetup()
	if configErr != nil {
		fmt.Println("******* WARNING *******")
		fmt.Println("You don't have a user.email and/or user.name set in your git config. dit won't work if this is not fixed!")
		fmt.Println("Please run 'git config user.name <GITHUB_USERNAME_HERE>' and 'git config user.email <GITHUB_EMAIL_HERE>'")
		fmt.Println("******* WARNING *******")
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
		// Create an empty README.md file
		cmdName = "touch"
		cmdArgs = []string{"README.md"}
		cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput()
		if err != nil {
			return errors.New("There was an error creating the README.md")
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
			fmt.Println("Wrong username or password for GitHub")
			fmt.Println("Note: If you have 2FA activated for GitHub, please go to: ")
			fmt.Println("https://github.blog/2013-09-03-two-factor-authentication/#how-does-it-work-for-command-line-git")
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
		fmt.Println(string(cmdOut))
		fmt.Println(err)
		return errors.New("There was an error running git commit")
	}

	// Push these changes to the new branch
	cmdName = "git"
	cmdArgs = []string{"push", "--set-upstream", "origin", branchName, "-v"}
	cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput()
	if strings.Contains(string(cmdOut), "Invalid username or password") {
		fmt.Println("Wrong username or password for GitHub")
		fmt.Println("Note: If you have 2FA activated for GitHub, please go to: ")
		fmt.Println("https://github.blog/2013-09-03-two-factor-authentication/#how-does-it-work-for-command-line-git")
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
		fmt.Println("Wrong username or password for GitHub")
		fmt.Println("Note: If you have 2FA activated for GitHub, please go to: ")
		fmt.Println("https://github.blog/2013-09-03-two-factor-authentication/#how-does-it-work-for-command-line-git")
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
		fmt.Println("Wrong username or password for GitHub")
		fmt.Println("Note: If you have 2FA activated for GitHub, please go to: ")
		fmt.Println("https://github.blog/2013-09-03-two-factor-authentication/#how-does-it-work-for-command-line-git")
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
