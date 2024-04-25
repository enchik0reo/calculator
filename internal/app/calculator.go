package app

func Calculate(exp []string) (string, error) {
	if err := expressionValidation(exp); err != nil {
		return "", err
	}

	num1, num2, isRoman, err := numsValidation(exp[0], exp[2])
	if err != nil {
		return "", err
	}

	var res string

	if isRoman {
		res, err = romanCalc(num1, num2, exp[1])
		if err != nil {
			return "", err
		}
	} else {
		res = arabCalc(num1, num2, exp[1])
	}

	return res, nil
}
