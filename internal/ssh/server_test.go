package ssh

import (
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/gliderlabs/ssh"
	gossh "golang.org/x/crypto/ssh"
)

/**
 * Starts a test SSH server on a random port.
 * @return address, stop function, error
 */
func startTestServer(t *testing.T) (string, func(), error) {
	// Generate a random host key for testing
	hostKeyPath := "test_host_key" // In current dir
	// Cleanup previous if exists
	os.Remove(hostKeyPath)

	cfg := &Config{
		Port:         0, // Random port
		Password:     "testpass",
		MaxPerMinute: 100,
		HostKeyPath:  hostKeyPath,
	}

	// Signer load/gen
	signer, err := LoadOrGenerateHostKey(cfg.HostKeyPath)
	if err != nil {
		return "", nil, fmt.Errorf("failed to gen host key: %v", err)
	}

	// Setup server
	rateLimiter := NewRateLimiter(cfg.MaxPerMinute)
	server := &ssh.Server{
		Addr:             "127.0.0.1:0", // Localhost random port
		Handler:          createSessionHandler(),
		PasswordHandler:  createPasswordHandler(cfg.Password),
		PublicKeyHandler: nil,
	}
	server.AddHostKey(signer)
	server.ConnCallback = func(ctx ssh.Context, conn net.Conn) net.Conn {
		ip := getIP(conn.RemoteAddr())
		if !rateLimiter.Allow(ip) {
			conn.Close()
			return nil
		}
		return conn
	}

	// Listen
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return "", nil, fmt.Errorf("failed to listen: %v", err)
	}
	addr := ln.Addr().String()

	// Serve in goroutine
	go func() {
		server.Serve(ln)
	}()

	// Cleanup function
	stop := func() {
		server.Close()
		ln.Close()
		os.Remove(hostKeyPath)
	}

	return addr, stop, nil
}

/**
 * Tests successful connection and authentication.
 */
func TestServer_ConnectivityAndAuth(t *testing.T) {
	addr, stop, err := startTestServer(t)
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer stop()

	// Wait a bit for server to be ready
	time.Sleep(100 * time.Millisecond)

	// Client config
	config := &gossh.ClientConfig{
		User: "testuser",
		Auth: []gossh.AuthMethod{
			gossh.Password("testpass"),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(), // Setup for test
	}

	// Connect
	client, err := gossh.Dial("tcp", addr, config)
	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}
	defer client.Close()

	// Open session
	session, err := client.NewSession()
	if err != nil {
		t.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	// Verify we can get output (even if it fails due to TUI/PTY issues in test, connection worked)
	// Just checking connectivity here.
}

/**
 * Tests failed authentication.
 */
func TestServer_AuthFail(t *testing.T) {
	addr, stop, err := startTestServer(t)
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer stop()

	time.Sleep(100 * time.Millisecond)

	config := &gossh.ClientConfig{
		User: "testuser",
		Auth: []gossh.AuthMethod{
			gossh.Password("wrongpass"),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
	}

	_, err = gossh.Dial("tcp", addr, config)
	if err == nil {
		t.Error("Expected authentication failure, got success")
	}
}

/**
 * Tests simple session interaction (if possible without full TUI).
 * We can at least request a PTY and see if it doesn't error immediately.
 */
func TestServer_Session(t *testing.T) {
	addr, stop, err := startTestServer(t)
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer stop()
	time.Sleep(100 * time.Millisecond)

	config := &gossh.ClientConfig{
		User: "testuser",
		Auth: []gossh.AuthMethod{
			gossh.Password("testpass"),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
	}

	client, err := gossh.Dial("tcp", addr, config)
	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		t.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	// Request PTY
	modes := gossh.TerminalModes{
		gossh.ECHO:          1,
		gossh.TTY_OP_ISPEED: 14400,
		gossh.TTY_OP_OSPEED: 14400,
	}
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		t.Fatalf("request for pseudo terminal failed: %s", err)
	}

	// We can't easily wait for TUI loop in this simple test without hanging,
	// but reaching here means PTY request succeeded and session is active.
}

/**
 * Tests rate limiting rejection.
 */
func TestServer_RateLimiting(t *testing.T) {
	// Custom host key for this test to avoid conflicts
	hostKeyPath := "test_host_key_rl"
	defer os.Remove(hostKeyPath)

	cfg := &Config{
		Port:         0,
		Password:     "testpass",
		MaxPerMinute: 1, // Only 1 connection allowed per minute per IP
		HostKeyPath:  hostKeyPath,
	}

	signer, err := LoadOrGenerateHostKey(cfg.HostKeyPath)
	if err != nil {
		t.Fatalf("Failed to gen host key: %v", err)
	}

	// Create RateLimiter with explicit values
	rateLimiter := NewRateLimiter(cfg.MaxPerMinute)

	server := &ssh.Server{
		Addr:            "127.0.0.1:0",
		Handler:         createSessionHandler(),
		PasswordHandler: createPasswordHandler(cfg.Password),
	}
	server.AddHostKey(signer)

	server.ConnCallback = func(ctx ssh.Context, conn net.Conn) net.Conn {
		ip := getIP(conn.RemoteAddr())
		if !rateLimiter.Allow(ip) {
			conn.Close()
			return nil
		}
		return conn
	}

	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	addr := ln.Addr().String()
	go server.Serve(ln)
	defer server.Close()
	defer ln.Close()

	config := &gossh.ClientConfig{
		User: "testuser",
		Auth: []gossh.AuthMethod{
			gossh.Password("testpass"),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
	}

	// Attempt multiple connections.
	// With Limit=1, Burst=1:
	// 1st conn: Success (tokens=0)
	// 2nd conn: Fail (tokens=0)

	successCount := 0
	attempts := 5

	for i := 0; i < attempts; i++ {
		client, err := gossh.Dial("tcp", addr, config)
		if err == nil {
			successCount++
			client.Close()
		} else {
			// Expected failure for subsequent connections
		}
	}

	// We expect exactly 1 success.
	if successCount != 1 {
		t.Errorf("Expected 1 successful connection (rate limited), got %d successes out of %d attempts", successCount, attempts)
	}
}
