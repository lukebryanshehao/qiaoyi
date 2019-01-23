package utils

import (
	"strconv"
	"time"
)

//获取系统当前时间
func GetTimeNow() *time.Time {
	//获取系统当前时间
	currentTimeData := time.Date(time.Now().Year(), time.Now().Month(),
		time.Now().Day(), time.Now().Hour(), time.Now().Minute(),
		time.Now().Second(), time.Now().Nanosecond(), time.Local)
	return &currentTimeData
}

//处理时间格式
func SubTime(str string) string {
	return str[0 : len(str)-15]
}

//获取当前时间戳
func GetTimeUnix() string {
	timeUnix := time.Now().UnixNano()
	i := int64(timeUnix)
	fileName := strconv.FormatInt(i, 10)
	return fileName
}

//time格式转string
func TimeFormatString(time time.Time) string {
	baseFormat := "2006-01-02 15:04:05"
	strTime := time.Format(baseFormat)
	return strTime
}

//string格式转time
func StringFormatTime(times string) time.Time {
	baseFormat := "2006-01-02 15:04:05"
	parseStrTime, _ := time.Parse(baseFormat, times)
	return parseStrTime
}
