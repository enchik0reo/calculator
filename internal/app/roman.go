package app

import (
	"strings"
)

func romanCalc(n1, n2 int, expression string) (string, error) {
	var resNum int

	switch expression {
	case "+":
		resNum = n1 + n2
	case "-":
		resNum = n1 - n2
	case "*":
		resNum = n1 * n2
	case "/":
		resNum = n1 / n2
	}

	if resNum <= 0 {
		return "", ErrInvalidRoman
	}

	roman := arabToRoman(resNum)

	return roman, nil
}

func arabToRoman(arab int) string {
	romanNums := []struct {
		num    int
		symbol string
	}{
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

	var result strings.Builder

	for _, rn := range romanNums {
		for arab >= rn.num {
			result.WriteString(rn.symbol)
			arab -= rn.num
		}
	}

	return result.String()
}
