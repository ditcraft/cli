package helpers

import (
	"bufio"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

// GetUserInputChoice will prompt the user if there are only two choices (like a or b)
func GetUserInputChoice(_prompt string, _selection1 string, _selection2 string) string {
	reader := bufio.NewReader(os.Stdin)

	var answer string
	for answer != _selection1 && answer != _selection2 {
		Printf(_prompt+" ("+_selection1+"/"+_selection2+"): ", 0)
		answer, _ = reader.ReadString('\n')
		answer = answer[0:1]
	}
	return answer
}

// GetUserInput will prompt the user if there are more possibilies than 2
func GetUserInput(_prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	Printf(_prompt+": ", 0)
	answer, _ := reader.ReadString('\n')
	answer = answer[0 : len(answer)-1]
	return answer
}

func PrintLine(_line string, _level int) {
	switch _level {
	case 0:
		fmt.Println(fmt.Sprintf("%s %s", aurora.Green("dit  >"), _line))
	case 1:
		fmt.Println(fmt.Sprintf("%s %s %s", aurora.Bold(aurora.Brown("warn")), aurora.Bold(aurora.Brown(">")), _line))
	case 2:
		fmt.Println(fmt.Sprintf("%s %s %s", aurora.Bold(aurora.BgRed(aurora.Gray("err"))), aurora.Bold(aurora.Red(" >")), _line))
	case 3:
		fmt.Println(fmt.Sprintf("%s %s %s", aurora.Cyan("demo"), aurora.Cyan(">"), _line))
	}
}

func Printf(_line string, _level int) {
	switch _level {
	case 0:
		fmt.Printf(fmt.Sprintf("%s %s", aurora.Green("dit  >"), _line))
	case 1:
		fmt.Printf(fmt.Sprintf("%s %s %s", aurora.Bold(aurora.Brown("warn")), aurora.Bold(aurora.Brown(">")), _line))
	case 2:
		fmt.Printf(fmt.Sprintf("%s %s %s", aurora.Bold(aurora.BgRed(aurora.Gray("err"))), aurora.Bold(aurora.Red(" >")), _line))
	case 3:
		fmt.Printf(fmt.Sprintf("%s %s %s", aurora.Cyan("demo"), aurora.Cyan(">"), _line))
	}
}
