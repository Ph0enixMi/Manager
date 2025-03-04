package tasktime

import (
	"errors"
	"slices"
	"strconv"
)

func AddTimeTask(metier_list []string, task string) ([]string, error) {
	if task == "" {
		return nil, errors.New("empty text")
	}
	return append(metier_list, task), nil
}

func DeleteTimeTask(metier_list []string, task string) ([]string, error) {
	num, err := strconv.Atoi(task)
	if err != nil || len(task) == 0 {
		return metier_list, errors.New("invalid argument")
	}
	return slices.Delete(metier_list, num-1, num), nil
}

func WriteTime(time_list []string, arguments []string) ([]string, error) {
	index, err1 := strconv.Atoi(arguments[0])
	hour, err2 := strconv.Atoi(arguments[1])
	minute, err3 := strconv.Atoi(arguments[2])

	if err1 != nil || err2 != nil || err3 != nil {
		return time_list, errors.New("invalid argumetn")
	}

	if len(time_list)-1 < index {
		for range index - len(time_list) {
			time_list = append(time_list, " ")
		}
	}

	time_list[index-1] = strconv.Itoa(hour) + " ч, " + strconv.Itoa(minute) + " м"

	return time_list, nil
}
