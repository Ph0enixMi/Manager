package main

import (
	"manager/console"
	"manager/tasktime"
)

var commands_list = []string{
	"help",
	"show time",
	"create time [name]",
	"delete time [index]",
	"write time [index] [time]",
	"clear",
	"exit",
}

var commands_descriptions = []string{
	"Список команд",
	"Список всех задач",
	"Создать задачу",
	"Удалить задачу",
	"Записать время на задачу (пример: 1 10)",
	"Очистить консоль",
	"Выход из программы",
}

func main() {
	var metier_list []string
	var time_list []string
	console.ClearScreen()
	for {
		command, arguments, err := console.GetCommand(commands_list)
		if err != nil {
			// console.PrintEmpty("Введите help")
		}

		del1, _ := console.GetDel(commands_list[2])
		del2, _ := console.GetDel(commands_list[3])
		del3, _ := console.GetDel(commands_list[4])

		// fmt.Println(arguments, len(arguments))
		// fmt.Println(command, len(command))

		switch command {
		case commands_list[0]:
			HelpCommand(arguments, commands_list, commands_descriptions)
		case commands_list[1]:
			ShowTimeCommand(metier_list, time_list)
		case commands_list[2][:del1[0]-1]:
			metier_list = CreateTimeCommand(metier_list, arguments)
		case commands_list[3][:del2[0]-1]:
			metier_list = DeleteTimeCommand(metier_list, arguments)
		case commands_list[4][:del3[0]-1]:
			time_list = WriteTimeCommand(metier_list, time_list, arguments)
		case commands_list[5]:
			console.ClearScreen()
		case commands_list[6]:
			return
		}
	}
}

func HelpCommand(argument []string, commands_list []string, commands_descriptions []string) {
	if len(argument) > 0 {
		console.PrintEmpty("Ошибка: комманда не поддерживает аргумент")
	} else {
		err := console.PrintList(commands_list, commands_descriptions, false)
		if err != nil {
			console.PrintEmpty("Ошибка: комманды отсутсвуют")
		}
	}
}

func ShowTimeCommand(metier_list []string, time_list []string) {
	if len(time_list) != 0 {
		console.PrintList(metier_list, time_list, true)
	} else {
		err := console.PrintList(metier_list, nil, false)
		if err != nil {
			console.PrintEmpty("Задачи отсутствуют")
		}
	}
}

func CreateTimeCommand(metier_list []string, argument []string) []string {
	res, err := tasktime.AddTimeTask(metier_list, argument[0])
	if err != nil {
		console.PrintEmpty("Пустой аргумент")
	}
	return res
}

func DeleteTimeCommand(metier_list []string, argument []string) []string {
	res, err := tasktime.DeleteTimeTask(metier_list, argument[0])
	if err != nil {
		console.PrintEmpty("Пустой или неверный аргумент")
	}
	return res
}

func WriteTimeCommand(metier_list []string, time_list []string, arguments []string) []string {
	if len(arguments) == 0 || len(arguments) > 3 {
		console.PrintEmpty("Недопустнипое количество атрибутов")
		return time_list
	}
	res, err := tasktime.WriteTime(time_list, arguments)
	if err != nil {
		console.PrintEmpty("Неправельно введены аргумент{ы}")
	}
	return res
}
