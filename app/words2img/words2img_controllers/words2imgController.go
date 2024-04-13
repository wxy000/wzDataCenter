package words2img_controllers

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/golang/freetype"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"unicode/utf8"
	"wzDataCenter/common"
)

func CreateImg(ctx *gin.Context) {
	words := ctx.DefaultQuery("words", "")
	color := ctx.DefaultQuery("color", "white")

	//高度
	wordsHeight := 40
	//宽度
	wordsLen := utf8.RuneCountInString(words)
	wordsWidth := wordsLen*wordsHeight + 1
	//创建画布
	img := image.NewRGBA(image.Rect(0, 0, wordsWidth, wordsHeight))

	//读字体
	fontBytes, err := ioutil.ReadFile("./app/words2img/fonts/康熙字典美化体.ttf")
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(float64(wordsHeight))
	c.SetClip(img.Bounds())
	c.SetDst(img)
	if color == "white" {
		c.SetSrc(image.White)
	} else {
		c.SetSrc(image.Black)
	}
	//设置字体显示位置
	_, err = c.DrawString(words, freetype.Pt(0, wordsHeight-4))
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}
	//保存
	path, err := saveFile(img)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}
	base64file, err := fileToBase64(path)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}
	common.OkWithData(base64file, ctx)
}

func saveFile(pic *image.RGBA) (string, error) {
	path := "./app/words2img/tmp/words.png"
	dstFile, _ := os.Create(path)
	defer dstFile.Close()

	err := png.Encode(dstFile, pic)
	if err != nil {
		return "", err
	}

	return path, nil
}

func fileToBase64(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	base64Str := base64.StdEncoding.EncodeToString(data)
	return base64Str, nil
}
