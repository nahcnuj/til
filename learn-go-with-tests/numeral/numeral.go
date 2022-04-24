package numeral

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var breakpoints = []RomanNumeral{
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	for _, breakpoint := range breakpoints {
		for arabic >= breakpoint.Value {
			result.WriteString(breakpoint.Symbol)
			arabic -= breakpoint.Value
		}
	}
	return result.String()
}
