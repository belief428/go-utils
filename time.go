package utils

import (
	"math"
	"time"
)

type Week int

const (
	Monday Week = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func IsEmptyTime(t time.Time) bool {
	return t.IsZero()
}

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func FormatDatetime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func FormatTimeForLayout(t time.Time, layout string) string {
	return t.Format(layout)
}

func DateTimeToTime(t string) time.Time {
	_time, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	return _time
}

func DataTimeToDate(t string) time.Time {
	_time, _ := time.ParseInLocation("2006-01-02", t, time.Local)
	return _time
}

func DataTimeForLayout(t string, layout string) time.Time {
	_time, _ := time.ParseInLocation(layout, t, time.Local)
	return _time
}

func GetMondayTime(t time.Time) time.Time {
	offset := int(time.Monday - t.Weekday())

	if offset > 0 {
		offset = -6
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
}

func MonthBeginAt(year, month int) time.Time {
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Location())
}

func MonthFinishAt(year, month int) time.Time {
	mark := time.Date(year, time.Month(month), 1, 23, 59, 59, 0, time.Now().Location())
	return mark.AddDate(0, 1, -1)
}

func DiffTimeMonth(time1, time2 time.Time) int {
	year := math.Abs(float64(time1.Year()) - float64(time2.Year()))
	month := math.Abs(float64(time1.Month()) - float64(time2.Month()))
	return int(year)*12 + int(month) + 1
}
