package time

import (
	"time"
)

//获取时间戳
func Time() int64 {
	return time.Now().Unix()
}

//时间戳转化成时间
func TimeToFormat(t int64, format string) string {
	//var format string="2006-01-02 15:04:05" 时间格式参考
	return time.Unix(t, 0).Format(format)
}

//将时间转化成时间戳
func FormatToTime(datetime, timeLayout string) int64 {
	//datetime := "2015-01-01 00:00:00"  //待转化为时间戳的字符串
	//timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64
	return timestamp
}

//当天0点时间
func TimeDayZero() int64 {
	return FormatToTime(TimeToFormat(Time(), "2006-01-02"), "2006-01-02")
}

//当天24点时间
func TimeThatDay24() int64 {
	return TimeDayZero() + 86400
}
//获取这一段时间的时间戳
func OneThouthTime(starTime, endTime int64) []int64 {
	var moth []int64
	for true {
		if starTime == endTime {
			goto Loop
		}
		moth = append(moth, starTime)
		starTime -= 86400
	}
	Loop:
	return moth
}
func GormTimeToTime(t  time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
func GormTimeToInt(t time.Time,timeLayout string) int64 {
	return  FormatToTime(t.Format(timeLayout),timeLayout)
}
