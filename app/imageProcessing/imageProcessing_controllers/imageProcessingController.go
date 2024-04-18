package imageProcessing_controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"image"
	"strconv"
	"wzDataCenter/app/imageProcessing/imageProcessing_service"
	"wzDataCenter/common"
)

type typeWaterMarkStruct struct {
	Width          int      `json:"width"`
	Height         int      `json:"height"`
	Color          string   `json:"color"`
	Rules          string   `json:"rules"`
	Size           int      `json:"size"`
	Position       string   `json:"position"`
	X              int      `json:"x"`
	Y              int      `json:"y"`
	WaterMarkFiles []string `json:"waterMarkFiles"`
}

// CreateWordsImg 输入文字生成图片
func CreateWordsImg(ctx *gin.Context) {
	words := ctx.DefaultQuery("words", "")
	color := ctx.DefaultQuery("color", "white")

	fontPath := "./app/imageProcessing/fonts/方正硬笔行书繁体.ttf"
	outputPath := "./app/imageProcessing/tmp/words.png"

	base64file, err := imageProcessing_service.Words2Img(words, color, fontPath, outputPath)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
		return
	}
	common.OkWithData(base64file, ctx)
}

// CreateImgWaterMarkWithWords 创建水印
func CreateImgWaterMarkWithWords(ctx *gin.Context) {
	widthtmp := ctx.DefaultPostForm("width", "0")    //原图宽度
	heighttmp := ctx.DefaultPostForm("height", "0")  //原图高度
	words := ctx.DefaultPostForm("words", "")        //文字
	color := ctx.DefaultPostForm("color", "white")   //颜色（黑白）
	sizetmp := ctx.DefaultPostForm("size", "0")      //水印大小
	position := ctx.DefaultPostForm("position", "8") //水印在原图的定位
	xtmp := ctx.DefaultPostForm("x", "0")            //水印x轴
	ytmp := ctx.DefaultPostForm("y", "0")            //水印y轴
	width, _ := strconv.Atoi(widthtmp)
	height, _ := strconv.Atoi(heighttmp)
	size, _ := strconv.Atoi(sizetmp)
	x, _ := strconv.Atoi(xtmp)
	y, _ := strconv.Atoi(ytmp)

	if width == 0 || height == 0 {
		common.FailWithMsg("请传入正确的图片长宽！", ctx)
		return
	}
	if words == "" {
		common.FailWithMsg("拍摄地址没有，还整啥水印啊！？", ctx)
		return
	}

	fontPath := "./app/imageProcessing/fonts/方正硬笔行书繁体.ttf"
	wordsPath := "./app/imageProcessing/tmp/words.png"
	waterMarkPath := "./app/imageProcessing/tmp/waterMark.png"

	//将传入的文字转为图片
	_, err := imageProcessing_service.Words2Img(words, color, fontPath, wordsPath)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
		return
	}

	//缩放图片
	length := 0
	if size == 0 {
		if width > height {
			length = height / 20
		} else {
			length = width / 20
		}
	} else {
		length = size
	}
	newImgWords, err := imageProcessing_service.ResizeImg(wordsPath, 0, length)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
		return
	}

	//创建一个原图大小的透明图片
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	//定位
	imageProcessing_service.SetLittleWaterMarkToBG(img, newImgWords, 0.06, position, x, y)

	//保存
	path, err := imageProcessing_service.SaveFile(img, waterMarkPath)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
		return
	}
	base64file, err := imageProcessing_service.FileToBase64(path)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
		return
	}

	common.OkWithData(base64file, ctx)
}

// CreateImgWaterMarkWithIdio 创建水印
func CreateImgWaterMarkWithIdio(ctx *gin.Context) {
	widthtmp := ctx.DefaultPostForm("width", "0")    //原图宽度
	heighttmp := ctx.DefaultPostForm("height", "0")  //原图高度
	idio := ctx.DefaultPostForm("idio", "")          //签名图片（base64字符串）
	sizetmp := ctx.DefaultPostForm("size", "0")      //水印大小
	position := ctx.DefaultPostForm("position", "8") //水印在原图的定位
	xtmp := ctx.DefaultPostForm("x", "0")            //水印x轴
	ytmp := ctx.DefaultPostForm("y", "0")            //水印y轴
	width, _ := strconv.Atoi(widthtmp)
	height, _ := strconv.Atoi(heighttmp)
	size, _ := strconv.Atoi(sizetmp)
	x, _ := strconv.Atoi(xtmp)
	y, _ := strconv.Atoi(ytmp)

	if width == 0 || height == 0 {
		common.FailWithMsg("请传入正确的图片长宽！", ctx)
		return
	}
	if idio == "" {
		common.FailWithMsg("签名没有，还整啥水印啊！？", ctx)
		return
	}

	idioPath := "./app/imageProcessing/tmp/idio.png"
	waterMarkPath := "./app/imageProcessing/tmp/waterMark.png"

	//将base64转为图片并暂存
	err := imageProcessing_service.Base64ToFile(idio, idioPath)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
		return
	}

	//缩放图片
	length := 0
	if size == 0 {
		if width > height {
			length = height / 20
		} else {
			length = width / 20
		}
	} else {
		length = size
	}
	newImgIdio, err := imageProcessing_service.ResizeImg(idioPath, 0, length)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
		return
	}

	//创建一个原图大小的透明图片
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	//定位
	imageProcessing_service.SetLittleWaterMarkToBG(img, newImgIdio, 0.06, position, x, y)

	//保存
	path, err := imageProcessing_service.SaveFile(img, waterMarkPath)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
		return
	}
	base64file, err := imageProcessing_service.FileToBase64(path)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
		return
	}

	common.OkWithData(base64file, ctx)
}

// CreateImgWaterMarkFORM 以表单方式传入参数
func CreateImgWaterMarkFORM(ctx *gin.Context) {
	widthtmp := ctx.DefaultPostForm("width", "0")            //原图宽度
	heighttmp := ctx.DefaultPostForm("height", "0")          //原图高度
	color := ctx.DefaultPostForm("color", "white")           //颜色（黑白）
	rules := ctx.DefaultPostForm("rules", "H")               //排列规则（水平或垂直）
	sizetmp := ctx.DefaultPostForm("size", "0")              //水印大小
	position := ctx.DefaultPostForm("position", "8")         //水印在原图的定位
	xtmp := ctx.DefaultPostForm("x", "0")                    //水印x轴
	ytmp := ctx.DefaultPostForm("y", "0")                    //水印y轴
	waterMarkFilesArr := ctx.PostFormArray("waterMarkFiles") //水印文件
	width, _ := strconv.Atoi(widthtmp)
	height, _ := strconv.Atoi(heighttmp)
	size, _ := strconv.Atoi(sizetmp)
	x, _ := strconv.Atoi(xtmp)
	y, _ := strconv.Atoi(ytmp)

	var waterMarkStruct = typeWaterMarkStruct{
		Width:          width,
		Height:         height,
		Color:          color,
		Rules:          rules,
		Size:           size,
		Position:       position,
		X:              x,
		Y:              y,
		WaterMarkFiles: waterMarkFilesArr,
	}

	base64file, err := createImgWaterMark(waterMarkStruct)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	} else {
		common.OkWithData(base64file, ctx)
	}
}

// CreateImgWaterMarkJSON 以json方式传入参数
func CreateImgWaterMarkJSON(ctx *gin.Context) {
	var waterMarkStruct typeWaterMarkStruct
	if err := ctx.ShouldBind(&waterMarkStruct); err != nil {
		common.FailWithMsg(err.Error(), ctx)
		return
	}

	base64file, err := createImgWaterMark(waterMarkStruct)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	} else {
		common.OkWithData(base64file, ctx)
	}
}

// 生成水印
func createImgWaterMark(waterMarkStruct typeWaterMarkStruct) (string, error) {

	if waterMarkStruct.Width == 0 || waterMarkStruct.Height == 0 {
		return "", errors.New("请传入正确的图片长宽！")
	}

	fontPath := "./app/imageProcessing/fonts/方正硬笔行书繁体.ttf"

	if len(waterMarkStruct.WaterMarkFiles) <= 0 {
		return "", errors.New("啥都不传？那还搞什么水印")
	}

	//地址数组
	var filesPath []string
	for key, value := range waterMarkStruct.WaterMarkFiles {
		fileType := value[:1]
		fileContent := value[1:]
		if fileContent == "" || len(fileContent) == 0 {
			return "", errors.New("不得传入空内容")
		}
		filePath := "./app/imageProcessing/tmp/waterMarkFile" + strconv.Itoa(key) + ".png"
		if fileType == "W" {
			//将传入的文字转为图片
			_, err := imageProcessing_service.Words2Img(fileContent, waterMarkStruct.Color, fontPath, filePath)
			if err != nil {
				return "", err
			}
		} else if fileType == "T" {
			//将base64转为图片并暂存
			err := imageProcessing_service.Base64ToFile(fileContent, filePath)
			if err != nil {
				return "", err
			}
		} else {
			return "", errors.New("传入文件类型有误，请检查")
		}
		filesPath = append(filesPath, filePath)
	}

	//缩放图片
	length := 0
	if waterMarkStruct.Size == 0 {
		if waterMarkStruct.Width > waterMarkStruct.Height {
			length = waterMarkStruct.Height / 20
		} else {
			length = waterMarkStruct.Width / 20
		}
	} else {
		length = waterMarkStruct.Size
	}
	allWidth := 0
	allHeight := 0
	if waterMarkStruct.Rules == "V" {
		allHeight = length * len(filesPath)
		for _, filePath := range filesPath {
			newFile, err := imageProcessing_service.ResizeImg(filePath, 0, length)
			if err != nil {
				return "", err
			}
			if allWidth < newFile.Bounds().Dx() {
				allWidth = newFile.Bounds().Dx()
			}
		}
	} else if waterMarkStruct.Rules == "H" {
		allHeight = length
		for _, filePath := range filesPath {
			newFile, err := imageProcessing_service.ResizeImg(filePath, 0, length)
			if err != nil {
				return "", err
			}
			allWidth = allWidth + newFile.Bounds().Dx()
		}
	} else {
		return "", errors.New("请输入H或V")
	}
	//创建组合水印
	xxtmp := 0
	imgTmp := image.NewRGBA(image.Rect(0, 0, allWidth, allHeight))
	for seq, filePath := range filesPath {
		newFile, err := imageProcessing_service.ResizeImg(filePath, 0, length)
		if err != nil {
			return "", err
		}
		if waterMarkStruct.Rules == "V" {
			imageProcessing_service.SetLittleWaterMarkToBG(imgTmp, newFile, 0, "0", (allWidth-newFile.Bounds().Dx())/2, newFile.Bounds().Dy()*seq)
		}
		if waterMarkStruct.Rules == "H" {
			imageProcessing_service.SetLittleWaterMarkToBG(imgTmp, newFile, 0, "0", xxtmp, 0)
			xxtmp = xxtmp + newFile.Bounds().Dx()
		}
	}

	waterMarkPath := "./app/imageProcessing/tmp/waterMark.png"

	//创建一个原图大小的透明图片
	img := image.NewRGBA(image.Rect(0, 0, waterMarkStruct.Width, waterMarkStruct.Height))

	//定位
	imageProcessing_service.SetLittleWaterMarkToBG(img, imgTmp, 0.06, waterMarkStruct.Position, waterMarkStruct.X, waterMarkStruct.Y)

	//保存
	path, err := imageProcessing_service.SaveFile(img, waterMarkPath)
	if err != nil {
		return "", err
	}
	base64file, err := imageProcessing_service.FileToBase64(path)
	if err != nil {
		return "", err
	}

	return base64file, nil
}
