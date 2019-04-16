package main

import (
	"errors"
	"fmt"
	"time"
)

const (
	DateTime     = "2006-01-02 15:04:05"
	DateTime2    = "2006-01-02 1504"
	Date         = "20060102"
	DateClock    = "2006010215"
	DayDuration  = 24 * time.Hour
	WeekDuration = 7 * 24 * time.Hour
	Second2500   = 16725196800 // 2100-01-01 时间戳， 如果超过这个时间认为不再是秒
	DaySecond    = 86400
)

func GetNow() int64 {
	return time.Now().Unix()
}

func GetNowInMilliSecond() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func GetDayInMilliSecond(t time.Time) int64 {
	return t.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func ToDateClock(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(DateClock)
}

func ToDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(Date)
}

func ToDateTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(DateTime)
}

func ToDateTimeInFormat(timestamp int64, format string) string {
	return time.Unix(timestamp, 0).Format(format)
}

func ToTimeInFormat(datetime string, format string) (time.Time, error) {
	return time.Parse(format, datetime)
}

func ToLocalTimeInFormat(datetime string, format string) (time.Time, error) {
	TZ, _ := time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation(format, datetime, TZ)
}

func GetTodayStartTimestamp() int64 {
	return time.Now().Unix() - (time.Now().Unix() % 3600) - int64(time.Now().Hour())*3600
}

func GetSomedayStartTimestamp(timestamp int64) int64 {
	return timestamp - (timestamp % 3600) - int64(time.Unix(timestamp, 0).Hour())*3600

}

// 一个时间戳距离当前时间的天数
func GetDayToNow(timestampSecond int64) (int64, error) {
	if timestampSecond > Second2500 {
		return 0, errors.New("timestamp must be second")
	}
	if timestampSecond <= 0 {
		return 0, errors.New("timestamp must > 0")
	}
	return (time.Now().Unix() - timestampSecond) / 86400, nil
}

// 距今时间的一个 Unix 时间戳
func GetUnixTimestampToNow(years, months, days int) int64 {
	return time.Now().AddDate(years, months, days).Unix()
}

func GetNowRangeInFormat(years int, months int, days int, format string) string {
	return time.Now().Local().AddDate(years, months, days).Format(format)
}

// 当天时间片编号，每 k 分钟一个时间片
func GetTimeIndexByMinutes(k int) int {
	return (time.Now().Minute() + time.Now().Hour()*60) / k
}

func main() {
	fmt.Println("time now: ", GetNow())
	fmt.Println("time now: ", GetNowInMilliSecond())
	t, err := ToTimeInFormat("2019-04-12 1200", DateTime2)
	fmt.Println("time now: ", GetDayInMilliSecond(t), err)
}
