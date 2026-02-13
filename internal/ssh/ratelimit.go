package ssh

import (
	"net"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

/**
 * Rate limiter that tracks connection attempts per IP address.
 */
type RateLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.RWMutex
	rate     rate.Limit
	burst    int
}

/**
 * Creates a new rate limiter with specified requests per minute.
 * @param requestsPerMinute - Maximum requests allowed per minute per IP
 * @return Configured RateLimiter instance
 */
func NewRateLimiter(requestsPerMinute int) *RateLimiter {
	r := rate.Every(time.Minute / time.Duration(requestsPerMinute))
	return &RateLimiter{
		limiters: make(map[string]*rate.Limiter),
		rate:     r,
		burst:    requestsPerMinute,
	}
}

/**
 * Checks if a connection from the given IP should be allowed.
 * @param ip - IP address to check
 * @return true if connection is allowed, false if rate limited
 * @effects Updates internal rate limit state for the IP
 */
func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	limiter, exists := rl.limiters[ip]
	if !exists {
		limiter = rate.NewLimiter(rl.rate, rl.burst)
		rl.limiters[ip] = limiter
	}

	return limiter.Allow()
}

/**
 * Cleans up old limiters to prevent memory leaks.
 * Should be called periodically in a goroutine.
 * @effects Removes inactive IP entries from the limiter map
 */
func (rl *RateLimiter) CleanupOldLimiters() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Simple cleanup: remove all limiters
	// In production, you'd track last access time
	rl.limiters = make(map[string]*rate.Limiter)
}

/**
 * Extracts IP address from network connection.
 * @param addr - Network address (could be IP:port format)
 * @return IP address as string
 */
func getIP(addr net.Addr) string {
	if tcpAddr, ok := addr.(*net.TCPAddr); ok {
		return tcpAddr.IP.String()
	}
	// Fallback: parse from string
	host, _, _ := net.SplitHostPort(addr.String())
	return host
}
