package utils

import (
	"fmt"
	"time"
)

func util2() {
	//包含时区，毫秒等信息
	nowtime := time.Now()
	fmt.Println(nowtime)

	fmt.Println(nowtime.Year())
	//英文月份
	fmt.Println(nowtime.Month())
	//数字月份
	fmt.Println(int(nowtime.Month()))
	fmt.Println(nowtime.Day())
	//24 小时制
	fmt.Println(nowtime.Hour())
	fmt.Println(nowtime.Minute())
	fmt.Println(nowtime.Second())

	//格式化输出
	stime := "年月日 %02d-%02d-%02d %02d:%02d:%02d"
	fmt.Printf(stime, nowtime.Year(), nowtime.Month(), nowtime.Day(), nowtime.Hour(), nowtime.Minute(), nowtime.Second())
	stime = fmt.Sprintf(stime, nowtime.Year(), nowtime.Month(), nowtime.Day(), nowtime.Hour(), nowtime.Minute(), nowtime.Second())
	fmt.Println(stime)

	//format := "2006-01-02 15:04:05"
	year := nowtime.Format("2006")
	fmt.Println(year)
	day := nowtime.Format("02")
	fmt.Println(day)
	second := nowtime.Format("05")
	fmt.Println(second)

	//休眠时间 1秒，100毫秒
	i := time.Second
	fmt.Println(i)
	i = time.Millisecond * 100
	fmt.Println(i)

	//时间戳，秒，纳秒
	secondnum := nowtime.Unix()
	secondnum = nowtime.UnixNano()
	fmt.Println(secondnum)
}

func GetNow() string {
	//这个日期是个固定值，是用于格式化日期的，是 go 语言的规范
	format := "2006-01-02 15:04:05"
	now := time.Now()
	return now.Format(format)
}

func GetThisYear() string {
	now := time.Now()
	return now.Format("2006")
}

func GetUnixTime() int64 {
	now := time.Now()
	return now.Unix()
}
