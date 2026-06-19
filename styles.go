package main

import (
	"image/color"
	"os"

	lipgloss "charm.land/lipgloss/v2"
)

var (
	layoutListStyle   lipgloss.Style
	sizeListStyle     lipgloss.Style
	quitConfirmStyle  lipgloss.Style
	quitCursorStyle   lipgloss.Style
	fingerStyle       map[Finger]lipgloss.Style
	fingerActive      map[Finger]lipgloss.Style
	infoBarStyle      lipgloss.Style
	layoutTitleStyle  lipgloss.Style
	layoutCursorStyle lipgloss.Style
	sizeTitleStyle    lipgloss.Style
	sizeCursorStyle   lipgloss.Style
)

var keyboardBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     ",",
	TopRight:    ",",
	BottomLeft:  "'",
	BottomRight: "'",
}

var darkColors = map[Finger]color.Color{
	FingerPinky:  lipgloss.BrightMagenta,
	FingerRing:   lipgloss.BrightBlue,
	FingerMiddle: lipgloss.BrightGreen,
	FingerIndex:  lipgloss.BrightYellow,
	FingerThumb:  lipgloss.BrightCyan,
	FingerAny:    lipgloss.BrightRed,
}

var lightColors = map[Finger]color.Color{
	FingerPinky:  lipgloss.Magenta,
	FingerRing:   lipgloss.Blue,
	FingerMiddle: lipgloss.Green,
	FingerIndex:  lipgloss.Yellow,
	FingerThumb:  lipgloss.Cyan,
	FingerAny:    lipgloss.Red,
}

func init() {
	isDark := lipgloss.HasDarkBackground(os.Stdin, os.Stdout)

	colors := darkColors

	if !isDark {
		colors = lightColors
	}

	if isDark {
		layoutListStyle = lipgloss.NewStyle().
			BorderForeground(lipgloss.BrightBlue).
			Border(keyboardBorder).
			Padding(1, 3)
		sizeListStyle = lipgloss.NewStyle().
			BorderForeground(lipgloss.BrightMagenta).
			Border(keyboardBorder).
			Padding(1, 3)
		quitConfirmStyle = lipgloss.NewStyle().
			BorderForeground(lipgloss.BrightBlack).
			Border(keyboardBorder).
			Padding(1, 3)
		quitCursorStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightRed)
	} else {
		layoutListStyle = lipgloss.NewStyle().
			BorderForeground(lipgloss.Blue).
			Border(keyboardBorder).
			Padding(1, 3)
		sizeListStyle = lipgloss.NewStyle().
			BorderForeground(lipgloss.Magenta).
			Border(keyboardBorder).
			Padding(1, 3)
		quitConfirmStyle = lipgloss.NewStyle().
			BorderForeground(lipgloss.Black).
			Border(keyboardBorder).
			Padding(1, 3)
		quitCursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Red)
	}

	fingerStyle = make(map[Finger]lipgloss.Style, len(colors))
	fingerActive = make(map[Finger]lipgloss.Style, len(colors))

	for finger, c := range colors {
		base := lipgloss.NewStyle().Foreground(c)

		if isDark {
			fingerStyle[finger] = base.Faint(true)
		} else {
			fingerStyle[finger] = base
		}

		fingerActive[finger] = base.Copy().Bold(true).Italic(true)
	}

	if isDark {
		infoBarStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
		layoutTitleStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightBlue)
		layoutCursorStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightBlue)
		sizeTitleStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightMagenta)
		sizeCursorStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightMagenta)
	} else {
		infoBarStyle = lipgloss.NewStyle().Faint(true)
		layoutTitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Blue)
		layoutCursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Blue)
		sizeTitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Magenta)
		sizeCursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Magenta)
	}
}
