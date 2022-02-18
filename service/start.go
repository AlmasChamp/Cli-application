package joker

import (
	"errors"
	"strconv"
)

func Start(args []string) error {

	var err error
	arg := ""
	num := 0

	if len(args) == 0 {
		return errors.New("not enough arguments")
	} else if len(args) == 1 {
		arg, err = checkArg(args[0])
		if err != nil {
			return err
		}
		num = 5
		doIt(arg, num)
		return nil
	} else if len(args) == 3 {
		arg, err = checkArg(args[0])
		if err != nil {
			return err
		}

		if err = checkFlag(args[0:]); err != nil {
			return err
		}

		num, err = strconv.Atoi(args[2])
		if err != nil {
			return errors.New("invalid arguments args (not a number)")
		}
		doIt(arg, num)
		return nil
	}

	return errors.New("invalid arguments or options. Example: go run . random or go run . dump -n 3")
}
