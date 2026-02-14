package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

/**
 * Renders the animated splash screen with full ASCII art.
 * @return Styled splash screen string
 */
func (m Model) renderSplashScreen() string {
	logo := `
   █████╗ ██████╗  █████╗ ███╗   ███╗███████╗    ███████╗███████╗██╗  ██╗
  ██╔══██╗██╔══██╗██╔══██╗████╗ ████║██╔════╝    ██╔════╝██╔════╝██║  ██║
  ███████║██║  ██║███████║██╔████╔██║███████╗    ███████╗███████╗███████║
  ██╔══██║██║  ██║██╔══██║██║╚██╔╝██║╚════██║    ╚════██║╚════██║██╔══██║
  ██║  ██║██████╔╝██║  ██║██║ ╚═╝ ██║███████║    ███████║███████║██║  ██║
  ╚═╝  ╚═╝╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝    ╚══════╝╚══════╝╚═╝  ╚═╝
                                                                          
  ██████╗  ██████╗ ██████╗ ████████╗███████╗ ██████╗ ██╗     ██╗ ██████╗ 
  ██╔══██╗██╔═══██╗██╔══██╗╚══██╔══╝██╔════╝██╔═══██╗██║     ██║██╔═══██╗
  ██████╔╝██║   ██║██████╔╝   ██║   █████╗  ██║   ██║██║     ██║██║   ██║
  ██╔═══╝ ██║   ██║██╔══██╗   ██║   ██╔══╝  ██║   ██║██║     ██║██║   ██║
  ██║     ╚██████╔╝██║  ██║   ██║   ██║     ╚██████╔╝███████╗██║╚██████╔╝
  ╚═╝      ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝      ╚═════╝ ╚══════╝╚═╝ ╚═════╝`

	subtitle := "\n\n                Welcome to my interactive portfolio\n"
	skip := "\n\n                   Press any key to continue..."

	splashStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colorAccent)).
		Bold(true).
		Align(lipgloss.Center)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colorHighlight)).
		Italic(true).
		Align(lipgloss.Center)

	skipStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colorMuted)).
		Align(lipgloss.Center)

	var b strings.Builder
	b.WriteString("\n\n\n")
	b.WriteString(splashStyle.Width(m.width).Render(logo))
	b.WriteString(subtitleStyle.Width(m.width).Render(subtitle))
	b.WriteString(skipStyle.Width(m.width).Render(skip))

	return b.String()
}
