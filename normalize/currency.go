package normalize

import "strings"

// Currency check the given code against the ISO 4217 and returns
// the uppercase version if it is valid. It returns EUR if € is passed and
// USD if $ is passed as an argument. All other cases return the given fallback value.
func Currency(given, fallback string) string {
	if given == "" {
		return fallback
	}
	if given == "$" {
		return "USD"
	}
	if given == "€" {
		return "EUR"
	}
	if len(given) != 3 {
		return fallback
	}
	upper := strings.ToUpper(given)
	if _, found := validCurrencies[upper]; found {
		return upper
	}
	return fallback
}

var validCurrencies = map[string]struct{}{
	"AFN": {},
	"ALL": {},
	"AMD": {},
	"ANG": {},
	"AOA": {},
	"ARS": {},
	"AUD": {},
	"AWG": {},
	"AZN": {},
	"BAM": {},
	"BBD": {},
	"BDT": {},
	"BGN": {},
	"BHD": {},
	"BIF": {},
	"BMD": {},
	"BND": {},
	"BOB": {},
	"BOV": {},
	"BRL": {},
	"BSD": {},
	"BTN": {},
	"BWP": {},
	"BYR": {},
	"BZD": {},
	"CAD": {},
	"CDF": {},
	"CHE": {},
	"CHF": {},
	"CHW": {},
	"CLF": {},
	"CLP": {},
	"CNY": {},
	"COP": {},
	"COU": {},
	"CRC": {},
	"CUP": {},
	"CVE": {},
	"CYP": {},
	"CZK": {},
	"DJF": {},
	"DKK": {},
	"DOP": {},
	"DZD": {},
	"EEK": {},
	"EGP": {},
	"ERN": {},
	"ETB": {},
	"EUR": {},
	"FJD": {},
	"FKP": {},
	"GBP": {},
	"GEL": {},
	"GHS": {},
	"GIP": {},
	"GMD": {},
	"GNF": {},
	"GTQ": {},
	"GYD": {},
	"HKD": {},
	"HNL": {},
	"HRK": {},
	"HTG": {},
	"HUF": {},
	"IDR": {},
	"ILS": {},
	"INR": {},
	"IQD": {},
	"IRR": {},
	"ISK": {},
	"JMD": {},
	"JOD": {},
	"JPY": {},
	"KES": {},
	"KGS": {},
	"KHR": {},
	"KMF": {},
	"KPW": {},
	"KRW": {},
	"KWD": {},
	"KYD": {},
	"KZT": {},
	"LAK": {},
	"LBP": {},
	"LKR": {},
	"LRD": {},
	"LSL": {},
	"LTL": {},
	"LVL": {},
	"LYD": {},
	"MAD": {},
	"MDL": {},
	"MGA": {},
	"MKD": {},
	"MMK": {},
	"MNT": {},
	"MOP": {},
	"MRO": {},
	"MTL": {},
	"MUR": {},
	"MVR": {},
	"MWK": {},
	"MXN": {},
	"MXV": {},
	"MYR": {},
	"MZN": {},
	"NAD": {},
	"NGN": {},
	"NIO": {},
	"NOK": {},
	"NPR": {},
	"NZD": {},
	"OMR": {},
	"PAB": {},
	"PEN": {},
	"PGK": {},
	"PHP": {},
	"PKR": {},
	"PLN": {},
	"PYG": {},
	"QAR": {},
	"RON": {},
	"RSD": {},
	"RUB": {},
	"RWF": {},
	"SAR": {},
	"SBD": {},
	"SCR": {},
	"SDG": {},
	"SEK": {},
	"SGD": {},
	"SHP": {},
	"SKK": {},
	"SLL": {},
	"SOS": {},
	"SRD": {},
	"STD": {},
	"SYP": {},
	"SZL": {},
	"THB": {},
	"TJS": {},
	"TMM": {},
	"TND": {},
	"TOP": {},
	"TRY": {},
	"TTD": {},
	"TWD": {},
	"TZS": {},
	"UAH": {},
	"UGX": {},
	"USD": {},
	"XAF": {},
	"XAG": {},
	"XAU": {},
	"XCD": {},
	"XPT": {},
	"XXX": {},
}