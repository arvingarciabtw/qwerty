package main

import (
	"strings"

	lipgloss "charm.land/lipgloss/v2"
)

type listModel struct {
	items       []string
	selected    int
	title       string
	titleStyle  lipgloss.Style
	cursorStyle lipgloss.Style
}

func (l listModel) windowStart() int {
	if len(l.items) <= 3 {
		return 0
	}
	return max(0, min(l.selected-1, len(l.items)-3))
}

func (l listModel) visibleItems() []string {
	start := l.windowStart()
	return l.items[start : start+3]
}

func (l listModel) View() string {
	var b strings.Builder

	titleLine := l.titleStyle.Render(l.title)
	b.WriteString(titleLine)
	b.WriteString("\n\n")

	maxWidth := lipgloss.Width(titleLine)

	start := l.windowStart()
	visible := l.visibleItems()

	var itemLines []string
	for i, item := range visible {
		var line string
		if start+i == l.selected {
			line = l.cursorStyle.Render("> " + item)
		} else {
			line = "  " + item
		}
		itemLines = append(itemLines, line)
		if w := lipgloss.Width(line); w > maxWidth {
			maxWidth = w
		}
	}

	for i, line := range itemLines {
		b.WriteString(line)
		if i < len(itemLines)-1 {
			b.WriteString("\n")
		}
	}

	b.WriteString("\n\n")

	help := infoBarStyle.Render("↵ / enter • q / quit")
	helpWidth := lipgloss.Width(help)
	padding := maxWidth - helpWidth
	if padding > 0 {
		b.WriteString(strings.Repeat(" ", padding))
	}
	b.WriteString(help)

	return b.String()
}

var keyboardLayoutItems = []string{
	"qwerty",
	"dvorak",
	"colemak",
	"colemak-dh",
	"workman",
	"azerty",
}

var keyboardSizeItems = []string{
	"60%",
	"65%",
	"75%",
	"80%",
	"96%",
	"100%",
}
