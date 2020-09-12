package ansi_escapes

import (
	"strconv"
	b64 "encoding/base64"
)

func ClearLines(count int) string {
	clear := ""
	for i := 0; i < count; i++ {
		if i < count - 1 {
			clear += EraseLine + CursorUp(1)
		} else {
			clear += EraseLine
		}
	}
	if count > 0 {
		clear += CursorLeft
	}
	return clear
}

func Image(buf []byte, width, height int, preserveAspectRatio bool) string {
	ret := OSC + "1337;File=inline=1"

	if width > 0 {
		ret += ";width=" + strconv.Itoa(width)
	}
	if height > 0 {
		ret += ";height=" + strconv.Itoa(height)
	}
	if preserveAspectRatio == false {
		ret += ";preserveAspectRatio=0"
	}

	return ret + ":" + b64.StdEncoding.EncodeToString(buf) + BEL
}