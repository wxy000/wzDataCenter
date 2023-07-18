package controllers_web

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"wzDataCenter/common"
)

// LayoutMenu 主页菜单
func LayoutMenu(ctx *gin.Context) {
	layoutMenu, err := json.Marshal(common.CONF)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "index.html",
			common.Response{
				Code: http.StatusInternalServerError,
				Msg:  "获取菜单失败，请重试",
				Data: nil,
			})
	} else {
		var v interface{}
		_ = json.Unmarshal(layoutMenu, &v)
		ctx.HTML(http.StatusOK, "index.html",
			common.Response{
				Code: http.StatusOK,
				Msg:  "菜单获取成功",
				Data: v,
			})
	}
}
