package ansi_escapes

import "strconv"

func CursorTo(x, y int)  string {
	return ESC + strconv.Itoa(y + 1) + SEP  + strconv.Itoa(x + 1) + "H"
}

func CursorMove(x, y int) string {
	returnCode := ""

	if x < 0 {
		returnCode += ESC + strconv.Itoa(-x) + "D"
	} else if x > 0 {
		returnCode += ESC + strconv.Itoa(x) + "C"
	}

	if y < 0 {
		returnCode += ESC + strconv.Itoa(-y) + "A"
	} else if y > 0 {
		returnCode += ESC + strconv.Itoa(y) + "B"
	}

	return returnCode
}

func CursorUp(count int) string {
	if count <= 0 {
		count = 1
	}
	return ESC + strconv.Itoa(count) + "A"
}

func CursorDown(count int) string  {
	if count <= 0 {
		count = 1
	}
	return ESC + strconv.Itoa(count) + "B"
}

func CursorForward(count int) string {
	if count <= 0 {
		count = 1
	}
	return ESC + strconv.Itoa(count) + "C"
}

func CursorBackward(count int) string  {
	if count <= 0 {
		count = 1
	}
	return ESC + strconv.Itoa(count) + "D"
}

func Link(text, url string) string  {
	return OSC + "8" + SEP + SEP + url + BEL + text + OSC + "8" + SEP + SEP + BEL
}