package common

func ToCommonMoney(units *int) *Money {
	if units == nil {
		return nil
	}
	return &Money{
		Units: int64(*units),
	}
}
