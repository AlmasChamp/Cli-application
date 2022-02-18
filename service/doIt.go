package joker

import (
	"errors"
	"fmt"
)

func doIt(word string, num int) error {
	if word == "random" {
		random()
	} else if word == "dump" {
		fmt.Println("herer")
		dump(num)
	}
	return errors.New("invalid arguments")
}
