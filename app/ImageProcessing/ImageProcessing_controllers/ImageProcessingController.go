package ImageProcessing_controllers

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/freetype"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
	"wzDataCenter/common"
)

// CreateWordsImg 输入文字生成图片
func CreateWordsImg(ctx *gin.Context) {
	words := ctx.DefaultQuery("words", "")
	color := ctx.DefaultQuery("color", "white")

	//定义一个包含所有ascii字符的字符串
	asciiChs := ""
	for i := 32; i < 127; i++ {
		asciiChs += fmt.Sprintf("%c", i)
	}
	//检查所有ascii字符的个数
	nohanzi := 0
	for _, ch := range words {
		if strings.ContainsRune(asciiChs, ch) {
			nohanzi += 1
		}
	}

	//高度
	wordsHeight := 40
	//宽度（ascii字符只占大概一半汉字的宽度）
	wordsLen := utf8.RuneCountInString(words)
	wordsWidth := (wordsLen-nohanzi/2)*wordsHeight + 1
	//创建画布
	img := image.NewRGBA(image.Rect(0, 0, wordsWidth, wordsHeight))

	//读字体
	fontBytes, err := ioutil.ReadFile("./app/ImageProcessing/fonts/康熙字典美化体.ttf")
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
	_, err = c.DrawString(words, freetype.Pt(0, wordsHeight-6))
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}
	//保存
	path, err := saveFile(img, "./app/ImageProcessing/tmp/words.png")
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}
	base64file, err := fileToBase64(path)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}
	common.OkWithData(base64file, ctx)
}

func saveFile(pic *image.RGBA, path string) (string, error) {
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

func base64ToFile(base64Str string, outputPath string) error {
	// 解码base64字符串
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return err
	}
	// 将解码的数据写入文件
	err = ioutil.WriteFile(outputPath, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
