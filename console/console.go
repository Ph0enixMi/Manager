package console

import (
	"bufio"
	"errors"
	"fmt"
	"manager/utils"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const bar string = "> "

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func PrintList(list []string, description []string) error {
	if len(list) == 0 {
		return errors.New("empty list")
	}

	fmt.Println()
	spaces := utils.GetSpaces(bar)
	for i, elem := range list {
		if description != nil {
			fmt.Printf("%v%v - %v\n", spaces, elem, description[i])
		} else {
			fmt.Printf("%v%v: %v\n", spaces, i, elem)
		}
	}
	fmt.Println()
	return nil
}

func PrintEmpty(text string) {
	spaces := utils.GetSpaces(bar)
	fmt.Println()
	fmt.Printf("%v%v\n", spaces, text)
	fmt.Println()

}

func GetCommand() (string, error) {
	var command string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(bar)

	if scanner.Scan() {
		command = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", errors.New("command input error")
	}
	return strings.TrimSpace(command), nil
}
