package main

import (
	"fmt"
	"os"

	list "charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func main() {
	keyboardLayoutItems := []list.Item{
		item{title: "QWERTY", desc: "The default"},
		item{title: "DVORAK", desc: "For nerds"},
		item{title: "COLEMAK", desc: "For mega nerds"},
		item{title: "COLEMAK-DH", desc: "Modern variant with better thumb positioning"},
		item{title: "WORKMAN", desc: "Reduces lateral finger movement"},
		item{title: "AZERTY", desc: "Standard French layout"},
	}
	keyboardSizeItems := []list.Item{
		item{title: "60%", desc: "No backtick :("},
		item{title: "65%", desc: "Really?"},
		item{title: "75%", desc: "The classic"},
		item{title: "80%", desc: "Oooh"},
		item{title: "100%", desc: "Damn"},
	}

	initModel := Model{
		keyboardLayout:         "qwerty",
		keyboardSize:           75,
		keyboardLayoutList:     list.New(keyboardLayoutItems, list.NewDefaultDelegate(), 0, 0),
		keyboardSizeList:       list.New(keyboardSizeItems, list.NewDefaultDelegate(), 0, 0),
		showInfoBar:            true,
		showKeyboardLayoutMenu: false,
		showKeyboardSizeMenu:   false,
	}
	initModel.keyboardLayoutList.Title = "Layouts"
	initModel.keyboardLayoutList.KeyMap.Quit.SetKeys("q")
	initModel.keyboardSizeList.Title = "Sizes"
	initModel.keyboardSizeList.KeyMap.Quit.SetKeys("q")

	p := tea.NewProgram(initModel)
	_, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
