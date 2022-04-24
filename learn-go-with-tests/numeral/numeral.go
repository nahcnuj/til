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

type RomanNumeralString string

func (s RomanNumeralString) Symbols() (symbols [][]byte) {
	for i := 0; i < len(s); i++ {
		symbol := s[i]
		if couldBeSubtractive(i, symbol, string(s)) {
			if allRomanNumerals.Exists(symbol, s[i+1]) {
				symbols = append(symbols, []byte{symbol, s[i+1]})
				i++
			} else {
				symbols = append(symbols, []byte{symbol})
			}
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return
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
	for _, symbol := range RomanNumeralString(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbol...)
	}
	return total
}

func isSubtractive(symbol byte) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}

func couldBeSubtractive(index int, currentSymbol byte, roman string) bool {
	return index+1 < len(roman) && isSubtractive(currentSymbol)
}
