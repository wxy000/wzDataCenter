package utils

import (
	"bufio"
	"os"
	"strings"
)

// ReadBlacklist 读取黑名单并判断ip是否在黑名单中
func ReadBlacklist(filepath string, ip string) (bool, error) {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ip == strings.TrimSpace(scanner.Text()) {
			return true, nil
		}
	}
	return false, nil
}
