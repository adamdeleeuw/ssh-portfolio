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
