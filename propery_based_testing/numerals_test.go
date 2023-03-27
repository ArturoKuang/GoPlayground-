package roman_numerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []romanNumeral{
	{Value: 1, Symbol: "I"},
	{Value: 2, Symbol: "II"},
	{Value: 3, Symbol: "III"},
	{Value: 4, Symbol: "IV"},
	{Value: 5, Symbol: "V"},
	{Value: 6, Symbol: "VI"},
	{Value: 7, Symbol: "VII"},
	{Value: 8, Symbol: "VIII"},
	{Value: 9, Symbol: "IX"},
	{Value: 10, Symbol: "X"},
	{Value: 14, Symbol: "XIV"},
	{Value: 18, Symbol: "XVIII"},
	{Value: 20, Symbol: "XX"},
	{Value: 39, Symbol: "XXXIX"},
	{Value: 40, Symbol: "XL"},
	{Value: 47, Symbol: "XLVII"},
	{Value: 49, Symbol: "XLIX"},
	{Value: 50, Symbol: "L"},
	{Value: 100, Symbol: "C"},
	{Value: 90, Symbol: "XC"},
	{Value: 400, Symbol: "CD"},
	{Value: 500, Symbol: "D"},
	{Value: 900, Symbol: "CM"},
	{Value: 1000, Symbol: "M"},
	{Value: 1984, Symbol: "MCMLXXXIV"},
	{Value: 3999, Symbol: "MMMCMXCIX"},
	{Value: 2014, Symbol: "MMXIV"},
	{Value: 1006, Symbol: "MVI"},
	{Value: 798, Symbol: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Value, test.Symbol), func(t *testing.T) {
			got := ConvertToRoman(test.Value)
			if got != test.Symbol {
				t.Errorf("got %q, want %q", got, test.Symbol)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Symbol, test.Value), func(t *testing.T) {
			got := ConvertToArabic(test.Symbol)
			if got != test.Value {
				t.Errorf("got %d, want %d", got, test.Value)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}

		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
