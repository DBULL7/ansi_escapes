package ansi_escapes

import "os"

const (
	ESC = "\u001B["
	OSC = "\u001B]"
	BEL = "\u0007"
	SEP = ";"
)

const (
	CursorLeft = ESC + "G"
	CursorGetPosition = ESC + "6n"
	CursorNextLine = ESC + "E"
	CursorPrevLine = ESC + "F"
	CursorHide = ESC + "?25l"
	CursorShow = ESC + "?25h"
	EraseEndLine = ESC + "K"
	EraseStartLine = ESC + "1k"
	EraseLine = ESC + "2K"
	EraseDown = ESC + "J"
	EraseUp = ESC + "1J"
	EraseScreen = ESC + "2J"
	ScrollUp = ESC + "S"
	ScrollDown = ESC + "T"
	ClearScreen = "\u001Bc"
	ClearTerminal = EraseScreen + ESC + "3J" + ESC + "H"
	Beep = BEL
)

var (
	IsTerminalApp = false
	CursorSavePositon = ESC + "s"
	CursorRestorePosition = ESC + "u"
)

func init() {
	termApp, exists := os.LookupEnv("TERM_PROGRAM")
	if exists && termApp == "Apple_Terminal" {
		IsTerminalApp = true
	}

	if IsTerminalApp {
		CursorSavePositon = "\u001B7"
		CursorRestorePosition = "\u001B8"
	}
}