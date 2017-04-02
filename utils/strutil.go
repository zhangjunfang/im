package utils

import (
	"regexp"
	"strconv"
	"strings"
)

//字符串截取  性能不是最佳的方式
func Substr(str string, start, length int) string {
	if start < 0 || length < 0 {
		return str
	}
	rs := []rune(str)
	end := start + length
	return string(rs[start:end])
}

//判断字符串是否为空
func IsStringEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

//字符串数字转化为int64
func Atoi64(str string) int64 {
	i, _ := strconv.Atoi(str)
	return int64(i)
}

//字符串数字转化为int32
func Atoi32(str string) int32 {
	i, _ := strconv.Atoi(str)
	return int32(i)
}

//字符串数字转化为int
func Atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

//字符串数字转化为int16
func Atoi16(str string) int16 {
	i, _ := strconv.Atoi(str)
	return int16(i)
}

//是否是字符串数字
func IsNum(ss ...string) bool {
	if ss == nil || len(ss) == 0 {
		return false
	}
	pattern := "^\\d{0,}$"
	for _, s := range ss {
		b, err := regexp.MatchString(pattern, s)
		if err != nil || !b {
			return false
		}
	}
	return true
}

//字符串连接  默认使用英文逗号连接
func Join(ss ...string) string {
	if ss == nil || len(ss) == 0 {
		return ""
	}
	return strings.Join(ss, ",")
}
