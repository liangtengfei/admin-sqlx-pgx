package datetime

import "time"

const (
	FormatDate     = "2006-01-02"
	FormatDatetime = "2006-01-02 15:04:05"
	FormatTime     = "15:04:05"
)

func ToDatetime(time time.Time) string {
	return time.Format(FormatDatetime)
}

func WeekBeginTime(p ...time.Time) string {
	t := time.Now()
	if len(p) == 1 {
		t = p[0]
	}
	return With(t).BeginningOfWeek().Format(FormatDatetime)
}

func WeekEndTime(p ...time.Time) string {
	t := time.Now()
	if len(p) == 1 {
		t = p[0]
	}
	return With(t).EndOfWeek().Format(FormatDatetime)
}
