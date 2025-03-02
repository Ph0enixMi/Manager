package main

import (
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
		command, argument, err := console.GetCommand(commands_list)
		if err != nil {
			console.PrintEmpty("Неверно введена комманда")
		}

		switch command {
		case commands_list[0]:
			if len(argument) > 0 {
				console.PrintEmpty("Ошибка, комманда не поддерживает аргумент")
			} else {
				err := console.PrintList(commands_list, commands_descriptions)
				if err != nil {
					console.PrintEmpty("Ошибка: комманды отсутсвуют")
				}
			}
		case commands_list[1]:
			err := console.PrintList(metier_list, nil)
			if err != nil {
				console.PrintEmpty("Задачи отсутсвуют")
			}
		case commands_list[2]:
			return
		case commands_list[3]:
			return
		}
	}
}
