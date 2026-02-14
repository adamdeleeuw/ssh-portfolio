package tui

import "github.com/charmbracelet/lipgloss"

// Tokyo Night color palette
const (
	colorBackground = "#1a1b26"
	colorForeground = "#c0caf5"
	colorAccent     = "#7aa2f7" // Headings, active tab
	colorHighlight  = "#bb9af7" // Links, emphasis
	colorSuccess    = "#9ece6a" // Success messages
	colorBorder     = "#414868" // Borders, dividers
	colorMuted      = "#565f89" // Dim text
)

var (
	// Base styles
	baseStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorForeground)).
			Background(lipgloss.Color(colorBackground))

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

	// Content area style
	contentStyle = lipgloss.NewStyle().
			Padding(1, 2)

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

	// Subtle border
	borderStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(colorBorder)).
			Padding(0, 1)
)
