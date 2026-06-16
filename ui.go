package main

import (
	list "charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Model struct {
	keyboardLayout         string
	keyboardSize           int
	keyboardLayoutList     list.Model
	keyboardSizeList       list.Model
	showInfoBar            bool
	showKeyboardLayoutMenu bool
	showKeyboardSizeMenu   bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			inFilter := (m.showKeyboardLayoutMenu && m.keyboardLayoutList.FilterState() == list.Filtering) ||
				(m.showKeyboardSizeMenu && m.keyboardSizeList.FilterState() == list.Filtering)
			if !inFilter {
				if m.showKeyboardLayoutMenu || m.showKeyboardSizeMenu {
					m.showKeyboardLayoutMenu = false
					m.showKeyboardSizeMenu = false
					return m, nil
				}
				return m, tea.Quit
			}
		case "ctrl+shift+l":
			m.showKeyboardLayoutMenu = !m.showKeyboardLayoutMenu
			m.showKeyboardSizeMenu = false
			return m, nil
		case "ctrl+shift+s":
			m.showKeyboardSizeMenu = !m.showKeyboardSizeMenu
			m.showKeyboardLayoutMenu = false
			return m, nil
		case "ctrl+shift+h":
			m.showInfoBar = !m.showInfoBar
			return m, nil
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.keyboardLayoutList.SetSize(msg.Width-h, msg.Height-v)
		m.keyboardSizeList.SetSize(msg.Width-h, msg.Height-v)
	}
	var cmd1, cmd2 tea.Cmd
	m.keyboardLayoutList, cmd1 = m.keyboardLayoutList.Update(msg)
	m.keyboardSizeList, cmd2 = m.keyboardSizeList.Update(msg)
	return m, tea.Batch(cmd1, cmd2)
}

func (m Model) View() tea.View {
	s := getView(m)

	if m.showKeyboardLayoutMenu {
		termW, termH, _ := getTerminalSize()

		pw := min(60, max(30, termW-4))
		ph := min(20, max(10, termH-4))
		x := (termW - pw) / 2
		y := (termH - ph) / 2

		m.keyboardLayoutList.SetSize(pw-6, ph-4)

		popup := lipgloss.NewStyle().
			Width(pw).
			Height(ph).
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Render(m.keyboardLayoutList.View())

		s = overlay(s, popup, x, y)
	}

	if m.showKeyboardSizeMenu {
		termW, termH, _ := getTerminalSize()

		pw := min(60, max(30, termW-4))
		ph := min(20, max(10, termH-4))
		x := (termW - pw) / 2
		y := (termH - ph) / 2

		m.keyboardSizeList.SetSize(pw-6, ph-4)

		popup := lipgloss.NewStyle().
			Width(pw).
			Height(ph).
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Render(m.keyboardSizeList.View())

		s = overlay(s, popup, x, y)
	}

	v := tea.NewView(s)
	v.AltScreen = true

	return v
}
