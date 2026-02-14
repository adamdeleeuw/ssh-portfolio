package tui

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

// Force TrueColor support for consistent colors in Docker/SSH
func init() {
	lipgloss.SetColorProfile(termenv.TrueColor)
	// Also set COLORTERM env var to signal TrueColor support
	os.Setenv("COLORTERM", "truecolor")
}

// Tokyo Night color palette
const (
	colorBackground = "#1a1b26"

	colorAccent    = "#7aa2f7" // Headings, active tab
	colorHighlight = "#bb9af7" // Links, emphasis

	colorBorder = "#414868" // Borders, dividers
	colorMuted  = "#565f89" // Dim text
)

var (
	// Header style (ASCII art container)
	headerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorAccent)).
			Bold(true).
			Align(lipgloss.Center)

	// Tab bar container
	tabBarStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(colorBorder)).
			BorderBottom(true).
			Padding(0, 1)

	// Active tab style
	activeTabStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorAccent)).
			Background(lipgloss.Color(colorBorder)).
			Bold(true).
			Padding(0, 2)

	// Inactive tab style
	inactiveTabStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorMuted)).
				Padding(0, 2)

	// Help bar style (bottom)
	helpBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted)).
			Background(lipgloss.Color(colorBackground)).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(colorBorder)).
			BorderTop(true).
			Padding(0, 1)

	// Stats bar style
	statsBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorHighlight)).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(colorBorder)).
			BorderTop(true).
			Padding(0, 1)
)
