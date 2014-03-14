package command

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

// trimsplit slices s into all substrings separated by sep and returns a
// slice of the substrings between the separator with all leading and trailing
// white space removed, as defined by Unicode.
func trimsplit(s, sep string) []string {
	raw := strings.Split(s, ",")
	trimmed := make([]string, 0)
	for _, r := range raw {
		trimmed = append(trimmed, strings.TrimSpace(r))
	}
	return trimmed
}

func argOrStdin(args []string, i int) (string, error) {
	if i < len(args) {
		return args[i], nil
	}
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return "", errors.New("no available argument and stdin")
	}
	return string(bytes), nil
}
