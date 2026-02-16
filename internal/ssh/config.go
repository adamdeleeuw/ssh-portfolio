package ssh

import (
	"os"
	"strconv"
)

/**
 * Server configuration loaded from environment variables.
 */
type Config struct {
	Port         int
	HostKeyPath  string
	MaxPerMinute int // Rate limit: connections per minute per IP
}

/**
 * Loads server configuration from environment variables with defaults.
 * @return Config populated with values from env or defaults
 */
func LoadConfig() *Config {
	port := 2222
	if p := os.Getenv("PORT"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			port = parsed
		}
	}

	hostKeyPath := os.Getenv("HOST_KEY_PATH")
	if hostKeyPath == "" {
		hostKeyPath = "./data/ssh_host_ed25519_key"
	}

	maxPerMinute := 60
	if m := os.Getenv("RATE_LIMIT"); m != "" {
		if parsed, err := strconv.Atoi(m); err == nil {
			maxPerMinute = parsed
		}
	}

	return &Config{
		Port:         port,
		HostKeyPath:  hostKeyPath,
		MaxPerMinute: maxPerMinute,
	}
}
