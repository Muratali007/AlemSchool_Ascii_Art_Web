package ascii

import (
	logger2 "ascii_art_web/logger"
	"crypto/md5"
	"errors"
	"fmt"
	"os"
)

var logger = logger2.GetLogger()

func Ascii(text, banner string) (string, error) {
	if len(text) <= 0 {
		logger.Info("error")
		return "", nil
	}

	fileName := ""

	switch banner {
	case "Standard":
		fileName = "standard.txt"
	case "Shadow":
		fileName = "shadow.txt"
	case "Thinkertoy":
		fileName = "thinkertoy.txt"
	default:
		logger.Info("error: font does not exist")
		return "", errors.New("error: font does not exist")
	}

	for _, val := range text {
		if !(val >= 0 && val <= '~') {
			logger.Info("error: not ASCII character")
			return "", errors.New("error: not ASCII character")
		}
	}

	content, err := os.ReadFile("./ascii/banners/" + fileName)
	if err != nil {
		logger.Info("error")
		return "", err
	}

	hash := MD5(string(content))
	if checkHash(hash, fileName) {
		logger.Info("error: banner should not be changed")
		return "", errors.New("error: banner should not be changed")
	}

	standFont, err := ReadFont(string(content))
	if err != nil {
		return "", fmt.Errorf("Error %v", err)
	}

	return ConvertedText(text, standFont), nil

}

func MD5(s string) string {
	h := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", h)
}

func checkHash(filename, check string) bool {
	hashStandard := "ac85e83127e49ec42487f272d9b9db8b"
	hashShadow := "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	hashThinkertoy := "86d9947457f6a41a18cb98427e314ff8"

	if check == hashStandard && filename == "standard.txt" {
		return true
	}
	if check == hashShadow && filename == "shadow.txt" {
		return true
	}
	if check == hashThinkertoy && filename == "thinkertoy.txt" {
		return true
	}

	return false
}
