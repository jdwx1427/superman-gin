package utils

import (
	"database/sql"
	"time"
)

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func TimeToUnix(t time.Time) int64 {
	return t.Unix()
}

func TimeToUnixString(t time.Time) string {
	unix := TimeToUnix(t)
	if unix < 0 {
		unix = 0
	}
	return Int64ToString(unix)
}

func NullTimeToUnixString(t sql.NullTime) string {
	if !t.Valid {
		return ""
	}

	return TimeToUnixString(t.Time)
}

func TimeToStringCompact(t time.Time) string {
	return t.Format("20060102150405")
}

func NullTimeToString(t sql.NullTime) string {
	if !t.Valid {
		return ""
	}

	return t.Time.Format("2006-01-02 15:04:05")
}

func NullTimeToUnix(t sql.NullTime) int64 {
	if !t.Valid {
		return 0
	}

	return t.Time.Unix()
}

func UnixToNullTime(t int64) sql.NullTime {
	if t <= 0 {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: time.Unix(t, 0), Valid: true}
}

func TimeToDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func NullTimeToDate(t sql.NullTime) string {
	if !t.Valid {
		return ""
	}

	return t.Time.Format("2006-01-02")
}

func TodayDate() string {
	t := time.Now()
	return t.Format("2006-01-02")
}

func NowTime() time.Time {
	t := time.Now()
	return t
}

func TodayStartTimeString() string {
	return TodayDate() + " 00:00:00"
}

func TodayStartTime() time.Time {
	todayStartTime, _ := DateTimeStringToTime(TodayStartTimeString())

	return todayStartTime
}

func TodayEndTimeString() string {
	return TodayDate() + " 23:59:59"
}

func TodayEndTime() time.Time {
	todayEndTime, _ := DateTimeStringToTime(TodayEndTimeString())

	return todayEndTime
}

func TodayEndTimeUnix() int64 {
	todayEndTime := TodayEndTime()

	return todayEndTime.Unix()
}

func OneHourLaterTime() time.Time {
	now := time.Now()
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	return later
}

func NowHour() int64 {
	t := int64(time.Now().Local().Hour())
	return t
}

func NowTimeString() string {
	return TimeToString(NowTime())
}

func NowTimeStringCompact() string {
	return TimeToStringCompact(NowTime())
}

func YesterdayDate() string {
	return TimeToDate(DateToTime(TodayDate()).AddDate(0, 0, -1))
}

// 前一天的日期
func PreviousDate(date string) string {
	return TimeToDate(DateToTime(date).AddDate(0, 0, -1))
}

// 获取x天前的日期，按照今天的日期开始算
func GetFewDaysAgo(t string, agoDay int) string {
	return TimeToDate(DateToTime(t).AddDate(0, 0, agoDay))
}

func Previous7daysDate() string {
	return TimeToDate(DateToTime(TodayDate()).AddDate(0, 0, -6))
}

func PreviousMonthDate() string {
	return TimeToDate(DateToTime(TodayDate()).AddDate(0, 0, -30))
}

func DateToTime(date string) time.Time {
	st, _ := time.Parse("2006-01-02", date)
	return st
}

func IsDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false
	}
	return true
}

func MonthMemo(month string) string {
	timeFormat := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t2, _ := time.ParseInLocation(timeFormat, month+"-01 00:00:00", loc)
	monthMemo := t2.Format("2006年01月")

	return monthMemo
}

func DateStringToTime(date string) sql.NullTime {
	st := sql.NullTime{Valid: false}

	if date != "" {
		tt := DateToTime(date)
		st = sql.NullTime{Time: tt, Valid: true}
	}

	return st
}

func TodayDatePath() string {
	t := time.Now()
	return t.Format("2006/01/02")
}

func NowTimeUnix() int64 {
	t := time.Now().Unix()
	return t
}

func NowTimeUnixString() string {
	t := Int64ToString(NowTimeUnix())
	return t
}

func DateTimeStringToTime(datetime string) (time.Time, error) {
	st, err := time.ParseInLocation("2006-01-02 15:04:05", datetime, time.Local)

	return st, err
}

func DateTimeStringToUnix(datetime string) int64 {
	st, err := time.ParseInLocation("2006-01-02 15:04:05", datetime, time.Local)

	if err != nil {
		return 0
	}

	return st.Unix()
}

func DateStringToUnix(date string) int64 {
	st, err := time.ParseInLocation("2006-01-02", date, time.Local)

	if err != nil {
		return 0
	}

	return st.Unix()
}

func DateTimeZoneToUnix(datetime string) int64 {
	st, err := time.Parse(time.RFC3339, datetime)

	if err != nil {
		return 0
	}

	return st.Unix()
}

func DateTimeStringToString(datetime string) string {
	st, err := time.Parse(time.RFC3339, datetime)

	if err != nil {
		return datetime
	}

	str := TimeToString(st)

	return str
}

func DateStringToString(date string) string {
	st, err := time.Parse(time.RFC3339, date)

	if err != nil {
		return date
	}

	str := TimeToDate(st)

	return str
}

func ExpireEndTime(endTimeBase int64, addDays int64) time.Time {
	if addDays%365 == 0 {
		return time.Unix(endTimeBase, 0).AddDate(int(addDays/365), 0, 0)
	} else if addDays%30 == 0 {
		return time.Unix(endTimeBase, 0).AddDate(0, int(addDays/30), 0)
	} else {
		return time.Unix(endTimeBase, 0).AddDate(0, 0, int(addDays))
	}
}

func DateDiffDays(dateMin, dateMax string) int64 {
	t1 := DateToTime(dateMin)
	t2 := DateToTime(dateMax)
	return int64(t2.Sub(t1).Hours()/24) + 1
}

func FormatSeconds(seconds int64) string {
	if seconds == 0 {
		return ""
	}

	secondTime := seconds  // 秒
	minuteTime := int64(0) // 分
	hourTime := int64(0)   // 小时
	if secondTime > 60 {   // 如果秒数大于60，将秒数转换成整数
		// 获取分钟，除以60取整数，得到整数分钟
		minuteTime = int64(secondTime / 60)
		// 获取秒数，秒数取佘，得到整数秒数
		secondTime = int64(secondTime % 60)
		// 如果分钟大于60，将分钟转换成小时
		if minuteTime > 60 {
			// 获取小时，获取分钟除以60，得到整数小时
			hourTime = int64(minuteTime / 60)
			// 获取小时后取佘的分，获取分钟除以60取佘的分
			minuteTime = int64(minuteTime % 60)
		}
	}
	result := "" + Int64ToString(secondTime) + ""
	result = "" + Int64ToString(minuteTime) + ":" + result
	result = "" + Int64ToString(hourTime) + ":" + result

	return result
}

func FormatSecondsToTimeLength(seconds int64) string {
	if seconds == 0 {
		return ""
	}

	secondTime := seconds  // 秒
	minuteTime := int64(0) // 分
	hourTime := int64(0)   // 小时
	if secondTime > 60 {   // 如果秒数大于60，将秒数转换成整数
		// 获取分钟，除以60取整数，得到整数分钟
		minuteTime = int64(secondTime / 60)
		// 获取秒数，秒数取佘，得到整数秒数
		secondTime = int64(secondTime % 60)
		// 如果分钟大于60，将分钟转换成小时
		if minuteTime > 60 {
			// 获取小时，获取分钟除以60，得到整数小时
			hourTime = int64(minuteTime / 60)
			// 获取小时后取佘的分，获取分钟除以60取佘的分
			minuteTime = int64(minuteTime % 60)
		}
	}

	result := ""

	if hourTime > 0 {
		result += Int64ToString(hourTime) + "小时"
	}
	if minuteTime > 0 {
		result += Int64ToString(minuteTime) + "分钟"
	}
	if secondTime > 0 {
		result += Int64ToString(secondTime) + "秒"
	}

	return result
}

func CompareDateLess(preDates []string, date string, addParams map[string]string) bool {
	for _, v := range preDates {
		if addParams[v] <= date {
			return true
		}
	}
	return false
}

func DateNowFormatStr() string {
	tm := time.Now()
	return tm.Format("2006-01-02 15:04:05")
}
