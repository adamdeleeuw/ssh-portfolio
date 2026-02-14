package tui

import (
	"testing"
)

/**
 * Tests model initialization with default values.
 */
func TestNewModel(t *testing.T) {
	tabs := []Tab{
		{Name: "Tab1", Content: "Content 1"},
		{Name: "Tab2", Content: "Content 2"},
	}

	m := NewModel(tabs, "test-session")

	if m.activeTab != 0 {
		t.Errorf("Expected activeTab 0, got %d", m.activeTab)
	}

	if len(m.tabs) != 2 {
		t.Errorf("Expected 2 tabs, got %d", len(m.tabs))
	}

	if m.sessionID != "test-session" {
		t.Errorf("Expected sessionID 'test-session', got %s", m.sessionID)
	}

	if m.showHelp != true {
		t.Error("Expected showHelp to be true by default")
	}
}
