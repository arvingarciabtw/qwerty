package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	evdev "github.com/gvalkov/golang-evdev"
)

func main() {
	devs, err := devices()
	if err != nil {
		printDeviceError(err)
		os.Exit(1)
	}

	p := tea.NewProgram(initModel())
	for _, dev := range devs {
		go listenToKeyboard(p, dev)
	}
	_, err = p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}

func listenToKeyboard(p *tea.Program, dev *evdev.InputDevice) {
	defer dev.File.Close()

	for {
		events, err := dev.Read()
		if err != nil {
			return
		}
		for _, ev := range events {
			if ev.Type == evdev.EV_KEY {
				p.Send(globalKeyMsg{
					code: uint16(ev.Code),
					down: ev.Value != 0,
				})
			}
		}
	}
}

func printDeviceError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)

	exe, exeErr := os.Executable()
	if exeErr != nil {
		exe = "qwerty"
	}

	fmt.Fprintf(os.Stderr, `
This app reads raw evdev keyboard events directly (rather than through
a display server) in order to work inside the TUI. That requires
read access to /dev/input/event*, which isn't readable by normal
users by default.

Fix: sudo setcap cap_dac_read_search=ep %s

This grants read access to just this binary. It doesn't run as
root, just bypasses one permission check.

Revoke anytime with: sudo setcap -r %s

Note: re-run this after rebuilding/reinstalling the binary.
`, exe, exe)
}
