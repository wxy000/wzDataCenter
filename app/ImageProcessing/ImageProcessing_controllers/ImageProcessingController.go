package ImageProcessing_controllers

import (
	"encoding/base64"
	"errors"
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

	fontPath := "./app/ImageProcessing/fonts/方正硬笔行书繁体.ttf"
	outputPath := "./app/ImageProcessing/tmp/words.png"

	base64file, err := words2Img(words, color, fontPath, outputPath)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
		return
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
		return
	}
	if words == "" && idio == "" {
		common.FailWithMsg("拍摄地址没有，签名也没有，还整啥水印啊！？", ctx)
		return
	}

	//创建一个原图大小的透明图片
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	//只存在签名
	if (words == "" && len(words) == 0) && (idio != "" && len(idio) != 0) {
		//将base64转为图片并暂存
		idioPath := "./app/ImageProcessing/tmp/idio.png"
		err := base64ToFile(idio, idioPath)
		if err != nil {
			common.FailWithMsg(err.Error(), ctx)
			return
		}

		//缩放图片
		newImgIdioDx := width / 6
		if width > height {
			newImgIdioDx = width / 6
		} else {
			newImgIdioDx = width / 4
		}
		newImgIdio, err := resizeImg(idioPath, newImgIdioDx, 0)
		if err != nil {
			common.FailWithMsg(err.Error(), ctx)
			return
		}

		if width > height {
			draw.Draw(img, img.Bounds().Add(image.Point{X: (width - newImgIdioDx) / 2, Y: height - int(float64(newImgIdio.Bounds().Dy())*1.5)}), newImgIdio, image.Point{}, draw.Src)
		} else {
			draw.Draw(img, img.Bounds().Add(image.Point{X: (width - newImgIdioDx) / 2, Y: height - int(float64(newImgIdio.Bounds().Dy())*3)}), newImgIdio, image.Point{}, draw.Src)
		}
	}

	//只存在地址
	if (words != "" && len(words) != 0) && (idio == "" && len(idio) == 0) {
		fontPath := "./app/ImageProcessing/fonts/方正硬笔行书繁体.ttf"
		wordsPath := "./app/ImageProcessing/tmp/words.png"

		//将传入的文字转为图片
		_, err := words2Img(words, color, fontPath, wordsPath)
		if err != nil {
			common.FailWithMsg(err.Error(), ctx)
			return
		}

		//缩放图片
		newImgWordsDx := width / 6
		if width > height {
			newImgWordsDx = width / 6
		} else {
			newImgWordsDx = width / 4
		}
		newImgWords, err := resizeImg(wordsPath, newImgWordsDx, 0)
		if err != nil {
			common.FailWithMsg(err.Error(), ctx)
			return
		}

		if width > height {
			draw.Draw(img, img.Bounds().Add(image.Point{X: (width - newImgWordsDx) / 2, Y: height - int(float64(newImgWords.Bounds().Dy())*1.5)}), newImgWords, image.Point{}, draw.Src)
		} else {
			draw.Draw(img, img.Bounds().Add(image.Point{X: (width - newImgWordsDx) / 2, Y: height - int(float64(newImgWords.Bounds().Dy())*3)}), newImgWords, image.Point{}, draw.Src)
		}
	}

	//签名、地址都存在
	if (words != "" && len(words) != 0) && (idio != "" && len(idio) != 0) {
		/*****************签名*****************/
		//将base64转为图片并暂存
		idioPath := "./app/ImageProcessing/tmp/idio.png"
		err := base64ToFile(idio, idioPath)
		if err != nil {
			common.FailWithMsg(err.Error(), ctx)
			return
		}
		//缩放图片
		newImgIdioDx := width / 6
		if width > height {
			newImgIdioDx = width / 6
		} else {
			newImgIdioDx = width / 4
		}
		newImgIdio, err := resizeImg(idioPath, newImgIdioDx, 0)
		if err != nil {
			common.FailWithMsg(err.Error(), ctx)
			return
		}
		/*****************签名*****************/
		/*****************地址*****************/
		fontPath := "./app/ImageProcessing/fonts/方正硬笔行书繁体.ttf"
		wordsPath := "./app/ImageProcessing/tmp/words.png"
		//将传入的文字转为图片
		_, err = words2Img(words, color, fontPath, wordsPath)
		if err != nil {
			common.FailWithMsg(err.Error(), ctx)
			return
		}
		newImgWords, err := resizeImg(wordsPath, 0, newImgIdio.Bounds().Dy())
		if err != nil {
			common.FailWithMsg(err.Error(), ctx)
			return
		}
		/*****************地址*****************/
		//合成到背景上
		//定位
		if rules == "v" {
			if width > height {
				draw.Draw(img, img.Bounds().Add(image.Point{X: (width - newImgIdio.Bounds().Dx()) / 2, Y: height - int(float64(newImgIdio.Bounds().Dy())*1.5)}), newImgIdio, image.Point{}, draw.Src)
				draw.Draw(img, img.Bounds().Add(image.Point{X: (width - newImgWords.Bounds().Dx()) / 2, Y: height - int(float64(newImgIdio.Bounds().Dy())*1.5+float64(newImgWords.Bounds().Dy()))}), newImgWords, image.Point{}, draw.Src)
			} else {
				draw.Draw(img, img.Bounds().Add(image.Point{X: (width - newImgIdio.Bounds().Dx()) / 2, Y: height - newImgIdio.Bounds().Dy()*3}), newImgIdio, image.Point{}, draw.Src)
				draw.Draw(img, img.Bounds().Add(image.Point{X: (width - newImgWords.Bounds().Dx()) / 2, Y: height - newImgIdio.Bounds().Dy()*3 - newImgWords.Bounds().Dy()}), newImgWords, image.Point{}, draw.Src)
			}
		} else {
			x := (width - newImgIdio.Bounds().Dx() - newImgWords.Bounds().Dx()) / 2
			y := float64(height) - float64(newImgIdio.Bounds().Dy())*1.5
			if width > height {
				y = float64(height) - float64(newImgIdio.Bounds().Dy())*1.5
			} else {
				y = float64(height) - float64(newImgIdio.Bounds().Dy())*3
			}
			draw.Draw(img, img.Bounds().Add(image.Point{X: x, Y: int(y)}), newImgWords, image.Point{}, draw.Src)
			draw.Draw(img, img.Bounds().Add(image.Point{X: x + newImgWords.Bounds().Dx() + 10, Y: int(y)}), newImgIdio, image.Point{}, draw.Src)
		}
	}

	//保存
	path, err := saveFile(img, "./app/ImageProcessing/tmp/waterMark.png")
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
		return
	}
	base64file, err := fileToBase64(path)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
		return
	}

	common.OkWithData(base64file, ctx)
}

// 调整图片大小
func resizeImg(imgPath string, afterWidth int, afterHeight int) (image.Image, error) {
	if afterWidth == 0 && afterHeight == 0 {
		return nil, errors.New("长宽都是0，那不没了吗？还调个屁啊")
	}

	//加载图片
	imgOpen, err := os.Open(imgPath)
	if err != nil {
		return nil, err
	}
	imgDecode, err := png.Decode(imgOpen)
	if err != nil {
		return nil, err
	}

	//获取原图大小
	imgDx := imgDecode.Bounds().Dx()
	imgDy := imgDecode.Bounds().Dy()

	imgDxNew := afterWidth
	imgDyNew := afterHeight

	//如果只传了width或者height，则按比例调整
	if afterWidth != 0 && afterHeight == 0 {
		imgDyNew = int(float64(imgDy) / float64(imgDx) * float64(afterWidth))
	}
	if afterWidth == 0 && afterHeight != 0 {
		imgDxNew = int(float64(imgDx) / float64(imgDy) * float64(afterHeight))
	}

	return resize.Resize(uint(imgDxNew), uint(imgDyNew), imgDecode, resize.Bilinear), nil
}

// 传入文字、颜色，生成透明图片，并返回base64串
func words2Img(words string, color string, fontPath string, outputPath string) (string, error) {
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
	fontBytes, err := ioutil.ReadFile(fontPath)
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
	path, err := saveFile(img, outputPath)
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
