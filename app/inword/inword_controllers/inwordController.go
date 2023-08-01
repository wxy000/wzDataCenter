package inword_controllers

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"math/rand"
	"os"
	"strings"
	"time"
	"wzDataCenter/common"
)

// 获取随机行数（rows表示传入总行数）
func getRandom(rows int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(rows)
}

// 获取总行数
func getRows(filepath string) int {
	fileHanle, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		return 0
	}
	defer fileHanle.Close()

	scanner := bufio.NewScanner(fileHanle)
	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0
	}
	return count
}

// GetRandomWord 随机获取一行文本
func GetRandomWord(ctx *gin.Context) {
	filename := ctx.DefaultQuery("file", "default")

	filepath := "./app/inword/words/" + filename + ".txt"

	fileHanle, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}
	defer fileHanle.Close()

	scanner := bufio.NewScanner(fileHanle)

	// 获取总行数
	count := getRows(filepath)
	// 获取随机行数
	randomRow := getRandom(count) + 1

	var result string
	// 按行处理txt
	i := 0
	for scanner.Scan() {
		i++
		if i == randomRow {
			result = strings.TrimSpace(scanner.Text())
			break
		}
	}
	if result == "" {
		result = "人生就是这样，你能见到的永远只有自己"
	}
	common.OkWithData(result, ctx)
}

// GetRandomImg 随机获取一行文本
func GetRandomImg(ctx *gin.Context) {
	filename := ctx.DefaultQuery("file", "default")

	filepath := "./app/inword/images/" + filename + ".txt"

	fileHanle, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	}
	defer fileHanle.Close()

	scanner := bufio.NewScanner(fileHanle)

	// 获取总行数
	count := getRows(filepath)
	// 获取随机行数
	randomRow := getRandom(count) + 1

	var result string
	// 按行处理txt
	i := 0
	for scanner.Scan() {
		i++
		if i == randomRow {
			result = strings.TrimSpace(scanner.Text())
			break
		}
	}
	if result == "" {
		result = "xxx"
	}
	common.OkWithData(result, ctx)
}
