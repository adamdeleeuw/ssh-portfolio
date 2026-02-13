package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/adamdeleeuw/ssh-portfolio/internal/ssh"
	"github.com/charmbracelet/log"
)

/**
 * Main entry point for the SSH portfolio server.
 * Loads configuration and starts the SSH server.
 */
func main() {
	// Set up structured logging
	log.SetLevel(log.InfoLevel)
	log.SetReportTimestamp(true)

	// Load configuration
	cfg := ssh.LoadConfig()

	log.Info("SSH Portfolio Server")
	log.Info("Configuration loaded",
		"port", cfg.Port,
		"hostKeyPath", cfg.HostKeyPath,
		"rateLimit", fmt.Sprintf("%d/min", cfg.MaxPerMinute),
	)

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Info("Shutdown signal received, stopping server...")
		os.Exit(0)
	}()

	// Start SSH server (blocks)
	if err := ssh.StartServer(cfg); err != nil {
		log.Fatal("Failed to start server", "error", err)
	}
}
