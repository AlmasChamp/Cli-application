package joker

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {

	//Arrange
	testTable := []struct {
		Name     string
		Args     []string
		Expected error
	}{
		{
			Name:     "emptyLine",
			Args:     []string{""},
			Expected: errors.New("invalid arguments or options. Example: go run . random or go run . dump"),
		},
		{
			Name:     "noArguments",
			Args:     []string{},
			Expected: errors.New("not enough arguments"),
		},
		{
			Name:     "manyArguments",
			Args:     []string{"dump", "-n", "2", "34"},
			Expected: errors.New("invalid arguments or options. Example: go run . random or go run . dump -n 3"),
		},
		{
			Name:     "inCorrectOptions",
			Args:     []string{"dump", "-p", "2"},
			Expected: errors.New("invalid arguments or options. Example: go run . random or go run . dump -n 3"),
		},
		{
			Name:     "inCorrectArguments",
			Args:     []string{"dump", "-n"},
			Expected: errors.New("invalid arguments or options. Example: go run . random or go run . dump -n 3"),
		},
		{
			Name:     "argsNotNumber",
			Args:     []string{"dump", "-n", "d"},
			Expected: errors.New("invalid arguments args (not a number)"),
		},
		{
			Name:     "inCorrectArguments",
			Args:     []string{"google", "-n", "2"},
			Expected: errors.New("invalid arguments or options. Example: go run . random or go run . dump"),
		},
	}

	// Act
	for _, testCase := range testTable {
		result := Start(testCase.Args)

		t.Logf("Calling Init(%v),result %d\n", testCase.Args, result)

		//Assert
		if assert.Equal(t, testCase.Expected, result) != true {
			t.Errorf(testCase.Name)
		}
	}

}
