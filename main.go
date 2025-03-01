package main

import (
	"bufio"
	"errors"
	"fmt"
	"manager/console"
	"os"
	"strings"
)

var commands_list = []string{"help", "exit"}
var commands_descriptions = []string{
	"Выводит список команд",
	"Выход из программы",
}

func main() {
	console.ClearScreen()
	for {
		command, _ := GetCommand()

		switch command {
		case commands_list[0]:
			PrintHelp(commands_list, commands_descriptions)
		case commands_list[1]:
			return
		}
	}
}

func GetCommand() (string, error) {
	var command string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		command = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", errors.New("command input error")
	}

	return strings.TrimSpace(command), nil
}

func PrintHelp(commands_list []string, commands_description []string) {
	fmt.Println()
	for i, elem := range commands_list {
		fmt.Printf("%v - %v\n", elem, commands_description[i])
	}
	fmt.Println()
}
