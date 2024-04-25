package app

import (
	"strconv"
)

func arabCalc(n1, n2 int, expression string) string {
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

	return strconv.Itoa(resNum)
}
