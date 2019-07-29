package helpers

import (
	"bufio"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

type outputType int

// outputTypes that can be printed
const (
	ERROR outputType = 0
	WARN  outputType = 1
	INFO  outputType = 2
	DEMO  outputType = 3
)

// GetUserInputChoice will prompt the user if there are only two choices (like a or b)
func GetUserInputChoice(_prompt string, _selection1 string, _selection2 string) string {
	reader := bufio.NewReader(os.Stdin)

	var answer string
	for answer != _selection1 && answer != _selection2 {
		Printf(_prompt+" ("+_selection1+"/"+_selection2+"): ", INFO)
		answer, _ = reader.ReadString('\n')
		answer = answer[0:1]
	}
	return answer
}

// GetUserInput will prompt the user if there are more possibilies than 2
func GetUserInput(_prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	Printf(_prompt+": ", INFO)
	answer, _ := reader.ReadString('\n')
	answer = answer[0 : len(answer)-1]
	return answer
}

// ColorizeCommand prints a dit command that will have a certain color to highlight it
func ColorizeCommand(_command string) string {
	return fmt.Sprintf("%s", aurora.Green("dit "+_command))
}

// PrintLine will print a log output with a line ending
func PrintLine(_line string, _level outputType) {
	Printf(_line, _level)
	fmt.Printf("\n")
}

// Printf will print a log output without a line ending
func Printf(_line string, _level outputType) {
	switch _level {
	case INFO:
		fmt.Printf(fmt.Sprintf("%s", _line))
	case WARN:
		fmt.Printf(fmt.Sprintf("%s %s %s", aurora.Bold(aurora.Brown("warn")), aurora.Bold(aurora.Brown(">")), _line))
	case ERROR:
		fmt.Printf(fmt.Sprintf("%s %s %s", aurora.Bold(aurora.BgRed(aurora.Gray(1-1, "err"))), aurora.Bold(aurora.Red(" >")), _line))
	case DEMO:
		fmt.Printf(fmt.Sprintf("%s %s %s", aurora.Cyan("demo"), aurora.Cyan(">"), _line))
	}
}
