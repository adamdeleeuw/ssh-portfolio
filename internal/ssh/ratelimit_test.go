package ssh

import (
	"testing"
)

/**
 * Tests that rate limiter allows connections within the limit.
 */
func TestRateLimiter_Allow(t *testing.T) {
	rl := NewRateLimiter(3) // 3 requests per minute

	// First 3 should be allowed
	for i := 0; i < 3; i++ {
		if !rl.Allow("192.168.1.1") {
			t.Errorf("Request %d should have been allowed", i+1)
		}
	}

	// 4th should be rate limited (burst is 3)
	if rl.Allow("192.168.1.1") {
		t.Error("Request 4 should have been rate limited")
	}
}

/**
 * Tests that different IPs have independent rate limits.
 */
func TestRateLimiter_SeparateIPs(t *testing.T) {
	rl := NewRateLimiter(2)

	// Use up IP1's quota
	rl.Allow("192.168.1.1")
	rl.Allow("192.168.1.1")

	// IP2 should still be allowed
	if !rl.Allow("192.168.1.2") {
		t.Error("Different IP should have separate limit")
	}
}

/**
 * Tests cleanup of old limiters.
 */
func TestRateLimiter_Cleanup(t *testing.T) {
	rl := NewRateLimiter(5)

	// Create some limiters
	rl.Allow("192.168.1.1")
	rl.Allow("192.168.1.2")
	rl.Allow("192.168.1.3")

	// Cleanup
	rl.CleanupOldLimiters()

	// Map should be empty
	rl.mu.RLock()
	count := len(rl.limiters)
	rl.mu.RUnlock()

	if count != 0 {
		t.Errorf("Expected 0 limiters after cleanup, got %d", count)
	}
}
