package main

import (
	"fmt"
	"manager/console"
)

var commands_list = []string{
	"help",
	"show time",
	"create time [name]",
	"exit",
}

var commands_descriptions = []string{
	"Список команд",
	"Список всех задач",
	"Создать задачу",
	"Выход из программы",
}

func main() {
	var metier_list []string
	console.ClearScreen()
	for {
		command, err := console.GetCommand()
		if err != nil {
			fmt.Println("Неверно введена комманда")
		}

		switch command {
		case commands_list[0]:
			err := console.PrintList(commands_list, commands_descriptions)
			if err != nil {
				fmt.Println("Ошибка: комманды отсутсвуют")
			}
		case commands_list[1]:
			err := console.PrintList(metier_list, nil)
			if err != nil {
				console.PrintEmpty("Задачи отсутсвуют")
			}
		case commands_list[2]:
		case commands_list[3]:
			return
		}
	}
}
