package ssh

import (
	"os"
	"testing"
)

/**
 * Tests configuration loading with environment variables.
 */
func TestLoadConfig(t *testing.T) {
	// Set environment variables
	os.Setenv("PORT", "3000")
	os.Setenv("SSH_PASSWORD", "test123")
	os.Setenv("RATE_LIMIT", "10")
	defer func() {
		os.Unsetenv("PORT")
		os.Unsetenv("SSH_PASSWORD")
		os.Unsetenv("RATE_LIMIT")
	}()

	cfg := LoadConfig()

	if cfg.Port != 3000 {
		t.Errorf("Expected port 3000, got %d", cfg.Port)
	}
	if cfg.Password != "test123" {
		t.Errorf("Expected password 'test123', got '%s'", cfg.Password)
	}
	if cfg.MaxPerMinute != 10 {
		t.Errorf("Expected rate limit 10, got %d", cfg.MaxPerMinute)
	}
}

/**
 * Tests default configuration values.
 */
func TestLoadConfig_Defaults(t *testing.T) {
	// Clear any env vars
	os.Unsetenv("PORT")
	os.Unsetenv("SSH_PASSWORD")
	os.Unsetenv("RATE_LIMIT")

	cfg := LoadConfig()

	if cfg.Port != 2222 {
		t.Errorf("Expected default port 2222, got %d", cfg.Port)
	}
	if cfg.Password != "portfolio" {
		t.Errorf("Expected default password 'portfolio', got '%s'", cfg.Password)
	}
	if cfg.MaxPerMinute != 5 {
		t.Errorf("Expected default rate limit 5, got %d", cfg.MaxPerMinute)
	}
}
