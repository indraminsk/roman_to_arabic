package converter

import "strings"

type RomanArabicPairsType map[string]int

func Convert(roman string) (arabic int) {
	var (
		ok bool

		arabicNumeralBy RomanArabicPairsType
	)

	arabicNumeralBy = buildRomanArabicNumeralPairs()

	arabic, ok = arabicNumeralBy[strings.ToUpper(roman)]
	if !ok {
		return 0
	}

	return arabic
}

func buildRomanArabicNumeralPairs() (pairs RomanArabicPairsType) {
	pairs = make(RomanArabicPairsType)

	pairs["I"] = 1
	pairs["II"] = 2
	pairs["III"] = 3
	pairs["IV"] = 4
	pairs["V"] = 5
	pairs["VI"] = 6
	pairs["VII"] = 7
	pairs["VIII"] = 8
	pairs["IX"] = 9
	pairs["X"] = 10
	pairs["L"] = 50
	pairs["C"] = 100
	pairs["D"] = 500
	pairs["M"] = 1000

	return pairs
}
