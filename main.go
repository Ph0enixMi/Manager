package main

import (
	"manager/console"
	"manager/tasktime"
)

var commands_list = []string{
	"help",
	"show time",
	"create time [name]",
	"delete time [num]",
	"clear",
	"exit",
}

var commands_descriptions = []string{
	"Список команд",
	"Список всех задач",
	"Создать задачу",
	"Удалить задачу",
	"Очистить консоль",
	"Выход из программы",
}

func main() {
	var metier_list []string
	console.ClearScreen()
	for {
		command, argument, err := console.GetCommand(commands_list)
		if err != nil {
			// console.PrintEmpty("Введите help")
		}

		del1, _ := console.GetDel(commands_list[2])
		del2, _ := console.GetDel(commands_list[3])

		switch command {
		case commands_list[0]:
			HelpCommand(argument, commands_list, commands_descriptions)
		case commands_list[1]:
			ShowTimeCommand(metier_list)
		case commands_list[2][:del1-1]:
			metier_list = CreateTimeCommand(metier_list, argument)
		case commands_list[3][:del2-1]:
			metier_list = DeleteTimeCommand(metier_list, argument)
		case commands_list[4]:
			console.ClearScreen()
		case commands_list[5]:
			return
		}
	}
}

func HelpCommand(argument string, commands_list []string, commands_descriptions []string) {
	if len(argument) > 0 {
		console.PrintEmpty("Ошибка: комманда не поддерживает аргумент")
	} else {
		err := console.PrintList(commands_list, commands_descriptions)
		if err != nil {
			console.PrintEmpty("Ошибка: комманды отсутсвуют")
		}
	}
}

func ShowTimeCommand(metier_list []string) {
	err := console.PrintList(metier_list, nil)
	if err != nil {
		console.PrintEmpty("Задачи отсутствуют")
	}
}

func CreateTimeCommand(metier_list []string, argument string) []string {
	res, err := tasktime.AddTimeTask(metier_list, argument)
	if err != nil {
		console.PrintEmpty("Пустой аргумент")
	}
	return res
}

func DeleteTimeCommand(metier_list []string, argument string) []string {
	res, err := tasktime.DeleteTimeTask(metier_list, argument)
	if err != nil {
		console.PrintEmpty("Пустой или неверный аргумент")
	}
	return res
}
