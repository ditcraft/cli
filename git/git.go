package git

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/ditcraft/cli/helpers"
)

var wg sync.WaitGroup

// GetRepository will return the name of the current repostiry
// minus the "https://github.com/" part
func GetRepository() (string, error) {
	// Retrieving the current repository url from git
	gitOutput, err := ExecuteCommandWithoutStdOut("config", "--get", "remote.origin.url")
	if err != nil {
		return "", errors.New("There was an error retrieving the git repositories' name. Is the repository set up correctly?")
	}

	repositoryRemote := strings.TrimSpace(gitOutput)

	// Removing unecessary stuff from the url
	repositoryName := sanitizeURL(repositoryRemote)

	return repositoryName, nil
}

// Clone wraps the git clone functionality and initialized the ditContract afterwards
func Clone(_repository string) (string, error) {
	// Calling git clone
	err := ExecuteCommand("clone", _repository, "--progress")
	if err != nil {
		return "", err
	}

	// if strings.Contains(string(cmdOut), "Invalid username or password") {
	// 	helpers.PrintLine("Wrong username or password", helpers.ERROR)
	// 	helpers.PrintLine("Note: If you have 2FA activated for GitHub, please go to: ", helpers.WARN)
	// 	helpers.PrintLine("https://github.blog/2013-09-03-two-factor-authentication/#how-does-it-work-for-command-line-git", helpers.WARN)
	// 	return "", errors.New("")
	// }
	// if err != nil {
	// 	return "", errors.New("There was an error running git clone")
	// }

	// Removing unecessary stuff from the url
	_repository = sanitizeURL(_repository)

	return _repository, nil
}

// GetHeadHashOfBranch will return the hash of a branches' head commit
func GetHeadHashOfBranch(_branch string) (string, error) {
	// Calling git clone
	gitOutput, err := ExecuteCommandWithoutStdOut("rev-parse", _branch)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(gitOutput), nil
}

// IsOnMaster will return a bool, indiciating whether the user is on the master branch
func IsOnMaster() (bool, error) {
	// Retrieving the current branch name
	gitOutput, err := ExecuteCommandWithoutStdOut("symbolic-ref", "--short", "HEAD")
	if err != nil {
		return false, errors.New("There was an error retrieving the current branch name")
	}
	if strings.TrimSpace(gitOutput) == "master" {
		return true, nil
	}
	return false, nil
}

// MasterIsClean will throw an error when the master branch is different from the remote master branch
// which will tell the merge command that there are changes that the users might try to cover up
func MasterIsClean() error {
	// Fetching the master branch
	err := ExecuteCommand("fetch")
	if err != nil {
		return errors.New("There was an error retrieving the current status of the remote master")
	}

	// Retrieving the head of the local master
	gitOutputLocal, err := ExecuteCommandWithoutStdOut("rev-parse", "@")
	if err != nil {
		return errors.New("There was an error retrieving the current head of the local master")
	}
	gitOutputLocal = strings.TrimSpace(gitOutputLocal)

	// Retrieving the head of the remote master
	gitOutputRemote, err := ExecuteCommandWithoutStdOut("rev-parse", "origin/master")
	if err != nil {
		return errors.New("There was an error retrieving the current head of the remote master")
	}
	gitOutputRemote = strings.TrimSpace(gitOutputRemote)

	// Retrieving the head of the merge base
	gitOutputBase, err := ExecuteCommandWithoutStdOut("merge-base", "@", "origin/master")
	if err != nil {
		return errors.New("There was an error retrieving the current head of the remote merge base")
	}
	gitOutputBase = strings.TrimSpace(gitOutputBase)

	if gitOutputLocal == gitOutputRemote {
		return nil
	} else if gitOutputLocal == gitOutputBase {
		return errors.New("You are behind the remote master branch - please pull first")
	} else if gitOutputRemote == gitOutputBase {
		return errors.New("You are ahead of the remote master branch - please push or resolve this first")
	}

	return errors.New("Your remote and local master branches have diverged - please resolve this first")
}

// SideBranchIsPushed will throw an error when the local side branch is different from the remote side branch
// which will tell the merge command that the user hasn't pushed his changes live
func SideBranchIsPushed(_branch string) error {
	// Fetching the branches
	err := ExecuteCommand("fetch")
	if err != nil {
		return errors.New("There was an error retrieving the current status of the remote branches")
	}

	// Retrieving the head of the local sidebranch
	gitOutputLocal, err := ExecuteCommandWithoutStdOut("rev-parse", _branch)
	if err != nil {
		return errors.New("There was an error retrieving the current head of the local sidebranch")
	}
	gitOutputLocal = strings.TrimSpace(gitOutputLocal)

	// Retrieving the head of the remote sidebranch
	gitOutputRemote, err := ExecuteCommandWithoutStdOut("rev-parse", "origin/"+_branch)
	if err != nil {
		helpers.PrintLine("There was an error retrieving the current head of the remote sidebranch", helpers.ERROR)
		helpers.PrintLine("This usually happens if you never pushed the commits of the sidebranch", helpers.ERROR)
		return errors.New("Please push before trying to merge into master")
	}
	gitOutputRemote = strings.TrimSpace(gitOutputRemote)

	// Retrieving the head of the merge base
	gitOutputBase, err := ExecuteCommandWithoutStdOut("merge-base", "@", "origin/"+_branch)
	if err != nil {
		return errors.New("There was an error retrieving the current head of the remote merge base")
	}
	gitOutputBase = strings.TrimSpace(gitOutputBase)

	if gitOutputLocal == gitOutputRemote {
		return nil
	} else if gitOutputLocal == gitOutputBase {
		return errors.New("The remote " + _branch + " branch is ahead of your local state. Please pull the changes before trying to merge into master")
	} else if gitOutputRemote == gitOutputBase {
		return errors.New("You didn't push all of your commits in the " + _branch + " branch. Please push before trying to merge into master")
	}

	helpers.PrintLine("The remote and local states of the "+_branch+" branch are not the same", helpers.ERROR)
	helpers.PrintLine("This usually happens if you didn't push the latest commits of the sidebranch", helpers.ERROR)
	return errors.New("Please resolve this before trying to merge into master")
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
	gitOutputBranch, err := ExecuteCommandWithoutStdOut("branch")
	if err != nil {
		return errors.New("There was an error retrieving the git repositories' branches. Is the repository set up correctly?")
	}

	// Verifying whether the connection works correctly
	gitOutputFetch, err := ExecuteCommandWithoutStdOut("fetch", "-v")
	if (!strings.Contains(gitOutputFetch, repoName) && len(strings.TrimSpace(gitOutputBranch)) > 0) || err != nil {
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
			helpers.PrintLine("Wrong username or password", helpers.ERROR)
			helpers.PrintLine("Note: If you have 2FA activated for GitHub, please go to: ", helpers.ERROR)
			helpers.PrintLine("https://github.blog/2013-09-03-two-factor-authentication/#how-does-it-work-for-command-line-git", helpers.ERROR)
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
		helpers.PrintLine("Wrong username or password", helpers.ERROR)
		helpers.PrintLine("Note: If you have 2FA activated for GitHub, please go to: ", helpers.ERROR)
		helpers.PrintLine("https://github.blog/2013-09-03-two-factor-authentication/#how-does-it-work-for-command-line-git", helpers.ERROR)
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
		helpers.PrintLine("Wrong username or password", helpers.ERROR)
		helpers.PrintLine("Note: If you have 2FA activated for GitHub, please go to: ", helpers.ERROR)
		helpers.PrintLine("https://github.blog/2013-09-03-two-factor-authentication/#how-does-it-work-for-command-line-git", helpers.ERROR)
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
		helpers.PrintLine("Wrong username or password", helpers.ERROR)
		helpers.PrintLine("Note: If you have 2FA activated for GitHub, please go to: ", helpers.ERROR)
		helpers.PrintLine("https://github.blog/2013-09-03-two-factor-authentication/#how-does-it-work-for-command-line-git", helpers.ERROR)
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

// ExecuteCommand will run git commands and redirect the output
func ExecuteCommand(_args ...string) error {
	cmd := exec.Command("git", _args[0:]...)
	defer cmd.Wait()

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

	go func() {
		wg.Wait()
	}()

	err = cmd.Wait()
	if err != nil {
		return errors.New("The git " + _args[0] + " command didn't execute properly")
	}

	return nil
}

// ExecuteCommandWithoutStdOut will run git commands and not print the output
func ExecuteCommandWithoutStdOut(_args ...string) (string, error) {
	cmd := exec.Command("git", _args[0:]...)
	defer cmd.Wait()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	cmd.Stdin = os.Stdin

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	var output string

	scannerStdout := bufio.NewScanner(stdout)
	wg.Add(1)
	go func() {
		for scannerStdout.Scan() {
			text := scannerStdout.Text()
			output += text

		}
		wg.Done()
	}()

	scannerStderr := bufio.NewScanner(stderr)
	scannerStderr.Split(customSplit)
	wg.Add(1)
	go func() {
		for scannerStderr.Scan() {
			text := scannerStderr.Text()
			output += text

		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
	}()

	err = cmd.Wait()
	if err != nil {
		return output, errors.New("The git " + _args[0] + " command didn't execute properly")
	}
	return output, nil
}

func customSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), "\r"); i >= 0 {
		return i + 1, data[0 : i+1], nil
	}

	if i := strings.Index(string(data), "\n"); i >= 0 {
		return i + 1, data[0 : i+1], nil
	}

	if atEOF {
		return len(data), data, nil
	}
	return
}
