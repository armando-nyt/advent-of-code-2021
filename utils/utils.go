package utils

import (
	"errors"
	"strings"
)

func SplitIntoList(input string) ([]string, error) {
	list := strings.Split(input, "\n")
	if len(list) < 1 {
		return nil, errors.New("empty list")
	}
	return list, nil
}