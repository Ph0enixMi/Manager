package console

import (
	"bufio"
	"errors"
	"fmt"
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

func PrintHelp(commands_list []string, commands_description []string) {
	fmt.Println()
	for i, elem := range commands_list {
		fmt.Printf("%v - %v\n", elem, commands_description[i])
	}
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
