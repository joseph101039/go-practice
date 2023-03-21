package jwt

import "math/rand"

// RandString 產生隨機英文字母
func RandString(strlen int) string {
	return randLetters([]rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"), strlen)
}

// RandStringNumber 產生隨機英文字母與數字
func RandStringNumber(strlen int) string {
	return randLetters([]rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"), strlen)
}

func randLetters(letterPool []rune, strlen int) string {
	letterLen := len(letterPool)
	str := make([]rune, strlen)
	for i := range str {
		str[i] = letterPool[rand.Intn(letterLen)]
	}
	return string(str)
}
