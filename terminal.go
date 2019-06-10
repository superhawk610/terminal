package terminal

import (
	"fmt"
	"os"
	"strings"

	dimensions "github.com/wayneashleyberry/terminal-dimensions"
)

const defaultWidth uint = 100

var terminalWidth uint
var clearer string

func init() {
	var err error
	if terminalWidth, err = dimensions.Width(); err != nil {
		fmt.Fprintf(os.Stderr, "couldn't determine the width of the terminal, defaulting to %d", defaultWidth)
		terminalWidth = defaultWidth
	}

	if terminalWidth <= 0 {
		panic("the terminal appears to have 0 visible columns")
	}

	clearer = strings.Repeat(" ", int(terminalWidth)-1)
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
