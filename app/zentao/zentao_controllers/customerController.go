package zentao_controllers

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/app/zentao/zentao_models"
	"wzDataCenter/common"
	"wzDataCenter/utils"
)

// GetAnalysisCustomer 获取‘类型’数据
/*
	p1:userId-用户id-必填项(直接通过token获取，无需传入)
	p2:dateStart-起始时间-非必填
	p3:dateEnd-截止时间-非必填
	p4:type0-数据格式类型-0.原始数据格式，1.预计实际都有，2.预计，3.实际
*/
func GetAnalysisCustomer(ctx *gin.Context) {
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

	succ, customerList, _ := zentao_models.GetAnalysisCustomer(userId, dateStart, dateEnd)
	if succ {
		var result interface{}
		result = customerList
		// 原始
		if type0 == "0" {
			result = customerList
		}
		// 都有
		if type0 == "1" {
			result = getCustomerYS(*customerList)
		}
		// 预计
		if type0 == "2" {
			result = getCustomerY(*customerList)
		}
		// 实际
		if type0 == "3" {
			result = getCustomerS(*customerList)
		}
		common.OkWithData(result, ctx)
	} else {
		common.FailWithMsg("数据查询失败", ctx)
	}
}

// getCustomerYS 按‘类型’分组，预计实际都有
/*
	p1:list-‘类型’列表
	description：格式样例-
		{
			"type1": ["x公司","y公司","z公司"],
			"yuji": [28,13,2],
			"shiji": [28,13,2]
		}
*/
func getCustomerYS(list []zentao_models.Customer) interface{} {
	var type1 []string
	var yuji []float64
	var shiji []float64
	for i := 0; i < len(list); i++ {
		type1 = append(type1, list[i].Customername)
		yuji = append(yuji, list[i].Esti)
		shiji = append(shiji, list[i].Cons)
	}
	return gin.H{
		"type1": type1,
		"yuji":  yuji,
		"shiji": shiji,
	}
}

// getCustomerY 按‘类型’分组，预计时长
/*
	p1:list-‘类型’列表
	description：格式样例-
		[
			{"name": "x公司","value": 28},
			{"name": "y公司","value": 13},
			{"name": "z公司","value": 2}
		]
*/
func getCustomerY(list []zentao_models.Customer) interface{} {
	var rlist []interface{}
	for i := 0; i < len(list); i++ {
		rlist = append(rlist, map[string]interface{}{"name": list[i].Customername, "value": list[i].Esti})
	}
	return rlist
}

// getCustomerS 按‘类型’分组，实际时长
/*
	p1:list-‘类型’列表
	description：格式样例-
		[
			{"name": "x公司","value": 28},
			{"name": "y公司","value": 13},
			{"name": "z公司","value": 2}
		]
*/
func getCustomerS(list []zentao_models.Customer) interface{} {
	var rlist []interface{}
	for i := 0; i < len(list); i++ {
		rlist = append(rlist, map[string]interface{}{"name": list[i].Customername, "value": list[i].Cons})
	}
	return rlist
}

// GetAnalysisLeixingDetail 按照类型取客户明细
func GetAnalysisLeixingDetail(ctx *gin.Context) {
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
	project := ctx.Query("project")

	succ, leixingList, count := zentao_models.GetAnalysisCustomerDetail(userId, project, dateStart, dateEnd)
	if succ {
		common.OkWithDataC(count, leixingList, ctx)
	} else {
		common.FailC(count, ctx)
	}
}
