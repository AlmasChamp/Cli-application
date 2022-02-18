package joker

import (
	"errors"
	"fmt"
)

func checkArg(s string) (string, error) {
	if s == "random" || s == "dump" {
		return s, nil
	}
	fmt.Println(s)
	return "", errors.New("invalid arguments or options. Example: go run . random or go run . dump")
}

func checkFlag(s []string) error {
	_, err := checkArg(s[0])
	if err != nil {
		return err
	} else if len(s) == 3 && s[1] == "-n" {
		return nil
	}
	return errors.New("invalid arguments or options. Example: go run . random or go run . dump -n 3")
}
