package log

import (
	"errors"
	"io/ioutil"
	"os/user"
	"strings"
	"time"

	"github.com/shirou/gopsutil/host"
)

// DitLog Object
var DitLog []string

// DitLogEntryCount contains the amount of log entries that we have
var DitLogEntryCount int
var maxLogEntries = 100

var upcommingEntry string

// AddCommand prepares a new log entry by inserting the command
func AddCommand(command string) error {
	dt := time.Now()
	upcommingEntry = "[" + dt.Format("01-02-2006 15:04:05") + "] " + command + "\n"
	return nil
}

// AddToLog adds new line(s) to the current log entry
func AddToLog(logContent string) error {
	upcommingEntry += logContent
	return nil
}

// FinalizeEntry will add an entry to the logfile until it has reached the maximum (then old ones will be shuffled out)
func FinalizeEntry() error {
	if DitLogEntryCount >= maxLogEntries {
		DitLog = append(DitLog[:1], DitLog[2:]...)
	}
	DitLog = append(DitLog, upcommingEntry)
	err := Save()
	return err
}

// RemoveLastLine will remove the last log entry due to confidentiality reasons
func RemoveLastLine(_before string) {
	splitLine := strings.Split(upcommingEntry, _before)
	upcommingEntry = splitLine[0]
}

// Load will load the log and set it to the exported variable "DitLog"
func Load() error {
	logFile, err := GetRawLog()
	if err != nil {
		return err
	}

	DitLogEntryCount = strings.Count(string(logFile), "#--------------------") - 2
	entry := strings.Split(string(logFile), "#--------------------\n")
	for i := 0; i < len(entry); i++ {
		if entry[i] != "" && entry[i] != "\n" {
			DitLog = append(DitLog, entry[i])
		}
	}

	return nil
}

// GetRawLog func
func GetRawLog() ([]byte, error) {
	// Retrieve the home directory of the user
	usr, err := user.Current()
	if err != nil {
		return nil, errors.New("Failed to retrieve home-directory of user")
	}

	// Reading the log file
	logFile, err := ioutil.ReadFile(usr.HomeDir + "/.ditlog")
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			dt := time.Now()
			architecture := "Unknown architecture"
			arch1, arch2, arch3, hostErr := host.PlatformInformation()
			if hostErr == nil {
				architecture = arch1 + " " + arch2 + " " + arch3
			}
			logFile = []byte("#--------------------\nditCLI Log File (v1)\nCreated on " + dt.Format("01-02-2006 15:04:05") + "\nArchitecture: " + architecture + "\n#--------------------\n")
			return logFile, nil
		}
		return nil, errors.New("Failed to load log file")
	}

	return logFile, nil
}

// Save will write the current log object to the file
func Save() error {
	fileString := "#--------------------\n"
	for _, logItem := range DitLog {
		fileString += logItem + "#--------------------\n"
	}

	// Retrieve the home directory of the user
	usr, err := user.Current()
	if err != nil {
		return errors.New("Failed to retrieve home-directory of user")
	}

	// Write the above into the log file
	err = ioutil.WriteFile(usr.HomeDir+"/.ditlog", []byte(fileString), 0644)
	if err != nil {
		return errors.New("Failed to write log file")
	}

	return nil
}
