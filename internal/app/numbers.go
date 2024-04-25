package app

import "strconv"

func numsValidation(n1, n2 string) (int, int, bool, error) {
	isArab1, isArab2 := false, false

	arab1, err := strconv.Atoi(n1)
	if err == nil {
		isArab1 = true
	}

	arab2, err := strconv.Atoi(n2)
	if err == nil {
		isArab2 = true
	}

	if isArab1 != isArab2 {
		return 0, 0, false, ErrDifferentSystems
	}

	if isArab1 {
		if valid := numberIsValid(arab1); !valid {
			return 0, 0, false, ErrInvalidNumber
		}

		if valid := numberIsValid(arab2); !valid {
			return 0, 0, false, ErrInvalidNumber
		}

		return arab1, arab2, false, nil
	}

	arab1, err = romanToInt(n1)
	if err != nil {
		return 0, 0, true, err
	}

	arab2, err = romanToInt(n2)
	if err != nil {
		return 0, 0, true, err
	}

	if valid := numberIsValid(arab1); !valid {
		return 0, 0, true, ErrInvalidNumber
	}

	if valid := numberIsValid(arab2); !valid {
		return 0, 0, true, ErrInvalidNumber
	}

	return arab1, arab2, true, nil
}

func romanToInt(s string) (int, error) {
	res := 0
	romanNums := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	for i := len(s) - 1; i >= 0; i-- {
		ni, ok := romanNums[s[i]]
		if !ok {
			return 0, ErrInvalidNumber
		}

		if i == len(s)-1 {
			res += ni
			continue
		}

		if ni >= romanNums[s[i+1]] {
			res += ni
		} else {
			res -= ni
		}
	}

	return res, nil
}

func numberIsValid(n int) bool {
	if n < 1 || n > 10 {
		return false
	}

	return true
}
