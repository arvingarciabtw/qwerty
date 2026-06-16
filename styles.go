package main

import (
	"fmt"
	"os"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
)

func getView(m Model) string {
	termWidth, termHeight, err := getTerminalSize()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}

	if !m.showInfoBar {
		return lipgloss.Place(termWidth, termHeight, lipgloss.Left, lipgloss.Bottom, "")
	}

	bg := lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color("#0000FF")).
		TabWidth(0)

	layoutCommand := bg.Render("C-S-L: Layout")
	sizeCommand := bg.Render("C-S-S: Size")
	hideCommand := bg.Render("C-S-H: Hide Bar")
	activeLayout := bg.Render(strings.ToUpper(m.keyboardLayout))
	activeSize := bg.Render(fmt.Sprintf("%d%%", m.keyboardSize))

	commands := lipgloss.JoinHorizontal(lipgloss.Bottom, layoutCommand, " ", sizeCommand, " ", hideCommand)
	activeConfigs := lipgloss.JoinHorizontal(lipgloss.Bottom, activeLayout, " ", activeSize)

	spacerWidth := termWidth - lipgloss.Width(activeConfigs) - lipgloss.Width(commands)
	spacer := strings.Repeat(" ", spacerWidth)

	bottomBar := lipgloss.JoinHorizontal(lipgloss.Bottom, activeConfigs, spacer, commands)
	s := lipgloss.Place(termWidth, termHeight, lipgloss.Left, lipgloss.Bottom, bottomBar)

	return s
}

func overlay(bg string, popup string, x, y int) string {
	bgLines := strings.Split(bg, "\n")
	popupLines := strings.Split(popup, "\n")
	for py, pl := range popupLines {
		by := y + py
		if by < 0 || by >= len(bgLines) {
			continue
		}
		bl := bgLines[by]
		bgW := ansi.StringWidth(bl)
		pw := ansi.StringWidth(pl)

		prefix := ansi.Cut(bl, 0, x)
		var suffix string
		if x+pw < bgW {
			suffix = ansi.Cut(bl, x+pw, bgW)
		}
		bgLines[by] = prefix + pl + suffix
	}
	return strings.Join(bgLines, "\n")
}
