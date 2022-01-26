package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDateToTime(t *testing.T) {
	date := Date{Year: 1999, Month: 6, Day: 21}

	_time := date.ToTime()
	assert.Equal(t, _time.Year(), 1999)
	assert.Equal(t, _time.Month(), time.Month(6))
	assert.Equal(t, _time.Day(), 21)
}
