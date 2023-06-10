package util

import (
	"time"
)

// 秒级时间戳=>日期
func TimeToDate(t time.Time) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	date := t.In(loc).Format("2006-01-02")
	return date
}

func TimeToDateTime(t time.Time) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	date := t.In(loc).Format("2006-01-02 15:04:05")
	return date
}

// 周X
func TimeToExportWeekDay(t time.Time) string {
	weekDay := t.Weekday()
	switch weekDay {
	case 0:
		return "周日"
	case 1:
		return "周一"
	case 2:
		return "周二"
	case 3:
		return "周三"
	case 4:
		return "周四"
	case 5:
		return "周五"
	case 6:
		return "周六"
	}
	return Val2string(weekDay)
}

// 800=>08:00
func FormatTimeToString(t int32) string {
	tInt := t / 100
	s1 := t - (t/10)*10
	s2 := t/10 - (t/100)*10
	var tStr string
	if tInt < 10 {
		tStr = "0" + Val2string(tInt)
	} else {
		tStr = Val2string(tInt)
	}
	return tStr + ":" + Val2string(s2) + Val2string(s1)

}

func DateTimeStrToTime(v string) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	resTime, _ := time.ParseInLocation("2006-01-02 15:04:05", v, loc)
	return resTime
}

func DateStrToTime(v string) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	resTime, _ := time.ParseInLocation("2006-01-02", v, loc)
	return resTime
}

func StrToTimeInt(v string) int64 {
	t := DateTimeStrToTime(v)
	return t.Unix()
}

func TimeIntToDateTimeStr(t int64) string {
	return TimeToDateTime(time.Unix(t, 0))
}

func FormatDurationToStr(d int64) string {
	var resStr string
	day := d / 24 / 3600
	if day > 0 {
		resStr += Val2string(day) + "天"
	}
	d = d - day*24*3600
	hour := d / 3600
	if hour > 0 {
		resStr += Val2string(hour) + "小时"
	}
	d = d - hour*3600
	minute := d / 60
	if minute > 0 {
		resStr += Val2string(minute) + "分钟"
	}
	return resStr
}
