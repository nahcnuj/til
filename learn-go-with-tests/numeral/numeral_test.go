package numeral

import (
	"fmt"
	"strconv"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Number uint16
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{19, "XIX"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func uint16ToStr(u uint16) string {
	return strconv.FormatUint(uint64(u), 10)
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		arabic := uint16ToStr(test.Number)
		t.Run(fmt.Sprintf("%q gets converted to %q", uint16ToStr(test.Number), test.Roman), func(t *testing.T) {
			got := ConvertToRoman(arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		want := uint16ToStr(test.Number)
		t.Run(fmt.Sprintf("%q gets converted to %q", test.Roman, want), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(number uint16) bool {
		if number > 3999 {
			return true
		}
		arabic := uint16ToStr(number)
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
