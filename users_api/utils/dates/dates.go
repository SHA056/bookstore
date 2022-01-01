package dates

import "time"

const (
	apiDatelayout = "2006 Jan 2 15:04:05 MST"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	now := GetNow()
	return now.Format(apiDatelayout)
}
