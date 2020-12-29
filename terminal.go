package terminal

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-tty"
)

const defaultWidth uint = 80

// Terminal represents an initialized terminal interface (its width
// has been measured)
type Terminal struct {
	width   uint
	clearer string
}

// New initializes a terminal (measures its width). This was originally
// done inside `init`, but this causes issues with some debugging/testing
// setups and is now done explicitly.
func New() *Terminal {
	var err error
	var width uint

	t, err := tty.Open()
	if err == nil {
		if w, _, err := t.Size(); err == nil {
			width = uint(w)
		}
		t.Close()
	}
	if err != nil || width <= 0 {
		fmt.Fprintf(os.Stderr, "couldn't determine the width of the terminal, defaulting to %d", defaultWidth)
		width = defaultWidth
	}

	return &Terminal{
		width:   width,
		clearer: strings.Repeat(" ", int(width)-1),
	}
}

// ClearLine clears the current line in the terminal.
func (term *Terminal) ClearLine() {
	fmt.Printf("\r%s\r", term.clearer)
}

// Overwritef clears the current line in the terminal and then performs
// a fmt.Printf() with the provided args.
func (term *Terminal) Overwritef(format string, s ...interface{}) {
	fmt.Printf("\r%s\r%s", term.clearer, fmt.Sprintf(format, s...))
}

// Width returns the width of the terminal instance.
func (term *Terminal) Width() uint {
	return term.width
}
