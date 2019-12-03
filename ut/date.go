/**
 * @Author: Machao
 * @Date: 2019-12-02 10:17
 * @To:
 */
package ut

import (
	"time"
)

type FormatT int

const (
	Day  FormatT = 0
	Hour         = 1 << iota
	Minute
	Second
)

func NowTimeStr(format FormatT) string {
	str := "2006-01-02"

	switch format {
	case Hour:
		str += " 15"
	case Minute:
		str += " 15:04"
	case Second:
		str += " 15:04:05"
	default:
	}

	return time.Now().Format(str)
}

func Strtotime(timeStr string, format FormatT) (time.Time, error) {
	str := "2006-01-02"

	switch format {
	case Hour:
		str += " 15"
	case Minute:
		str += " 15:04"
	case Second:
		str += " 15:04:05"
	default:
	}

	return time.ParseInLocation(str, timeStr, time.Local)
}
