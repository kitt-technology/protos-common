package common

import "time"

func (date *Date) ToTime() time.Time {
	return time.Date(int(date.Year), time.Month(date.Month), int(date.Day), 0, 0, 0, 0, time.UTC)
}

func (dates *DateRange) IsBetween(t time.Time) bool {
	if dates.Min != nil && dates.Max != nil {
		return dates.Min.ToTime().Before(t) && dates.Max.ToTime().After(t)
	} else if dates.Min != nil {
		return dates.Min.ToTime().Before(t)
	} else if dates.Max != nil {
		return dates.Max.ToTime().After(t)
	}
	return true
}

func ToCommonDate(t time.Time) *Date {
	return &Date{
		Year:  int32(t.Year()),
		Month: int32(t.Month()),
		Day:   int32(t.Day()),
	}
}

func ToNullableCommonDate(t *time.Time) *Date {
	if t == nil {
		return nil
	}
	return &Date{
		Year:  int32(t.Year()),
		Month: int32(t.Month()),
		Day:   int32(t.Day()),
	}
}
