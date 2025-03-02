package tasktime

import (
	"errors"
)

func AddTimeTask(metier_list []string, task string) ([]string, error) {
	if task == "" {
		return nil, errors.New("empty text")
	}
	res := append(metier_list, task)
	return res, nil
}
