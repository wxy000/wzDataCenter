package zentao_controllers

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/app/zentao/zentao_models"
	"wzDataCenter/common"
	"wzDataCenter/utils"
)

// GetAnalysisLeixing 获取‘类型’数据
/*
	p1:userId-用户id-必填项(直接通过token获取，无需传入)
	p2:dateStart-起始时间-非必填
	p3:dateEnd-截止时间-非必填
	p4:type0-数据格式类型-0.原始数据格式，1.预计实际都有，2.预计，3.实际
*/
func GetAnalysisLeixing(ctx *gin.Context) {
	user, _ := ctx.Get("claims")
	userId := user.(*utils.CustomClaims).Users.ID

	dateStart := ctx.DefaultQuery("dateStart", "1900-01-01")
	if dateStart == "" {
		dateStart = "1900-01-01"
	}
	dateStart = dateStart + " 00:00:00"
	dateEnd := ctx.DefaultQuery("dateEnd", "3000-12-31")
	if dateEnd == "" {
		dateEnd = "3000-12-31"
	}
	dateEnd = dateEnd + " 23:59:59"
	type0 := ctx.DefaultQuery("type0", "0")

	succ, leixingList, _ := zentao_models.GetAnalysisLeixing(userId, dateStart, dateEnd)
	if succ {
		var result interface{}
		result = leixingList
		// 原始
		if type0 == "0" {
			result = leixingList
		}
		// 都有
		if type0 == "1" {
			result = getLeixingYS(*leixingList)
		}
		// 预计
		if type0 == "2" {
			result = getLeixingY(*leixingList)
		}
		// 实际
		if type0 == "3" {
			result = getLeixingS(*leixingList)
		}
		common.OkWithData(result, ctx)
	} else {
		common.FailWithMsg("数据查询失败", ctx)
	}
}

// getLeixingYS 按‘类型’分组，预计实际都有
/*
	p1:list-‘类型’列表
	description：格式样例-
		{
			"type1": ["BUG-其他程序问题","产品BUG","需求评估","客制需求(转个案)","客制需求(线上免费)"],
			"yuji": [28,13,2,44,54],
			"shiji": [28,13,2,31,54]
		}
*/
func getLeixingYS(list []zentao_models.Leixing) interface{} {
	var type1 []string
	var yuji []float64
	var shiji []float64
	for i := 0; i < len(list); i++ {
		type1 = append(type1, list[i].Cloudname)
		yuji = append(yuji, list[i].Esti)
		shiji = append(shiji, list[i].Cons)
	}
	return gin.H{
		"type1": type1,
		"yuji":  yuji,
		"shiji": shiji,
	}
}

// getLeixingY 按‘类型’分组，预计时长
/*
	p1:list-‘类型’列表
	description：格式样例-
		[
			{"name": "BUG-其他程序问题","value": 28},
			{"name": "产品BUG","value": 13},
			{"name": "需求评估","value": 2},
			{"name": "客制需求(转个案)","value": 44},
			{"name": "客制需求(线上免费)","value": 54}
		]
*/
func getLeixingY(list []zentao_models.Leixing) interface{} {
	var rlist []interface{}
	for i := 0; i < len(list); i++ {
		rlist = append(rlist, map[string]interface{}{"name": list[i].Cloudname, "value": list[i].Esti})
	}
	return rlist
}

// getLeixingS 按‘类型’分组，实际时长
/*
	p1:list-‘类型’列表
	description：格式样例-
		[
			{"name": "BUG-其他程序问题","value": 28},
			{"name": "产品BUG","value": 13},
			{"name": "需求评估","value": 2},
			{"name": "客制需求(转个案)","value": 31},
			{"name": "客制需求(线上免费)","value": 54}
		]
*/
func getLeixingS(list []zentao_models.Leixing) interface{} {
	var rlist []interface{}
	for i := 0; i < len(list); i++ {
		rlist = append(rlist, map[string]interface{}{"name": list[i].Cloudname, "value": list[i].Cons})
	}
	return rlist
}
