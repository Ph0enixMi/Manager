package main

import (
	"fmt"
	"manager/console"
)

var commands_list = []string{
	"help",
	"exit",
}

var commands_descriptions = []string{
	"Выводит список команд",
	"Выход из программы",
}

func main() {
	console.ClearScreen()
	for {
		command, err := console.GetCommand()
		if err != nil {
			fmt.Println("Неверно введена комманда")
		}

		switch command {
		case commands_list[0]:
			console.PrintHelp(commands_list, commands_descriptions)
		case commands_list[1]:
			return
		}
	}
}
