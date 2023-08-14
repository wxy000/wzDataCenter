package utils

import (
	"golang.org/x/time/rate"
	"sync"
	"wzDataCenter/common"
)

type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

var (
	RateLimiter *IPRateLimiter
)

// SetupIPRateLimiter 创建一个RateLimiter
func SetupIPRateLimiter() error {
	var r rate.Limit
	r = 1
	b := common.CONF.Limiter.CountPerSecond
	RateLimiter = &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}
	return nil
}

// AddIP 添加一个ip到map
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip] = limiter
	return limiter
}

// GetLimiter 通过ip得到limiter
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]
	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}
	i.mu.Unlock()
	return limiter
}
