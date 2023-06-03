package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"wzDataCenter/common"
	"wzDataCenter/models"
	"wzDataCenter/utils"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get(common.CONF.TokenName)
		if authHeader == "" {
			common.FailLoginWithMsg("无权限访问，请求未携带token", ctx)
			ctx.Abort() //结束后续操作
			return
		}
		// log.Print("token:", authHeader)

		// 解析token包含的信息
		claims, err := utils.NewJWT().ParseToken(authHeader)
		if err != nil {
			common.FailLoginWithMsg("无效的Token", ctx)
			ctx.Abort()
			return
		}

		if err := CheckUserInfo(claims); err != nil {
			common.FailLoginWithMsg(err.Error(), ctx)
			ctx.Abort()
			return
		}

		// 将当前请求的claims信息保存到请求的上下文c上
		ctx.Set("claims", claims)
		// 后续的处理函数可以用过ctx.Get("claims")来获取当前请求的用户信息
		ctx.Next()

	}
}

// CheckUserInfo 检查用户信息
func CheckUserInfo(claims *utils.CustomClaims) error {
	var user1 models.Users
	result := common.DB.Where(&models.Users{Username: claims.Username}).Find(&user1)
	// 检查用户是否存在
	if result.RowsAffected >= 1 {
		return nil
	}
	return errors.New("用户不存在")
}
