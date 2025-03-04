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
	spaces := utils.GetSpacesBar(bar)
	mx := utils.GetMaxLn(list)
	for i, elem := range list {
		if description != nil {
			fmt.Printf("%v%v%v - %v\n", spaces, utils.GetSpaces(elem, mx), elem, description[i])
		} else {
			fmt.Printf("%v%v: %v\n", spaces, i+1, elem)
		}
	}
	fmt.Println()
	return nil
}

func PrintEmpty(text string) {
	spaces := utils.GetSpacesBar(bar)
	fmt.Println()
	fmt.Printf("%v%v\n", spaces, text)
	fmt.Println()
}

func GetCommand(commands_list []string) (string, []string, error) {
	var command string
	var arguments []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(bar)

	if scanner.Scan() {
		command = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", arguments, errors.New("command input error")
	}

	command = strings.TrimSpace(command)

	if len(command) == 0 {
		return "", arguments, errors.New("empty input")
	}

	for _, elem := range commands_list {
		del, _ := GetDel(elem)
		flag := true

		for i, rn := range elem {
			if len(del) == 0 {
				if string(rn) != string(command[i]) {
					flag = false
					break
				}
			} else if i < del[0]-1 && len(command) >= len(elem[:del[0]-1]) {
				if string(rn) != string(command[i]) {
					flag = false
					break
				}
			} else if len(command) < len(elem[:del[0]-1]) {
				flag = false
				break
			}
		}

		if len(del) == 0 {
			if flag && len(command) == len(elem) {
				break
			} else if flag && len(command) > len(elem) {
				arguments = append(arguments, command[len(elem)+1:])
			}
		} else {
			if flag && len(command) == len(elem[:del[0]-1]) {
				break
			} else if flag && len(command) > len(elem[:del[0]-1]) {
				if len(del) == 1 {
					arguments = append(arguments, command[del[0]:])
				} else {
					new_arguments := GetSpaceArguments(command[del[0]:])
					arguments = append(arguments, new_arguments...)
				}
			}
		}
	}

	if len(arguments) != 0 {
		// Вырезать список аргументов
		command = command[:len(command)-len(arguments[0])-1]
	}
	for i, arg := range arguments {
		arguments[i] = strings.TrimSpace(arg)
	}

	return command, arguments, nil
}

func GetDel(command string) ([]int, error) {
	var del_list []int
	for i, rn := range command {
		if string(rn) == "[" {
			del_list = append(del_list, i)
		}
	}
	if len(del_list) == 0 {
		return del_list, errors.New("command dont have [")
	}
	return del_list, nil
}

func GetSpaceArguments(arguments string) []string {
	var arguments_list []string
	var del []int
	for i, rn := range arguments {
		if string(rn) == " " {
			del = append(del, i)
		}
	}

	if len(del) == 0 {
		return arguments_list
	}

	for i, v := range del {
		if i == 0 {
			arguments_list = append(arguments_list, arguments[:v])
		} else if i > 0 && i <= len(del)-1 {
			arguments_list = append(arguments_list, arguments[del[i-1]+1:v])
		}
		if i == len(del)-1 {
			arguments_list = append(arguments_list, arguments[v:])
		}
	}
	return arguments_list
}
