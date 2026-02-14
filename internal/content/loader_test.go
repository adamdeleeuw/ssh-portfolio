package content

import (
	"os"
	"path/filepath"
	"testing"
)

/**
 * Tests loading tabs from markdown files.
 */
func TestLoadTabs(t *testing.T) {
	// Create temp directory with test markdown files
	tempDir := t.TempDir()

	testFiles := map[string]string{
		"welcome.md":  "# Welcome\n\nTest content",
		"about.md":    "# About\n\nAbout content",
		"projects.md": "# Projects\n\nProjects content",
		"future.md":   "# Future\n\nFuture content",
	}

	for filename, content := range testFiles {
		path := filepath.Join(tempDir, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatalf("Failed to create test file %s: %v", filename, err)
		}
	}

	// Load tabs
	tabs, err := LoadTabs(tempDir)
	if err != nil {
		t.Fatalf("LoadTabs failed: %v", err)
	}

	// Verify we got 4 tabs
	if len(tabs) != 4 {
		t.Errorf("Expected 4 tabs, got %d", len(tabs))
	}

	// Verify tab names
	expectedNames := []string{"Welcome", "About", "Projects", "Future"}
	for i, name := range expectedNames {
		if tabs[i].Name != name {
			t.Errorf("Tab %d: expected name %s, got %s", i, name, tabs[i].Name)
		}
	}

	// Verify content is rendered (should contain the original text)
	if len(tabs[0].Content) == 0 {
		t.Error("Tab content should not be empty")
	}
}

/**
 * Tests graceful handling of missing files.
 */
func TestLoadTabs_MissingFiles(t *testing.T) {
	// Use empty temp directory
	tempDir := t.TempDir()

	tabs, err := LoadTabs(tempDir)
	if err != nil {
		t.Fatalf("LoadTabs should not error on missing files: %v", err)
	}

	// Should still return 4 tabs with placeholder content
	if len(tabs) != 4 {
		t.Errorf("Expected 4 tabs with placeholders, got %d", len(tabs))
	}
}
