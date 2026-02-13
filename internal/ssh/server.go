package ssh

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gliderlabs/ssh"
)

/**
 * Starts the SSH server and begins accepting connections.
 * @param cfg - Server configuration (port, password, keys)
 * @return error if server fails to start or crashes
 * @effects Blocks current goroutine until server stops
 */
func StartServer(cfg *Config) error {
	// Load or generate host key
	signer, err := LoadOrGenerateHostKey(cfg.HostKeyPath)
	if err != nil {
		return fmt.Errorf("failed to load host key: %w", err)
	}

	// Create rate limiter
	rateLimiter := NewRateLimiter(cfg.MaxPerMinute)

	// Start cleanup goroutine
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			rateLimiter.CleanupOldLimiters()
		}
	}()

	// Configure SSH server
	server := &ssh.Server{
		Addr:             fmt.Sprintf(":%d", cfg.Port),
		Handler:          createSessionHandler(cfg),
		PasswordHandler:  createPasswordHandler(cfg.Password),
		PublicKeyHandler: nil, // No public key auth for now
	}

	server.AddHostKey(signer)

	// Connection callback for rate limiting
	server.ConnCallback = func(ctx ssh.Context, conn net.Conn) net.Conn {
		ip := getIP(conn.RemoteAddr())
		if !rateLimiter.Allow(ip) {
			log.Warn("Rate limit exceeded", "ip", ip)
			conn.Close()
			return nil
		}
		log.Info("New connection", "ip", ip)
		return conn
	}

	log.Info("Starting SSH server", "port", cfg.Port)
	return server.ListenAndServe()
}

/**
 * Creates password authentication handler.
 * @param expectedPassword - The password to check against
 * @return SSH PasswordHandler function
 */
func createPasswordHandler(expectedPassword string) ssh.PasswordHandler {
	return func(ctx ssh.Context, password string) bool {
		if password == expectedPassword {
			log.Info("Successful authentication", "user", ctx.User())
			return true
		}
		log.Warn("Failed authentication attempt", "user", ctx.User())
		return false
	}
}

/**
 * Creates the SSH session handler that manages each connection.
 * @param cfg - Server configuration
 * @return SSH Handler function
 */
func createSessionHandler(cfg *Config) ssh.Handler {
	return func(sess ssh.Session) {
		// Handle PTY requests
		ptyReq, winCh, isPty := sess.Pty()
		if !isPty {
			io.WriteString(sess, "Error: PTY required\n")
			sess.Exit(1)
			return
		}

		log.Info("Session started",
			"user", sess.User(),
			"term", ptyReq.Term,
			"width", ptyReq.Window.Width,
			"height", ptyReq.Window.Height,
		)

		// Handle window size changes
		go func() {
			for win := range winCh {
				log.Debug("Window resized", "width", win.Width, "height", win.Height)
				// TODO: Send resize event to TUI (Phase 3)
			}
		}()

		// Placeholder: In Phase 3, we'll launch Bubble Tea here
		// For now, just show a welcome message
		io.WriteString(sess, "\r\n")
		io.WriteString(sess, "╔═══════════════════════════════════════╗\r\n")
		io.WriteString(sess, "║   Welcome to Adam's SSH Portfolio!    ║\r\n")
		io.WriteString(sess, "╚═══════════════════════════════════════╝\r\n")
		io.WriteString(sess, "\r\n")
		io.WriteString(sess, "SSH server is working! \r\n")
		io.WriteString(sess, "Phase 3 will add the TUI interface.\r\n")
		io.WriteString(sess, "\r\n")
		io.WriteString(sess, "Press '~' then '.' to disconnect.\r\n")
		io.WriteString(sess, "\r\n")

		// Keep session alive until user disconnects
		<-sess.Context().Done()
		log.Info("Session ended", "user", sess.User())
	}
}
