package parser12306_controllers

import (
	"github.com/gin-gonic/gin"
	"regexp"
	"wzDataCenter/common"
)

// ParserTicketCalendar 解析车票信息（日历）
func ParserTicketCalendar(ctx *gin.Context) {
	name := ctx.PostForm("name")
	rawData := ctx.PostForm("rawData")

	basePattern := "\\d\\." + name + "，\\d{4}年(.*)检票口([0-9a-zA-Z]*)(，电子客票。|。)"
	timePattern := "(\\d{4})年(\\d{1,2})月(\\d{1,2})日(\\d{2}:\\d{2})开"
	sitePattern := "([\u4e00-\u9fa5]*)[—-]([\u4e00-\u9fa5]*)"
	trainNumberPattern := "([0-9a-zA-Z]*)次列车"
	seatPattern := "([0-9a-zA-Z]*)车([0-9a-zA-Z]*)号"
	ticketEntrancePattern := "检票口([0-9a-zA-Z]*)"

	// 取到有用的那条字符串
	baseRegex := regexp.MustCompile(basePattern)
	baseStr := baseRegex.FindString(rawData)

	// 获取时间
	timeRegex := regexp.MustCompile(timePattern)
	timeArr := timeRegex.FindStringSubmatch(baseStr)

	// 获取站点
	siteRegex := regexp.MustCompile(sitePattern)
	siteArr := siteRegex.FindStringSubmatch(baseStr)

	// 获取车次
	trainNumberRegex := regexp.MustCompile(trainNumberPattern)
	trainNumberArr := trainNumberRegex.FindStringSubmatch(baseStr)

	// 获取座位
	seatRegex := regexp.MustCompile(seatPattern)
	seatArr := seatRegex.FindStringSubmatch(baseStr)

	// 检票口
	ticketEntranceRegex := regexp.MustCompile(ticketEntrancePattern)
	ticketEntranceArr := ticketEntranceRegex.FindStringSubmatch(baseStr)

	// fmt.Println("提取字符串内容：", baseStr)

	common.OkWithData(gin.H{
		"time":           timeArr[0],
		"year":           timeArr[1],
		"month":          timeArr[2],
		"day":            timeArr[3],
		"hour":           timeArr[4],
		"site":           siteArr[0],
		"siteStart":      siteArr[1],
		"siteEnd":        siteArr[2],
		"trainNumber":    trainNumberArr[1],
		"carNumber":      seatArr[1],
		"seatNumber":     seatArr[2],
		"ticketEntrance": ticketEntranceArr[1],
	}, ctx)
}