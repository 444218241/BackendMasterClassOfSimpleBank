package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// IsSupportedCurrency return true if the currency is supported.
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case EUR, USD:
		return true
	}
	return false
}
