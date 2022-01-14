package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoneyFromArgs(t *testing.T) {
	valid := map[string]interface{}{
		"currencyCode": "GBP",
		"units":        1234,
	}
	validMoney := MoneyFromArgs(valid)
	assert.Equal(t, validMoney.Units, int64(1234))
	assert.Equal(t, validMoney.CurrencyCode, "GBP")

	empty := map[string]interface{}{}
	assert.Panics(t, func() {
		MoneyFromArgs(empty)
	})
}
