package main

import (
	"fmt"

	help "charm.land/bubbles/v2/help"
)

type Model struct {
	layoutList     listModel
	sizeList       listModel
	activeLayout   string
	activeSize     int
	showLayoutList bool
	showSizeList   bool
	showQuitConfirm bool
	quitSelected   int
	showInfoBar    bool
	helpModel      help.Model
	pressedKeys    map[uint16]bool
}

type GlobalKeyMsg struct {
	Code uint16
	Down bool
}

func getInitModel() Model {
	cfg := loadConfig()
	layoutList := listModel{
		items:       keyboardLayoutItems,
		selected:    0,
		title:       "Layouts",
		titleStyle:  layoutTitleStyle,
		cursorStyle: layoutCursorStyle,
	}
	for i, item := range layoutList.items {
		if item == cfg.ActiveLayout {
			layoutList.selected = i
			break
		}
	}
	sizeList := listModel{
		items:       keyboardSizeItems,
		selected:    0,
		title:       "Sizes",
		titleStyle:  sizeTitleStyle,
		cursorStyle: sizeCursorStyle,
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
		activeLayout:   cfg.ActiveLayout,
		activeSize:     cfg.ActiveSize,
		showLayoutList: false,
		showSizeList:   false,
		showInfoBar:    true,
		helpModel:      help.New(),
		pressedKeys:    make(map[uint16]bool),
	}
	initModel.helpModel.Styles = help.Styles{
		FullKey:        infoBarStyle,
		FullDesc:       infoBarStyle,
		FullSeparator:  infoBarStyle,
		ShortKey:       infoBarStyle,
		ShortDesc:      infoBarStyle,
		ShortSeparator: infoBarStyle,
		Ellipsis:       infoBarStyle,
	}

	return initModel
}
