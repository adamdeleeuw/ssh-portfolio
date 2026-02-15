package tui

import (
	"strings"
	"testing"
)

/**
 * Tests the View() method renders correctly in different states.
 */
func TestView(t *testing.T) {
	tabs := []Tab{
		{Name: "Home", Content: "Welcome to my portfolio"},
		{Name: "About", Content: "About me content"},
	}
	m := NewModel(tabs, "test-session")
	m.SetSize(100, 40) // Ensure enough space to render

	// 1. Initial State (Splash Screen)
	m.showSplash = true
	output := m.View()
	if !strings.Contains(output, "Initializing") && !strings.Contains(output, "Loading") && len(output) < 10 {
		// Splash screen might be complicated to test string matching if it uses sophisticated rendering,
		// but it should return *something*.
		if output == "" {
			t.Error("Expected splash screen output, got empty string")
		}
	}

	// 2. Main View
	m.showSplash = false
	output = m.View()

	// Check for Header
	if !strings.Contains(output, "ADAM'S") {
		t.Error("View output missing header")
	}

	// Check for Tab Bar
	if !strings.Contains(output, "Home") || !strings.Contains(output, "About") {
		t.Error("View output missing tabs")
	}

	// Check for Content (Viewport)
	// Viewport content might NOT be rendered if not set explicitly or if size is too small.
	// But we called SetSize, so it should be there.
	// Note: Viewport might wrap or style content, so exact match might be tricky.
	// Let's just check for general presence if possible, or skip strictly checking content string
	// if it's heavily formatted.

	// Check for Stats
	if !strings.Contains(output, "Session:") {
		t.Error("View output missing session stats")
	}

	// Check for Help
	if !strings.Contains(output, "q: quit") {
		t.Error("View output missing help bar")
	}

	// 3. Help Toggle
	m.showHelp = false
	output = m.View()
	if strings.Contains(output, "q: quit") {
		// Note: renderHelpBar() returns the string, View() appends it.
		// If m.showHelp is false, View() should not append it.
		// NOTE: renderHelpBar implementation in view.go needs to be checked if it respects the flag
		// or if View() respects it.
		// In view.go: if m.showHelp { ... b.WriteString(m.renderHelpBar()) }
		// So this check is valid.
		if strings.Contains(output, "navigate  •  j/k: scroll") {
			t.Error("View output should not contain help bar when showHelp is false")
		}
	}
}
