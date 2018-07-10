package utils

import (
	"time"
	"strconv"
	"strings"
)

func StrToIntMonth(month string) int {
	date := map[string]int {
		"January": 0,
		"February": 1,
		"March": 2,
		"April": 3,
		"May": 4,
		"June": 5,
		"July": 6,
		"August": 7,
		"September": 8,
		"October": 9,
		"November": 10,
		"December": 11,
	}
	return date[month]
}

func GetTodayYMD(sep string) string {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()

	monthStr := strconv.Itoa(int(month))
	if month < 10 {
		monthStr = "0" + monthStr
	}

	dayStr := strconv.Itoa(day)
	if day < 10 {
		dayStr = "0" + dayStr
	}

	yearStr := strconv.Itoa(year)

	return yearStr + sep + monthStr + sep + dayStr
}

func GetTodayYM(sep string) string {
	now := time.Now()
	year := now.Year()
	month := now.Month()

	yearStr := strconv.Itoa(year)

	monthStr := strconv.Itoa(int(month))
	if month < 10 {
		monthStr = "0" + monthStr
	}
	return yearStr + sep + monthStr
}

func GetYesterdayYMD(sep string) string {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	todaySec := today.Unix()
	yesterdaySec := todaySec - 60*60*24
	yesterday := time.Unix(yesterdaySec, 0)
	yesterdayYMD := yesterday.Format("2006-01-02")
	return strings.Replace(yesterdayYMD, "-", sep, -1)
}

func GetTomorrowYMD(sep string) string {
	now := time.Now()
	todaySec := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Unix()
	tomorrowSec := todaySec + 60*60*24
	tomorrow := time.Unix(tomorrowSec, 0)
	tomorrowYMD := tomorrow.Format("2006-01-02")
	return strings.Replace(tomorrowYMD, "-", sep, -1)
}

func GetTodayMidnightTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
}

func GetYesterdayMidnightTime() time.Time {
	now := time.Now()
	todaySec := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Unix()
	yesterdaySec := todaySec - 60*60*24
	return time.Unix(yesterdaySec, 0)
}