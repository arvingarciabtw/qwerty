package main

import (
	"fmt"
)

type Model struct {
	activeLayout   string
	activeSize     int
	layoutList     listModel
	sizeList       listModel
	quitDialog     dialogModel
	showLayoutList bool
	showSizeList   bool
	showQuitDialog bool
	showAllInfo    bool
	pressedKeys     map[uint16]bool
	terminalWidth   int
	terminalHeight  int
}

func initModel() Model {
	cfg := loadConfig()

	layoutList := listModel{
		items:       layoutListItems,
		selected:    0,
		title:       "Layouts",
		accentColor: layoutColor,
	}
	for i, item := range layoutList.items {
		if item == cfg.ActiveLayout {
			layoutList.selected = i
			break
		}
	}

	sizeList := listModel{
		items:       layoutSizeItems,
		selected:    0,
		title:       "Sizes",
		accentColor: sizeColor,
	}
	for i, item := range sizeList.items {
		if item == fmt.Sprintf("%d%%", cfg.ActiveSize) {
			sizeList.selected = i
			break
		}
	}

	initModel := Model{
		layoutList:     layoutList,
		sizeList:       sizeList,
		quitDialog:     dialogModel{accentColor: quitColor},
		activeLayout:   cfg.ActiveLayout,
		activeSize:     cfg.ActiveSize,
		showLayoutList: false,
		showSizeList:   false,
		showAllInfo:    true,
		pressedKeys:    make(map[uint16]bool),
	}

	return initModel
}
