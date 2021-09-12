package converter

import "testing"

func TestConvert(t *testing.T) {
	var (
		romanArabicPairs RomanArabicPairsType
	)

	romanArabicPairs = RomanArabicPairsType{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9,
		"X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}

	for roman, arabic := range romanArabicPairs {
		var (
			converted int
		)

		converted = Convert(roman)
		if arabic != converted {
			t.Errorf("roman numeral %s expected: %d, in actual: %d", roman, arabic, converted)
		}
	}
}
