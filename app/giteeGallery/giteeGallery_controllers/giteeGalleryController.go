package giteeGallery_controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/unti-io/go-utils/utils"
	"strconv"
	"strings"
	"time"
	"wzDataCenter/app/giteeGallery/giteeGallery_common"
	"wzDataCenter/common"
)

// Update 新建文件
func Update(ctx *gin.Context) {
	picture := ctx.PostForm("picture")
	suffix := ctx.PostForm("suffix")

	currentTime := time.Now()
	msgTime := currentTime.Format("2006-01-02 15:04:05")

	filename := currentTime.Format("20060102150405")

	message := "Update " + filename + "." + suffix + " by 快捷指令 - " + msgTime

	// url
	url := "https://gitee.com/api/v5/repos/" + giteeGallery_common.GITEEGALLERY_CONF.Gitee.Owner + "/" + giteeGallery_common.GITEEGALLERY_CONF.Gitee.Repo + "/contents/" + filename

	curl := utils.Curl(utils.CurlRequest{
		Method: "POST",
		Url:    url,
		Body: map[string]any{
			"access_token": giteeGallery_common.GITEEGALLERY_CONF.Gitee.AccessToken,
			"content":      picture,
			"message":      message,
			"branch":       giteeGallery_common.GITEEGALLERY_CONF.Gitee.Branch,
		},
	}).Send()

	if !strings.HasPrefix(strconv.Itoa(curl.StatusCode), "20") {
		common.FailWithMsg("网络连接失败，请稍后再试", ctx)
	} else {
		common.OkWithData(curl.Text, ctx)
	}
}

// Get 获取仓库具体路径下的内容
func Get(ctx *gin.Context) {
	filename := ctx.DefaultQuery("filename", "")

	url := "https://gitee.com/api/v5/repos/" + giteeGallery_common.GITEEGALLERY_CONF.Gitee.Owner + "/" + giteeGallery_common.GITEEGALLERY_CONF.Gitee.Repo + "/contents/" + filename

	curl := utils.Curl(utils.CurlRequest{
		Method: "GET",
		Url:    url,
		Query: map[string]any{
			"access_token": giteeGallery_common.GITEEGALLERY_CONF.Gitee.AccessToken,
		},
	}).Send()

	if !strings.HasPrefix(strconv.Itoa(curl.StatusCode), "20") {
		common.FailWithMsg("网络连接失败，请稍后再试", ctx)
	} else {
		common.OkWithData(curl.Text, ctx)
	}
}

// Del 删除图片
func Del(ctx *gin.Context) {
	sha := ctx.PostForm("sha")
	path := ctx.PostForm("path")

	currentTime := time.Now()
	msgTime := currentTime.Format("2006-01-02 15:04:05")

	message := "Update " + path + " by 快捷指令 - " + msgTime

	// url
	url := "https://gitee.com/api/v5/repos/" + giteeGallery_common.GITEEGALLERY_CONF.Gitee.Owner + "/" + giteeGallery_common.GITEEGALLERY_CONF.Gitee.Repo + "/contents/" + path

	curl := utils.Curl(utils.CurlRequest{
		Method: "DELETE",
		Url:    url,
		Query: map[string]any{
			"access_token": giteeGallery_common.GITEEGALLERY_CONF.Gitee.AccessToken,
			"sha":          sha,
			"message":      message,
		},
	}).Send()

	if !strings.HasPrefix(strconv.Itoa(curl.StatusCode), "20") {
		common.FailWithMsg("网络连接失败，请稍后再试", ctx)
	} else {
		common.OkWithData(curl.Text, ctx)
	}
}
