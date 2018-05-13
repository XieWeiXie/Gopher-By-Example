package infra

import (
	"errors"
	"strings"
)

var (
	ErrorStringSpace = errors.New("string trim error")
)

func HandleCommon(oldString string) (string, error) {
	newReplacer := strings.NewReplacer("\n", "", "\t", "")
	return strings.TrimSpace(newReplacer.Replace(oldString)), nil
}

func HandlerURL(oldString string) (string, error) {
	return "https://github.com" + strings.TrimSpace(oldString), nil
}
