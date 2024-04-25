package app

func expressionValidation(exp []string) error {
	if len(exp) < 3 {
		return ErrNoOperation
	}

	if len(exp) > 3 {
		return ErrInvalidExpression
	}

	if exp[1] != "+" && exp[1] != "-" && exp[1] != "*" && exp[1] != "/" {
		return ErrInvalidOperator
	}

	return nil
}
