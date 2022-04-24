package numeral

import (
	"strconv"
	"strings"
)

type RomanNumeral struct {
	Value  uint16
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

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
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

func strToUint16(s string) uint16 {
	n, _ := strconv.ParseUint(s, 10, 16)
	return uint16(n)
}

func ConvertToRoman(arabic string) string {
	number := strToUint16(arabic)

	var result strings.Builder
	for _, numeral := range allRomanNumerals {
		for number >= numeral.Value {
			result.WriteString(numeral.Symbol)
			number -= numeral.Value
		}
	}
	return result.String()
}

func ConvertToArabic(roman string) string {
	total := uint16(0)
	for _, symbol := range romanNumeralString(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbol...)
	}
	return strconv.FormatUint(uint64(total), 10)
}
