package helpers

import (
	"bufio"
	"fmt"
	"os"
)

// GetUserInputChoice will prompt the user if there are only two choices (like a or b)
func GetUserInputChoice(_prompt string, _selection1 string, _selection2 string) string {
	reader := bufio.NewReader(os.Stdin)

	var answer string
	for answer != _selection1 && answer != _selection2 {
		fmt.Print(_prompt + " (" + _selection1 + "/" + _selection2 + "): ")
		answer, _ = reader.ReadString('\n')
		answer = answer[0:1]
	}
	return answer
}

// GetUserInput will prompt the user if there are more possibilies than 2
func GetUserInput(_prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(_prompt + ": ")
	answer, _ := reader.ReadString('\n')
	answer = answer[0 : len(answer)-1]
	return answer
}
