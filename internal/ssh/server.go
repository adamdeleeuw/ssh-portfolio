package ssh

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/adamdeleeuw/ssh-portfolio/internal/content"
	"github.com/adamdeleeuw/ssh-portfolio/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
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
		Handler:          createSessionHandler(),
		PasswordHandler:  createPasswordHandler(cfg.Password),
		PublicKeyHandler: nil, // No public key auth for now
	}

	server.AddHostKey(signer)

	// Connection callback for rate limiting
	server.ConnCallback = func(_ ssh.Context, conn net.Conn) net.Conn {
		// key is just the IP, no DNS lookup
		ip := getIP(conn.RemoteAddr())

		if !rateLimiter.Allow(ip) {
			// Don't log "Rate limit exceeded" for every attempt to avoid log spam,
			// but do return nil to drop the connection.

			// Small delay to slow down brute force/spam
			time.Sleep(500 * time.Millisecond)
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
 * @return SSH Handler function
 */
func createSessionHandler() ssh.Handler {
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

		// Load content tabs
		tabs, err := content.LoadTabs("./content")
		if err != nil {
			io.WriteString(sess, fmt.Sprintf("Error loading content: %v\n", err))
			sess.Exit(1)
			return
		}

		// Create TUI model
		sessionID := fmt.Sprintf("%s-%d", sess.User(), time.Now().Unix())
		model := tui.NewModel(tabs, sessionID)

		// Set initial window dimensions before starting program
		model.SetSize(ptyReq.Window.Width, ptyReq.Window.Height)

		// Create Bubble Tea program with custom input/output
		p := tea.NewProgram(
			model,
			tea.WithInput(sess),
			tea.WithOutput(sess),
			tea.WithAltScreen(),
		)

		// Handle window size changes
		go func() {
			for win := range winCh {
				p.Send(tea.WindowSizeMsg{
					Width:  win.Width,
					Height: win.Height,
				})
			}
		}()

		// Run the program (blocks until quit)
		if _, err := p.Run(); err != nil {
			log.Error("TUI error", "error", err)
		}

		log.Info("Session ended", "user", sess.User())
	}
}
