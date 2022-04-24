package numeral

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

func (r RomanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
}

type romanNumeralString string

func (s romanNumeralString) Symbols() (symbols [][]byte) {
	for i := 0; i < len(s); i++ {
		symbol := s[i]
		if i+1 < len(s) && isSubtractive(symbol) && allRomanNumerals.Exists(symbol, s[i+1]) {
			symbols = append(symbols, []byte{symbol, s[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return
}

func isSubtractive(symbol byte) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

func ConvertToArabic(roman string) int {
	total := 0
	for _, symbol := range romanNumeralString(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbol...)
	}
	return total
}
