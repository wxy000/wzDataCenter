package ImageProcessing_controllers

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/freetype"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
	"wzDataCenter/common"
)

// CreateWordsImg 输入文字生成图片
func CreateWordsImg(ctx *gin.Context) {
	words := ctx.DefaultQuery("words", "")
	color := ctx.DefaultQuery("color", "white")

	base64file, err := words2Img(words, color)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}
	common.OkWithData(base64file, ctx)
}

// CreateImgWaterMarkWithWordsAndIdio 创建水印
func CreateImgWaterMarkWithWordsAndIdio(ctx *gin.Context) {
	widthtmp := ctx.DefaultPostForm("width", "0")
	heighttmp := ctx.DefaultPostForm("height", "0")
	words := ctx.DefaultPostForm("words", "")
	color := ctx.DefaultPostForm("color", "white")
	idio := ctx.DefaultPostForm("idio", "")
	rules := ctx.DefaultPostForm("rules", "h")
	width, _ := strconv.Atoi(widthtmp)
	height, _ := strconv.Atoi(heighttmp)
	if width == 0 || height == 0 {
		common.FailWithMsg("请传入正确的图片长宽！", ctx)
	}

	//将传入的文字转为图片
	base64fileWords, err := words2Img(words, color)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}

	//将base64转为图片并暂存
	idioPath := "./app/ImageProcessing/tmp/idio.png"
	err = base64ToFile(idio, idioPath)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}
	wordsPath := "./app/ImageProcessing/tmp/words.png"
	err = base64ToFile(base64fileWords, wordsPath)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}

	//加载图片
	imgIdioOpen, err := os.Open(idioPath)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}
	imgIdio, err := png.Decode(imgIdioOpen)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}

	imgWordsOpen, err := os.Open(wordsPath)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}
	imgWords, err := png.Decode(imgWordsOpen)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}

	//创建一个原图大小的透明图片
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	//按比例缩小idio和words
	imgIdioDx := imgIdio.Bounds().Dx()
	imgIdioDy := imgIdio.Bounds().Dy()
	newImgIdioDx := width / 9
	newImgIdioDy := float64(imgIdioDy) / float64(imgIdioDx) * float64(newImgIdioDx)
	/*newImgIdio := image.NewRGBA(image.Rect(0, 0, newImgIdioDx, int(newImgIdioDy)))
	draw.Draw(newImgIdio, newImgIdio.Bounds(), imgIdio, image.Point{}, draw.Src)*/
	newImgIdio := resize.Resize(uint(newImgIdioDx), uint(newImgIdioDy), imgIdio, resize.Bilinear)

	imgWordsDx := imgWords.Bounds().Dx()
	imgWordsDy := imgWords.Bounds().Dy()
	newImgWordsDy := newImgIdioDy
	newImgWordsDx := float64(imgWordsDx) / float64(imgWordsDy) * float64(newImgWordsDy)
	/*newImgWords := image.NewRGBA(image.Rect(0, 0, int(newImgWordsDx), int(newImgWordsDy)))
	draw.Draw(newImgWords, newImgWords.Bounds(), imgWords, image.Point{}, draw.Src)*/
	newImgWords := resize.Resize(uint(newImgWordsDx), uint(newImgWordsDy), imgWords, resize.Bilinear)

	//合成到背景上
	//定位
	if rules == "v" {
		draw.Draw(img, img.Bounds().Add(image.Point{X: (width - newImgIdioDx) / 2, Y: height - int(newImgIdioDy*1.5)}), newImgIdio, image.Point{}, draw.Src)
		draw.Draw(img, img.Bounds().Add(image.Point{X: (width - int(newImgWordsDx)) / 2, Y: height - int(newImgIdioDy*1.5+newImgWordsDy)}), newImgWords, image.Point{}, draw.Src)
	} else {
		x := (float64(width) - float64(newImgIdioDx) - newImgWordsDx) / 2
		y := float64(height) - newImgIdioDy*1.5
		draw.Draw(img, img.Bounds().Add(image.Point{X: int(x), Y: int(y)}), newImgWords, image.Point{}, draw.Src)
		draw.Draw(img, img.Bounds().Add(image.Point{X: int(x + newImgWordsDx + 10), Y: int(y)}), newImgIdio, image.Point{}, draw.Src)
	}

	//保存
	path, err := saveFile(img, "./app/ImageProcessing/tmp/waterMark.png")
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}
	base64file, err := fileToBase64(path)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}

	common.OkWithData(base64file, ctx)
}

// 传入文字、颜色，生成透明图片，并返回base64串
func words2Img(words string, color string) (string, error) {
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
		return "", err
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return "", err
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
		return "", err
	}
	//保存
	path, err := saveFile(img, "./app/ImageProcessing/tmp/words.png")
	if err != nil {
		return "", err
	}
	base64file, err := fileToBase64(path)
	if err != nil {
		return "", err
	}
	return base64file, nil
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
	err = ioutil.WriteFile(outputPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
