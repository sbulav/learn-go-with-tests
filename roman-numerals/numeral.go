package numerals

import "strings"

import "fmt"

type romanNumeral struct {
	Value  int
	Symbol string
}

type romanNumerals []romanNumeral

var allRomanNumerals = romanNumerals{
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

func (r romanNumerals) ValueOf(symbol string) int {
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func ConvertToArabic(roman string) int {
	total := 0

	fmt.Println("------------------")
	fmt.Println("Converting to arabic", roman)
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		// look ahead to next symbol if we can and, the current symbol is base 10 (only valid subtractors)
		if couldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]

			// build the two character string
			potentialNumber := string([]byte{symbol, nextSymbol})

			if value := allRomanNumerals.ValueOf(potentialNumber); value != 0 {
				total += value
				i++ // move past this character too for the next loop
			} else {
				total++ // this is fishy...
			}
		} else {
			total += allRomanNumerals.ValueOf(string([]byte{symbol}))
		}
		fmt.Println("Total", total)
	}
	return total
}
func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	return index+1 < len(roman) && currentSymbol == 'I'
}
