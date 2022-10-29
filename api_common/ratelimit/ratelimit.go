package ratelimit

import (
	"sync"

	"golang.org/x/time/rate"
)

// IPRateLimiter an ip rate limiting object
type IPRateLimiter struct {
	ipTokenMap      map[string]*rate.Limiter
	mu              *sync.RWMutex
	tokensPerSecond rate.Limit
	maxBurstSize    int64
}

// NewIPRateLimiter inits and returns a new IPRateLimiter
func NewIPRateLimiter(tokensPerSecond rate.Limit, maxBurstSize int64) *IPRateLimiter {
	i := &IPRateLimiter{
		ipTokenMap:      make(map[string]*rate.Limiter),
		mu:              &sync.RWMutex{},
		tokensPerSecond: tokensPerSecond,
		maxBurstSize:    maxBurstSize,
	}
	return i
}

// AddIP creates a new rate limiter and adds it to the ips map,
// using the IP address as the key
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()
	limiter := rate.NewLimiter(i.tokensPerSecond, int(i.maxBurstSize))
	i.ipTokenMap[ip] = limiter
	return limiter
}

// GetLimiter returns the rate limiter for the provided IP address if it exists.
// Otherwise, calls AddIP to add IP address to the map
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ipTokenMap[ip]
	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}
	i.mu.Unlock()
	return limiter
}
