package utils

import (
	"strconv"
	"time"
)

func FormatTimeDuration(d *time.Duration) string {
	if d == nil {
		return "0s"
	}
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	resultStr := ""
	if days := hours / 24; days > 0 {
		resultStr += strconv.Itoa(days) + "d"
	}
	if hoursNum := hours % 24; hoursNum > 0 {
		resultStr += strconv.Itoa(hoursNum) + "h"
	}
	if minutes > 0 {
		resultStr += strconv.Itoa(minutes) + "m"
	}
	if seconds > 0 {
		resultStr += strconv.Itoa(seconds) + "s"
	}
	return resultStr
}

func FormatTime(t time.Time) string {
	// 将时间格式化为指定格式的字符串
	return t.Format("2006-01-02 15:04:05")
}
