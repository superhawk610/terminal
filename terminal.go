package terminal

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

var terminalWidth uint
var clearer string

func init() {
	terminalWidth = getTerminalWidth()
	clearer = strings.Repeat(" ", int(terminalWidth)-1)
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getTerminalWidth() uint {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
	return uint(ws.Col)
}

// ClearLine clears the current line in the terminal
func ClearLine() {
	fmt.Printf("\r%s\r", clearer)
}

// Overwritef clears the current line in the terminal and then performs
// a fmt.Printf() with the provided args
func Overwritef(format string, s ...interface{}) {
	fmt.Printf("\r%s\r%s", clearer, fmt.Sprintf(format, s...))
}
