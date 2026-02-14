package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

/**
 * Renders the entire TUI as a string.
 * @return String representation of the UI
 */
func (m Model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}

	// Show splash screen if active
	if m.showSplash {
		return m.renderSplashScreen()
	}

	var b strings.Builder

	// Header
	b.WriteString(m.renderHeader())
	b.WriteString("\n\n")

	// Tabs
	b.WriteString(m.renderTabBar())
	b.WriteString("\n\n")

	// Viewport content
	b.WriteString(m.viewport.View())
	b.WriteString("\n\n")

	// Stats bar
	b.WriteString(m.renderStatsBar())

	// Help bar (if enabled)
	if m.showHelp {
		b.WriteString("\n")
		b.WriteString(m.renderHelpBar())
	}

	return b.String()
}

/**
 * Renders a compact header title bar.
 * @return Styled header string
 */
func (m Model) renderHeader() string {
	title := "╔══════════════════════════════════════╗\n" +
		"║        ADAM'S  SSH  PORTFOLIO        ║\n" +
		"╚══════════════════════════════════════╝"

	return headerStyle.Width(m.width).Render(title)
}

/**
 * Renders the tab navigation bar.
 * @return Styled tab bar string
 */
func (m Model) renderTabBar() string {
	var tabs []string

	for i, tab := range m.tabs {
		var style lipgloss.Style
		if i == m.activeTab {
			style = activeTabStyle
		} else {
			style = inactiveTabStyle
		}
		tabs = append(tabs, style.Render(tab.Name))
	}

	tabBar := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
	return tabBarStyle.Width(m.width).Render(tabBar)
}

/**
 * Renders the statistics bar showing uptime and session info.
 * @return Styled stats bar string
 */
func (m Model) renderStatsBar() string {
	uptime := time.Since(m.startTime)
	hours := int(uptime.Hours())
	minutes := int(uptime.Minutes()) % 60

	stats := fmt.Sprintf("⏱ %dh %dm • Session: %s", hours, minutes, m.sessionID)
	return statsBarStyle.Width(m.width).Render(stats)
}

/**
 * Renders the help bar with keybindings.
 * @return Styled help bar string
 */
func (m Model) renderHelpBar() string {
	help := "Tab/h/l: navigate  •  j/k: scroll  •  g/G: top/bottom  •  ?: help  •  q: quit"
	return helpBarStyle.Width(m.width).Render(help)
}
