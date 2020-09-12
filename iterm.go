package ansi_escapes

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func itermSetCwd(cwd string) string {
	if cwd == "" {
		var err error
		cwd, err = os.Getwd()
		if err != nil {
			return ""
		}
	}

	return OSC + "50;CurrentDir=" + cwd + BEL
}

func itermAnnotation(message string, x, y, length int, hidden bool) (string, error) {
	ret := OSC + "1337;"

	if (x == 0 || y == 0) && !(x != 0 && y != 0 && length != 0) {
		return "", errors.New("args x, y, and length must be defined when args x or y are defined")
	}
	message = strings.ReplaceAll(message, "|", "")

	if length > 0 && x != 0 {
		ret += message + "|" + strconv.Itoa(length) + "|" + strconv.Itoa(x) + "|" + strconv.Itoa(y)
	} else if length > 0 && x == 0 {
		ret += strconv.Itoa(length) + "|" + message
	} else {
		ret += message
	}

	return ret + BEL, nil
}
