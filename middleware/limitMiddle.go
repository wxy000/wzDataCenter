package middleware

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/common"
	"wzDataCenter/utils"
)

// LimitMiddleware 限流器
func LimitMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//得到ip地址
		ipAddr := utils.GetRealIp(ctx)
		limiter := utils.RateLimiter.GetLimiter(ipAddr)
		if !limiter.Allow() {
			common.FailWithMsg("访问超出限制", ctx)
			ctx.Abort()
			return
		} else {
			ctx.Next()
		}
	}
}
