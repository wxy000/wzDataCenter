package utils

import (
	"github.com/gin-gonic/gin"
	"net"
	"strings"
)

// GetRealIp 获取ip地址
func GetRealIp(ctx *gin.Context) string {
	ip, _, err := net.SplitHostPort(ctx.Request.RemoteAddr)
	if err != nil {
		ip = ctx.Request.RemoteAddr
	}
	if ip != "127.0.0.1" {
		return ip
	}
	// Check if behide nginx or apache
	xRealIP := ctx.Request.Header.Get("X-Real-Ip")
	xForwardedFor := ctx.Request.Header.Get("X-Forwarded-For")

	for _, address := range strings.Split(xForwardedFor, ",") {
		address = strings.TrimSpace(address)
		if address != "" {
			return address
		}
	}

	if xRealIP != "" {
		return xRealIP
	}
	return ip
}
