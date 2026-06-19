package main

import (
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	hasOpenList := m.showLayoutList || m.showSizeList

	switch msg := msg.(type) {
	case tea.KeyPressMsg:

		switch msg.String() {
		case "ctrl+shift+l":
			m.showLayoutList = !m.showLayoutList
			m.showSizeList = false
			return m, nil

		case "ctrl+shift+s":
			m.showSizeList = !m.showSizeList
			m.showLayoutList = false
			return m, nil

		case "ctrl+shift+h":
			m.showInfoBar = !m.showInfoBar
			return m, nil

		case "?":
			if hasOpenList || m.showQuitConfirm {
				break
			}
			m.helpModel.ShowAll = !m.helpModel.ShowAll
			return m, nil

		case "up", "k", "left":
			if m.showLayoutList {
				if m.layoutList.selected > 0 {
					m.layoutList.selected--
				}
				return m, nil
			}
			if m.showSizeList {
				if m.sizeList.selected > 0 {
					m.sizeList.selected--
				}
				return m, nil
			}

		case "down", "j", "right":
			if m.showLayoutList {
				if m.layoutList.selected < len(m.layoutList.items)-1 {
					m.layoutList.selected++
				}
				return m, nil
			}
			if m.showSizeList {
				if m.sizeList.selected < len(m.sizeList.items)-1 {
					m.sizeList.selected++
				}
				return m, nil
			}

		case "enter":
			if m.showLayoutList {
				m.activeLayout = strings.ToLower(m.layoutList.items[m.layoutList.selected])
				m.showLayoutList = false
				saveConfig(Config{ActiveLayout: m.activeLayout, ActiveSize: m.activeSize})
				return m, nil
			}
			if m.showSizeList {
				sizeStr := strings.TrimSuffix(m.sizeList.items[m.sizeList.selected], "%")
				if size, err := strconv.Atoi(sizeStr); err == nil {
					m.activeSize = size
				}
				m.showSizeList = false
				saveConfig(Config{ActiveLayout: m.activeLayout, ActiveSize: m.activeSize})
				return m, nil
			}

			}
			if hasOpenList {
				m.showLayoutList = false
				m.showSizeList = false
				return m, nil
			}
			return m, tea.Quit

		case "ctrl+c":
			return m, tea.Quit
		}

	case GlobalKeyMsg:
		m.pressedKeys[msg.Code] = msg.Down

	case tea.WindowSizeMsg:
		m.helpModel.SetWidth(msg.Width)
	}

	return m, nil
}
