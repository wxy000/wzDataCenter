package zentao_controllers

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/app/zentao/zentao_models"
	"wzDataCenter/common"
	"wzDataCenter/utils"
)

func GetLeixingInfo(ctx *gin.Context) {
	user, err := ctx.Get("claims")
	if !err {
		common.FailWithMsg("获取‘类型’数据失败", ctx)
	} else {
		userid := user.(*utils.CustomClaims).Users.ID
		succ, leixingList, _ := zentao_models.GetLeixing(userid)
		if succ {
			common.OkWithData(leixingList, ctx)
		} else {
			common.FailWithMsg("数据查询失败", ctx)
		}
	}
}
