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
		return metier_list, errors.New("wrong argument")
	}
	return slices.Delete(metier_list, num-1, num), nil
}
