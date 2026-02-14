package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

/**
 * Handles all user input and system events.
 * @param msg - Message from Bubble Tea (keypress, window resize, etc.)
 * @return Updated model and optional command
 * @effects Updates model state based on message type
 */
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case splashTimeoutMsg:
		// Splash screen timer finished
		m.showSplash = false
		return m, nil

	case tea.KeyMsg:
		// Allow any key to skip splash screen
		if m.showSplash {
			m.showSplash = false
			return m, nil
		}
		switch msg.String() {
		// Quit
		case "q", "ctrl+c":
			return m, tea.Quit

		// Tab navigation (using Tab key for now, gt/gT will be added later for true Vim motions)
		case "tab", "l", "right":
			m.activeTab++
			if m.activeTab >= len(m.tabs) {
				m.activeTab = 0
			}
			m.updateViewportContent()

		case "shift+tab", "h", "left":
			m.activeTab--
			if m.activeTab < 0 {
				m.activeTab = len(m.tabs) - 1
			}
			m.updateViewportContent()

		// Scrolling
		case "j", "down":
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd

		case "k", "up":
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd

		case "d", "ctrl+d":
			m.viewport.PageDown()

		case "u", "ctrl+u":
			m.viewport.PageUp()

		case "g":
			m.viewport.GotoTop()

		case "G":
			m.viewport.GotoBottom()

		// Toggle help
		case "?":
			m.showHelp = !m.showHelp
		}

	case tea.WindowSizeMsg:
		// Terminal was resized
		m.width = msg.Width
		m.height = msg.Height

		if !m.ready {
			m.ready = true
		}

		// Recalculate viewport dimensions
		// Reserve space for: header (3), tabs (3), stats (2), help (2), padding (2)
		headerHeight := 3
		tabHeight := 3
		statsHeight := 2
		helpHeight := 2
		if !m.showHelp {
			helpHeight = 0
		}

		m.viewport.Width = msg.Width - 4
		m.viewport.Height = msg.Height - headerHeight - tabHeight - statsHeight - helpHeight - 2

		m.updateViewportContent()
	}

	return m, cmd
}

/**
 * Updates viewport content when tab changes.
 * @effects Sets viewport content to current tab's content
 */
func (m *Model) updateViewportContent() {
	if m.activeTab >= 0 && m.activeTab < len(m.tabs) {
		m.viewport.SetContent(m.tabs[m.activeTab].Content)
		m.viewport.GotoTop()
	}
}
