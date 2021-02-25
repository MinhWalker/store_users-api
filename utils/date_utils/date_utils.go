package date_utils

import "time"

const(
	apiDatelayout = "02-01-2006T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDatelayout)
}
