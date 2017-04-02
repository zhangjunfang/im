package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const TIME_FORMAT_YYYYMMDDHHMMSS string = "2006-01-02 15:04:05"
const TIME_FORMAT_YYYYMMDD string = "2006-01-02"
const TIME_FORMAT_YYYYMM string = "200601"
const TIME_FORMAT_MMDD string = "0102"

//const TIMEFORMAT string = "2006-01-02"
const TIMEFORMAT_YEAR string = "2006"

var count int64 = 0

//当前时间格式
func NowTime() string {
	return time.Now().Format(TIME_FORMAT_YYYYMMDDHHMMSS)
}

//当前时间格式
func NowDate() string {
	return time.Now().Format(TIME_FORMAT_YYYYMMDD)
}

//字符串时间 格式化为当地时区时间
func Str2Time(str string) (time.Time, error) {
	return time.ParseInLocation(TIME_FORMAT_YYYYMMDDHHMMSS, str, time.Local)
}

//字符串时间 格式化为当地时区时间
func Str2Date(str string) (time.Time, error) {
	return time.ParseInLocation(TIME_FORMAT_YYYYMMDD, str, time.Local)
}

//随机生成随机数
func GetRand(r int) int {
	atomic.AddInt64(&count, 1)
	if atomic.LoadInt64(&count) > 100000000 {
		count = 0
	}
	rd := rand.New(rand.NewSource(time.Now().UnixNano() + count))
	return rd.Intn(r)
}

//
func Chatid(fid, tid, domain string) string {
	if fid < tid {
		fid, tid = tid, fid
	}
	m := md5.New()
	m.Write([]byte(fmt.Sprint(fid, "tim", domain, tid)))
	return hex.EncodeToString(m.Sum(nil))
}

func MD5(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return strings.ToUpper(hex.EncodeToString(m.Sum(nil)))
}

func Sha1(s string) string {
	m := sha1.New()
	m.Write([]byte(s))
	return strings.ToUpper(hex.EncodeToString(m.Sum(nil)))
}

// 毫秒
func TimeMills() string {
	return fmt.Sprint(time.Now().UnixNano() / 1000000)
}

//
func TimeMillsInt64() int64 {
	return time.Now().UnixNano() / 1000000
}

//  毫秒转换为格式化时间
func TimeMills2TimeFormat(mill string) string {
	millint, _ := strconv.Atoi(mill)
	return time.Unix(int64(millint/1000), 0).Format(TIME_FORMAT_YYYYMMDDHHMMSS)
}

//—-------------------------------------------------------------------------------------------
// 今天剩余时间 单位 秒
func RestSecond() int {
	return ((24-time.Now().Hour())*60 - time.Now().Minute()) * 60
}
