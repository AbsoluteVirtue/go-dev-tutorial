package go_module

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name not provided")
	}
	msg := fmt.Sprintf(randomFormat(), name)
	return msg, nil
}

func HelloList(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	format := []string{
		"Hi, %v, sup!",
		"Hello, esteemed %v!",
	}
	return format[rand.Intn(len(format))]
}
