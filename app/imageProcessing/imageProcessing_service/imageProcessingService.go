package imageProcessing_service

import (
	"errors"
	"github.com/golang/freetype"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"unicode/utf8"
	"wzDataCenter/utils"
)

// SetLittleWaterMarkToBG 将小水印打在背景上
// position：1.左上，2.上中，3.右上，4.左中，5.中，6.右中，7.左下，8.下中，9.右下
// 优先position定位，如果position为0，则用x，y
func SetLittleWaterMarkToBG(bgImg *image.RGBA, waterMark image.Image, marginRadio float64, position string, x int, y int) {
	//背景的宽高
	width := bgImg.Bounds().Dx()
	height := bgImg.Bounds().Dy()
	//遮罩margin的宽度（向内的边框宽度）
	marginWidth := float64(width) * marginRadio
	marginHeight := float64(height) * marginRadio
	//小水印的宽高
	waterMarkDx := waterMark.Bounds().Dx()
	waterMarkDy := waterMark.Bounds().Dy()

	//定位
	if position == "1" {
		draw.Draw(bgImg, bgImg.Bounds().Add(image.Point{X: int(marginWidth), Y: int(marginHeight)}), waterMark, image.Point{}, draw.Src)
	}
	if position == "2" {
		draw.Draw(bgImg, bgImg.Bounds().Add(image.Point{X: (width - waterMarkDx) / 2, Y: int(marginHeight)}), waterMark, image.Point{}, draw.Src)
	}
	if position == "3" {
		draw.Draw(bgImg, bgImg.Bounds().Add(image.Point{X: width - waterMarkDx - int(marginWidth), Y: int(marginHeight)}), waterMark, image.Point{}, draw.Src)
	}
	if position == "4" {
		draw.Draw(bgImg, bgImg.Bounds().Add(image.Point{X: int(marginWidth), Y: (height - waterMarkDy) / 2}), waterMark, image.Point{}, draw.Src)
	}
	if position == "5" {
		draw.Draw(bgImg, bgImg.Bounds().Add(image.Point{X: (width - waterMarkDx) / 2, Y: (height - waterMarkDy) / 2}), waterMark, image.Point{}, draw.Src)
	}
	if position == "6" {
		draw.Draw(bgImg, bgImg.Bounds().Add(image.Point{X: width - waterMarkDx - int(marginWidth), Y: (height - waterMarkDy) / 2}), waterMark, image.Point{}, draw.Src)
	}
	if position == "7" {
		draw.Draw(bgImg, bgImg.Bounds().Add(image.Point{X: int(marginWidth), Y: height - waterMarkDy - int(marginHeight)}), waterMark, image.Point{}, draw.Src)
	}
	if position == "8" {
		draw.Draw(bgImg, bgImg.Bounds().Add(image.Point{X: (width - waterMarkDx) / 2, Y: height - waterMarkDy - int(marginHeight)}), waterMark, image.Point{}, draw.Src)
	}
	if position == "9" {
		draw.Draw(bgImg, bgImg.Bounds().Add(image.Point{X: width - waterMarkDx - int(marginWidth), Y: height - waterMarkDy - int(marginHeight)}), waterMark, image.Point{}, draw.Src)
	}
	if position == "0" {
		draw.Draw(bgImg, bgImg.Bounds().Add(image.Point{X: x, Y: y}), waterMark, image.Point{}, draw.Src)
	}
}

// ResizeImg 调整图片大小
func ResizeImg(imgPath string, afterWidth int, afterHeight int) (image.Image, error) {
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

// Words2Img 传入文字、颜色，生成透明图片，并返回base64串
func Words2Img(words string, color string, rules string, fontPath string, outputPath string) (string, error) {

	if rules == "" || len(rules) == 0 {
		rules = "H"
	}
	if rules != "H" && rules != "V" {
		rules = "H"
	}

	//读字体
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return "", err
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return "", err
	}

	//多少个字
	wordsLen := utf8.RuneCountInString(words)
	wordsHeight := 0
	wordsWidth := 0
	if rules == "H" {
		//高度
		wordsHeight = 65
		//宽度
		wordsWidth = wordsLen*wordsHeight + 1
	} else if rules == "V" {
		wordsWidth = 65
		wordsHeight = wordsLen*wordsWidth + 1
	}

	//创建画布
	img := image.NewRGBA(image.Rect(0, 0, wordsWidth, wordsHeight))

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	if rules == "H" {
		c.SetFontSize(float64(wordsHeight))
	} else if rules == "V" {
		c.SetFontSize(float64(wordsWidth))
	}
	c.SetClip(img.Bounds())
	c.SetDst(img)
	if color == "white" {
		c.SetSrc(image.White)
	} else {
		c.SetSrc(image.Black)
	}

	//设置字体显示位置
	if rules == "H" {
		_, err = c.DrawString(words, freetype.Pt(0, wordsHeight-6))
		if err != nil {
			return "", err
		}
	} else if rules == "V" {
		i := 0
		for _, word := range words {
			_, err = c.DrawString(string(word), freetype.Pt(0, wordsWidth*(i+1)-6))
			if err != nil {
				return "", err
			}
			i += 1
		}
	}

	//切掉透明的部分
	bounds := img.Bounds()
	x1 := 0
	y1 := 0
	x2 := 0
	y2 := 0
outloop1:
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a > 0 {
				x1 = x
				break outloop1
			}
		}
	}
outloop2:
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a > 0 {
				y1 = y
				break outloop2
			}
		}
	}
outloop3:
	for x := bounds.Max.X; x > bounds.Min.X; x-- {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a > 0 {
				x2 = x + 1
				break outloop3
			}
		}
	}
outloop4:
	for y := bounds.Max.Y; y > bounds.Min.Y; y-- {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a > 0 {
				y2 = y + 1
				break outloop4
			}
		}
	}

	_ = img.SubImage(image.Rect(x1, y1, x2, y2)).(*image.RGBA)

	//保存
	path, err := SaveFile(img, outputPath)
	if err != nil {
		return "", err
	}
	base64file, err := utils.FileToBase64(path)
	if err != nil {
		return "", err
	}
	return base64file, nil
}

func SaveFile(pic *image.RGBA, path string) (string, error) {
	dstFile, _ := os.Create(path)
	defer dstFile.Close()

	err := png.Encode(dstFile, pic)
	if err != nil {
		return "", err
	}

	return path, nil
}
