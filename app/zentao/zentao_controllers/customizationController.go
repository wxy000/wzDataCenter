package zentao_controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"wzDataCenter/app/zentao/zentao_models"
	"wzDataCenter/common"
	"wzDataCenter/utils"
)

// dayOfYear 函数接收一个整数，表示本年第几天，并返回对应的日期。
func dayOfYear(day int) time.Time {
	now := time.Now()
	start := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	return start.AddDate(0, 0, day-1)
}

// GetAnalysisLineCustomization 获取年度‘客制返工关系图’（预计时长）
func GetAnalysisLineCustomization(ctx *gin.Context) {
	user, _ := ctx.Get("claims")
	userId := user.(*utils.CustomClaims).Users.ID

	dateYear := ctx.DefaultQuery("dateYear", strconv.Itoa(time.Now().Year()))
	if dateYear == "" {
		dateYear = strconv.Itoa(time.Now().Year())
	}

	succ, customizationList, _ := zentao_models.GetAnalysisLineCustomization(userId, dateYear)

	endtime := dateYear + "-12-31"
	starttime := dateYear + "-01-01"
	e, _ := time.Parse("2006-01-02", endtime)
	s, _ := time.Parse("2006-01-02", starttime)
	day := e.Sub(s).Hours()/24 + 1

	//月份
	var dataDate []string
	for i := 0; i < int(day); i++ {
		dataDate = append(dataDate, dayOfYear((i + 1)).Format("01-02"))
	}

	//客制
	var dataKezhi []float64
	var flagKezhi = false
	var tmpKezhi = 0.0
	for i := 0; i < int(day); i++ {
		for j := 0; j < len(*customizationList); j++ {
			if (*customizationList)[j].Clid == "T0343" && (*customizationList)[j].Dayofyear == (i+1) {
				flagKezhi = true
				tmpKezhi = (*customizationList)[j].Esti
			}
		}
		if flagKezhi {
			dataKezhi = append(dataKezhi, tmpKezhi)
		} else {
			dataKezhi = append(dataKezhi, 0)
		}
		flagKezhi = false
		tmpKezhi = 0.0
	}

	//客制
	var dataFangong []float64
	var flagFangong = false
	var tmpFangong = 0.0
	for i := 0; i < int(day); i++ {
		for j := 0; j < len(*customizationList); j++ {
			if (*customizationList)[j].Clid == "T0344" && (*customizationList)[j].Dayofyear == (i+1) {
				flagFangong = true
				tmpFangong = (*customizationList)[j].Esti
			}
		}
		if flagFangong {
			dataFangong = append(dataFangong, tmpFangong)
		} else {
			dataFangong = append(dataFangong, 0)
		}
		flagFangong = false
		tmpFangong = 0.0
	}

	if succ {
		common.OkWithData(gin.H{
			"dataDate":    dataDate,
			"dataKezhi":   dataKezhi,
			"dataFangong": dataFangong,
		}, ctx)
	} else {
		common.FailWithMsg("数据查询失败", ctx)
	}
}
