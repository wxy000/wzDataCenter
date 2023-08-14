package middleware

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/common"
	"wzDataCenter/utils"
)

// BlacklistMiddleware 黑名单
func BlacklistMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//得到ip地址
		ipAddr := utils.GetRealIp(ctx)
		succ, err := utils.ReadBlacklist("./common/blacklist.txt", ipAddr)
		if succ {
			common.FailWithMsg("无权限访问", ctx)
			ctx.Abort()
			return
		} else {
			if err != nil {
				common.FailWithMsg("读取黑名单发生异常", ctx)
				ctx.Abort()
				return
			} else {
				ctx.Next()
			}
		}
	}
}
