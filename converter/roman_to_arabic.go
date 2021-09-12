package converter

import "strings"

type RomanArabicPairsType map[string]int

func Convert(roman string) (arabic int) {
	var (
		ok bool

		romanArabicPairs RomanArabicPairsType
	)

	romanArabicPairs = RomanArabicPairsType{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9,
		"X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}

	arabic, ok = romanArabicPairs[strings.ToUpper(roman)]
	if !ok {
		return 0
	}

	return arabic
}
