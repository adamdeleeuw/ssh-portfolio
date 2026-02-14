package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

/**
 * Represents a single tab in the portfolio.
 */
type Tab struct {
	Name    string
	Content string
}

/**
 * Main Bubble Tea model holding application state.
 */
type Model struct {
	activeTab  int            // Current tab index (0-3)
	tabs       []Tab          // All tabs
	viewport   viewport.Model // Scrollable content area
	width      int            // Terminal width
	height     int            // Terminal height
	ready      bool           // Whether viewport is initialized
	showHelp   bool           // Show help bar
	showSplash bool           // Show splash screen animation
	startTime  time.Time      // Server start time for uptime
	sessionID  string         // Unique session identifier
}

/**
 * Creates a new TUI model with default state.
 * @param tabs - List of tabs to display
 * @param sessionID - Unique identifier for this session
 * @return Initialized Model
 */
func NewModel(tabs []Tab, sessionID string) Model {
	vp := viewport.New(80, 20)
	vp.MouseWheelEnabled = false // SSH doesn't support mouse

	return Model{
		activeTab:  0,
		tabs:       tabs,
		viewport:   vp,
		ready:      false,
		showHelp:   true,
		showSplash: true, // Start with splash screen
		startTime:  time.Now(),
		sessionID:  sessionID,
	}
}

/**
 * Initializes the Bubble Tea program.
 * @return Initial command to start splash timer
 */
func (m Model) Init() tea.Cmd {
	return splashTimer()
}

/**
 * Sets the window size and initializes the viewport.
 * This should be called before starting the Bubble Tea program.
 * @param width - Terminal width
 * @param height - Terminal height
 */
func (m *Model) SetSize(width, height int) {
	m.width = width
	m.height = height
	m.ready = true

	// Calculate viewport dimensions
	headerHeight := 3
	tabHeight := 3
	statsHeight := 2
	helpHeight := 2
	if !m.showHelp {
		helpHeight = 0
	}

	m.viewport.Width = width - 4
	m.viewport.Height = height - headerHeight - tabHeight - statsHeight - helpHeight - 2

	// Set initial content
	m.updateViewportContent()
}

/**
 * Message sent when splash screen timer completes.
 */
type splashTimeoutMsg struct{}

/**
 * Creates a command that sends a message after the splash duration.
 * @return Command that fires after 2.5 seconds
 */
func splashTimer() tea.Cmd {
	return tea.Tick(2500*time.Millisecond, func(t time.Time) tea.Msg {
		return splashTimeoutMsg{}
	})
}
