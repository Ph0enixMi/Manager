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

func GetCommand(commands_list []string) (string, string, error) {
	var command, argument string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(bar)

	if scanner.Scan() {
		command = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", "", errors.New("command input error")
	}

	command = strings.TrimSpace(command)

	if len(command) == 0 {
		return "", "", errors.New("empty input")
	}

	for _, elem := range commands_list {
		del := 0
		flag := true

		for i, rn := range elem {
			if string(rn) == "[" {
				del = i
			}
		}

		for i, rn := range elem {
			if del == 0 {
				if string(rn) != string(command[i]) {
					flag = false
					break
				}
			} else if i < del-1 && len(command) >= len(elem[:del-1]) {
				if string(rn) != string(command[i]) {
					flag = false
					break
				}
			} else if len(command) < len(elem[:del-1]) {
				flag = false
				break
			}
		}

		if del == 0 {
			if flag && len(command) == len(elem) {
				break
			} else if flag && len(command) > len(elem) {
				argument = command[len(elem)+1:]
			}
		} else {
			if flag && len(command) == len(elem[:del-1]) {
				break
			} else if flag && len(command) > len(elem[:del-1]) {
				argument = command[del:]
			}
		}
	}

	if argument != "" {
		command = command[:len(command)-len(argument)-1]
	}

	return command, strings.TrimSpace(argument), nil
}
