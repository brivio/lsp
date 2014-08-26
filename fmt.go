// fmt.go, like render.go, contains stuff concerning output formatting in the stdoutt/terminal,
// but fmt.go is for more bash-specific/lower level stuff
package main

import (
	"fmt"
	"strings"
	"syscall"
	"unicode/utf8"
	"unsafe"

	c "github.com/mitchellh/colorstring"
)

const (
	dashesNumber = 2
)

var (
	terminalWidth   = 80
	columnSize      = 39 // characters in the filename column
	maxFileNameSize = columnSize - 7
)

func printCentered(o string) {
	length := utf8.RuneCount([]byte(o))
	sideburns := (6+2*columnSize-length)/2 - dashesNumber
	fmt.Printf(strings.Repeat(" ", sideburns))
	fmt.Printf(c.Color("[red]" + strings.Repeat("-", dashesNumber)))
	fmt.Printf(c.Color("[red]" + o + "[white]"))
	fmt.Printf(c.Color("[red]"+strings.Repeat("-", dashesNumber)) + "\n")
}

// SetTerminalSize returns the dimensions of the given terminal.
func SetColumnSize() {
	const stdoutFD = 1
	var dimensions [4]uint16

	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(stdoutFD), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&dimensions)), 0, 0, 0); err != 0 {
		return
	}
	terminalWidth = int(dimensions[1])
	if terminalWidth < 3 {
		return
	}
	columnSize = (terminalWidth - 2) / 2
}

func printHR() {
	fmt.Printf(c.Color("\n[cyan]" + strings.Repeat("-", terminalWidth) + "\n"))
}
