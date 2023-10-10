package giteeGallery_controllers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
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

	// 创建发送客户端
	client := &http.Client{}
	// url
	uri := "https://gitee.com/api/v5/repos/" + giteeGallery_common.GITEEGALLERY_CONF.Gitee.Owner + "/" + giteeGallery_common.GITEEGALLERY_CONF.Gitee.Repo + "/contents/" + filename
	// 发送的数据
	data := url.Values{}
	data.Set("access_token", giteeGallery_common.GITEEGALLERY_CONF.Gitee.AccessToken)
	data.Set("content", picture)
	data.Set("message", message)
	data.Set("branch", giteeGallery_common.GITEEGALLERY_CONF.Gitee.Branch)
	// 执行请求
	resp, err := client.PostForm(uri, data)
	defer resp.Body.Close()
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}

	if !strings.HasPrefix(strconv.Itoa(resp.StatusCode), "20") {
		common.FailWithMsg("网络连接失败，请稍后再试", ctx)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	/*var v interface{}
	json.Unmarshal(body, &v)
	respData := v.(map[string]interface{})
	contentData := respData["content"].(map[string]interface{})

	common.OkWithData(contentData["download_url"], ctx)*/
	common.OkWithData(string(body), ctx)
}

// Get 获取仓库具体路径下的内容
func Get(ctx *gin.Context) {
	filename := ctx.DefaultQuery("filename", "")

	// 创建发送客户端
	client := &http.Client{}
	// url
	uri := "https://gitee.com/api/v5/repos/" + giteeGallery_common.GITEEGALLERY_CONF.Gitee.Owner + "/" + giteeGallery_common.GITEEGALLERY_CONF.Gitee.Repo + "/contents/" + filename
	// 发送的数据
	data := url.Values{}
	data.Set("access_token", giteeGallery_common.GITEEGALLERY_CONF.Gitee.AccessToken)
	payload := strings.NewReader(data.Encode())
	req, err := http.NewRequest("GET", uri, payload)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}
	// 执行请求
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}

	if !strings.HasPrefix(strconv.Itoa(resp.StatusCode), "20") {
		common.FailWithMsg("网络连接失败，请稍后再试", ctx)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	common.OkWithData(string(body), ctx)
}
