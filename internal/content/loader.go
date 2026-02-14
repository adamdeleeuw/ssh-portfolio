package content

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adamdeleeuw/ssh-portfolio/internal/tui"
	"github.com/charmbracelet/glamour"
)

/**
 * Loads markdown content files and converts to ANSI-styled strings.
 * @param contentDir - Directory containing markdown files
 * @return Slice of tabs with rendered content
 * @return error if files cannot be loaded
 */
func LoadTabs(contentDir string) ([]tui.Tab, error) {
	// Define tab order and filenames
	tabFiles := []struct {
		name     string
		filename string
	}{
		{"Welcome", "welcome.md"},
		{"About", "about.md"},
		{"Projects", "projects.md"},
		{"Future", "future.md"},
	}

	// Create glamour renderer with dark theme for terminal
	// Enable hyperlinks for clickable links in compatible terminals (OSC 8)
	renderer, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(100),
		glamour.WithPreservedNewLines(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create markdown renderer: %w", err)
	}

	var tabs []tui.Tab

	for _, tf := range tabFiles {
		path := filepath.Join(contentDir, tf.filename)

		// Read markdown file
		content, err := os.ReadFile(path)
		if err != nil {
			// If file doesn't exist, use placeholder
			tabs = append(tabs, tui.Tab{
				Name:    tf.name,
				Content: fmt.Sprintf("Content for %s coming soon!\n\nFile not found: %s", tf.name, tf.filename),
			})
			continue
		}

		// Render markdown to ANSI
		rendered, err := renderer.Render(string(content))
		if err != nil {
			return nil, fmt.Errorf("failed to render %s: %w", tf.filename, err)
		}

		tabs = append(tabs, tui.Tab{
			Name:    tf.name,
			Content: rendered,
		})
	}

	return tabs, nil
}
