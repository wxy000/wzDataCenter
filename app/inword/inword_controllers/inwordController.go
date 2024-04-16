package inword_controllers

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"wzDataCenter/common"
)

// 获取随机数（rows表示传入总行数，num表示产生的随机数的个数）
func getRandoms(rows int, num int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, num)
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(rows) + 1
	loop:
		for k := 0; k < i; k++ {
			if nums[k] == nums[i] {
				nums[i] = rand.Intn(rows) + 1
				goto loop
			}
		}
	}
	sort.Ints(nums)
	return nums
}

// 二分法查找元素是否在数组中
func binarySearch(des []int, goal int) bool {
	locate := len(des) / 2
	//当递归到只剩一个元素，判断此元素与目标元素是否相同，如果相同，则在数组内，不相同
	if len(des) == 1 {
		if des[0] == goal {
			return true
		} else {
			return false
		}
	}
	//如果数数组中间的元素大于目标元素，则进行左半部分递归
	if des[locate] > goal {
		return binarySearch(des[:locate], goal)
	} else {
		//同理，相反情况，右半部分进行递归
		return binarySearch(des[locate:], goal)
	}
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

// GetRandomWords 随机获取n行文本
func GetRandomWords(ctx *gin.Context) {
	nums := ctx.DefaultQuery("nums", "1")
	numss, err := strconv.Atoi(nums)
	if err != nil {
		numss = 1
	}
	filename := ctx.DefaultQuery("file", "default")

	filepath := "./app/inword/words/" + filename + ".txt"

	fileHanle, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
		return
	}
	defer fileHanle.Close()

	scanner := bufio.NewScanner(fileHanle)

	// 获取总行数
	count := getRows(filepath)
	if numss > count {
		numss = count
	}
	// 获取随机行数
	randomRows := getRandoms(count, numss)

	result := make([]string, numss)
	// 按行处理txt
	i := 0
	j := 0
	for scanner.Scan() {
		i++
		if binarySearch(randomRows, i) {
			result[j] = strings.TrimSpace(scanner.Text())
			j++
		}
	}
	if result[0] == "" {
		result[0] = "人生就是这样，你能见到的永远只有自己"
	}
	common.OkWithData(result, ctx)
}

// GetRandomImgs 随机获取一行文本
func GetRandomImgs(ctx *gin.Context) {
	nums := ctx.DefaultQuery("nums", "1")
	numss, err := strconv.Atoi(nums)
	if err != nil {
		numss = 1
	}
	filename := ctx.DefaultQuery("file", "default")
	isPage := ctx.DefaultQuery("isPage", "N")
	if isPage != "N" && isPage != "Y" {
		isPage = "N"
	}

	filepath := "./app/inword/images/" + filename + ".txt"

	fileHanle, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
		return
	}
	defer fileHanle.Close()

	scanner := bufio.NewScanner(fileHanle)

	// 获取总行数
	count := getRows(filepath)
	if numss > count {
		numss = count
	}
	// 获取随机行数
	randomRows := getRandoms(count, numss)

	result := make([]string, numss)
	// 按行处理txt
	i := 0
	j := 0
	for scanner.Scan() {
		i++
		if binarySearch(randomRows, i) {
			result[j] = strings.TrimSpace(scanner.Text())
			j++
		}
	}
	if result[0] == "" {
		result[0] = "xxx"
	}
	if isPage == "N" {
		common.OkWithData(result, ctx)
	} else {
		ctx.HTML(http.StatusOK, "index.html",
			common.Response{
				Code: http.StatusOK,
				Msg:  "获取图片成功",
				Data: result,
			})
	}
}
