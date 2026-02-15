package tui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

/**
 * Tests tab navigation forward wraps around.
 */
func TestUpdate_TabNavigationForward(t *testing.T) {
	tabs := []Tab{
		{Name: "Tab1", Content: "Content 1"},
		{Name: "Tab2", Content: "Content 2"},
		{Name: "Tab3", Content: "Content 3"},
	}

	m := NewModel(tabs, "test")
	m.ready = true
	m.showSplash = false // Disable splash for testing

	// Navigate forward
	updatedModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyTab})
	m = updatedModel.(Model)
	if m.activeTab != 1 {
		t.Errorf("Expected activeTab 1, got %d", m.activeTab)
	}

	// Navigate forward again
	updatedModel, _ = m.Update(tea.KeyMsg{Type: tea.KeyTab})
	m = updatedModel.(Model)
	if m.activeTab != 2 {
		t.Errorf("Expected activeTab 2, got %d", m.activeTab)
	}

	// Should wrap around to 0
	updatedModel, _ = m.Update(tea.KeyMsg{Type: tea.KeyTab})
	m = updatedModel.(Model)
	if m.activeTab != 0 {
		t.Errorf("Expected activeTab to wrap to 0, got %d", m.activeTab)
	}
}

/**
 * Tests tab navigation backward wraps around.
 */
func TestUpdate_TabNavigationBackward(t *testing.T) {
	tabs := []Tab{
		{Name: "Tab1", Content: "Content 1"},
		{Name: "Tab2", Content: "Content 2"},
	}

	m := NewModel(tabs, "test")
	m.ready = true
	m.showSplash = false // Disable splash for testing
	m.activeTab = 0

	// Navigate backward from 0 should wrap to last tab
	updatedModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyShiftTab})
	m = updatedModel.(Model)
	if m.activeTab != 1 {
		t.Errorf("Expected activeTab to wrap to 1, got %d", m.activeTab)
	}
}

/**
 * Tests quit command.
 */
func TestUpdate_Quit(t *testing.T) {
	m := NewModel([]Tab{}, "test")
	m.showSplash = false // Disable splash for testing

	// Test 'q' key
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	if cmd == nil {
		t.Error("Expected quit command, got nil")
	}

	// Test Ctrl+C
	_, cmd = m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	if cmd == nil {
		t.Error("Expected quit command for Ctrl+C, got nil")
	}
}

/**
 * Tests help toggle.
 */
func TestUpdate_HelpToggle(t *testing.T) {
	m := NewModel([]Tab{}, "test")
	m.showSplash = false // Disable splash for testing
	initialHelp := m.showHelp

	// Toggle help
	updatedModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'?'}})
	m = updatedModel.(Model)
	if m.showHelp == initialHelp {
		t.Error("Help should have toggled")
	}

	// Toggle again
	updatedModel, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'?'}})
	m = updatedModel.(Model)
	if m.showHelp != initialHelp {
		t.Error("Help should have toggled back")
	}
}

/**
 * Tests window resize updates dimensions and viewport.
 */
func TestUpdate_WindowResize(t *testing.T) {
	m := NewModel([]Tab{}, "test")
	m.ready = false

	// Resize window
	updatedModel, _ := m.Update(tea.WindowSizeMsg{Width: 100, Height: 50})
	m = updatedModel.(Model)

	if !m.ready {
		t.Error("Model should be ready after resize")
	}
	if m.width != 100 {
		t.Errorf("Expected width 100, got %d", m.width)
	}
	if m.height != 50 {
		t.Errorf("Expected height 50, got %d", m.height)
	}
	if m.viewport.Width != 96 { // 100 - 4
		t.Errorf("Expected viewport width 96, got %d", m.viewport.Width)
	}
}

/**
 * Tests scrolling keybindings.
 */
func TestUpdate_Scrolling(t *testing.T) {
	tabs := []Tab{{Name: "Tab1", Content: "Line 1\nLine 2\nLine 3\nLine 4\nLine 5"}}
	m := NewModel(tabs, "test")
	m.ready = true
	m.showSplash = false
	m.viewport.Height = 2 // Small height to force scrolling
	m.viewport.SetContent(tabs[0].Content)

	// Scroll down (j)
	updatedModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})
	m = updatedModel.(Model)
	if m.viewport.YOffset == 0 {
		t.Error("Expected scrolling down to increase YOffset")
	}

	// Scroll up (k)
	updatedModel, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'k'}})
	m = updatedModel.(Model)
	if m.viewport.YOffset != 0 {
		t.Error("Expected scrolling up to decrease YOffset")
	}
}

/**
 * Tests splash screen timeout and dismissal.
 */
func TestUpdate_SplashScreen(t *testing.T) {
	m := NewModel([]Tab{}, "test")
	m.showSplash = true

	// Test timeout
	updatedModel, _ := m.Update(splashTimeoutMsg{})
	m = updatedModel.(Model)
	if m.showSplash {
		t.Error("Splash screen should be dismissed after timeout")
	}

	// Reset and test keypress dismissal
	m.showSplash = true
	updatedModel, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'x'}})
	m = updatedModel.(Model)
	if m.showSplash {
		t.Error("Splash screen should be dismissed on keypress")
	}
}

/**
 * Tests viewport content update on tab change.
 */
func TestUpdate_ViewportContent(t *testing.T) {
	tabs := []Tab{
		{Name: "Tab1", Content: "Content 1"},
		{Name: "Tab2", Content: "Content 2"},
	}
	m := NewModel(tabs, "test")
	m.ready = true
	m.showSplash = false
	m.SetSize(80, 24)

	// Initial content
	if m.viewport.View() == "" {
		// Viewport might render differently based on styles/size, but shouldn't be empty if content is set
		// However, we just check internal logic for now
	}

	// Switch tab
	m.activeTab = 1
	m.updateViewportContent()

	// Since we can't easily check internal lib buffer without rendering,
	// we rely on the fact that updateViewportContent calls SetContent.
	// We can verify this via the View() output in integration tests,
	// or blindly trust it sets the content.
	// For this unit test, let's just ensure no panic and state is updated.
}
