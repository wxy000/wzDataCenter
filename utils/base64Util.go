package utils

import (
	"encoding/base64"
	"io/ioutil"
)

// FileToBase64 文件转base64
func FileToBase64(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	base64Str := base64.StdEncoding.EncodeToString(data)
	return base64Str, nil
}

// Base64ToFile base64转文件
func Base64ToFile(base64Str string, outputPath string) error {
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
