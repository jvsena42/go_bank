package util

const (
	USD = "USD"
	CAD = "CAD"
	EUR = "EUR"
	BTC = "BTC"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, CAD, EUR, BTC:
		return true
	}
	return false
}
