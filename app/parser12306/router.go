package parser12306

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/app/parser12306/parser12306_controllers"
)

func Router(r *gin.Engine) {
	parser12306 := r.Group("/12306")
	{
		parser12306.POST("/parserTicketCalendar", parser12306_controllers.ParserTicketCalendar)
	}
}
