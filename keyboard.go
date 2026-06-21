package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"

	evdev "github.com/gvalkov/golang-evdev"
)

type globalKeyMsg struct {
	code uint16
	down bool
}

const maxEventDevices = 32

func devices() ([]*evdev.InputDevice, error) {
	var result []*evdev.InputDevice
	sawPermissionError := false

	for e := range maxEventDevices {
		device, err := evdev.Open(fmt.Sprintf("/dev/input/event%d", e))
		if err != nil {
			if os.IsPermission(err) {
				sawPermissionError = true
			}
			continue
		}
		if isKeyboardDevice(e) {
			result = append(result, device)
		} else {
			device.File.Close()
		}
	}

	if len(result) == 0 {
		if sawPermissionError {
			return nil, fmt.Errorf("permission denied reading input devices")
		}
		return nil, fmt.Errorf("no keyboard device found")
	}

	return result, nil
}

func isKeyboardDevice(eventNum int) bool {
	if checkUeventProperty(eventNum) {
		return true
	}
	return checkUdevadm(eventNum)
}

func readUeventFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	return string(data), err
}

func checkUeventProperty(eventNum int) bool {
	data, err := readUeventFile(fmt.Sprintf("/sys/class/input/event%d/device/uevent", eventNum))
	if err == nil && strings.Contains(data, "ID_INPUT_KEYBOARD=1") {
		return true
	}
	data, err = readUeventFile(fmt.Sprintf("/sys/class/input/event%d/uevent", eventNum))
	if err == nil && strings.Contains(data, "ID_INPUT_KEYBOARD=1") {
		return true
	}
	return false
}

func checkUdevadm(eventNum int) bool {
	out, err := exec.Command("udevadm", "info", "--query=property", fmt.Sprintf("--name=/dev/input/event%d", eventNum)).Output()
	if err != nil {
		return false
	}
	return strings.Contains(string(out), "ID_INPUT_KEYBOARD=1")
}

func checkInputGroup() error {
	groups, err := os.Getgroups()
	if err != nil {
		return err
	}
	lookup, err := user.LookupGroup("input")
	if err != nil {
		return err
	}
	gid, err := strconv.Atoi(lookup.Gid)
	if err != nil {
		return err
	}
	for _, g := range groups {
		if g == gid {
			return nil
		}
	}
	return fmt.Errorf("user is not in the input group")
}
