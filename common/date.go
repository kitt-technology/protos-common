package common

import "time"

func (date *Date) ToTime() time.Time {
	return time.Date(int(date.Year), time.Month(date.Month), int(date.Day), 0, 0, 0, 0, time.UTC)
}