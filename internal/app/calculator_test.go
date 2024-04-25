package app

import (
	"errors"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    string
		wantErr bool
	}{
		{
			name:  "TestCalculate_1, arab `+` OK",
			input: []string{"1", "+", "1"},
			want:  "2",
		},
		{
			name:  "TestCalculate_2, roman `+` OK",
			input: []string{"I", "+", "I"},
			want:  "II",
		},
		{
			name:    "TestCalculate_3, roman `+` Error",
			input:   []string{"I", "-", "I"},
			wantErr: true,
		},
		{
			name:    "TestCalculate_4, Expression Error",
			input:   []string{"1"},
			wantErr: true,
		},
		{
			name:    "TestCalculate_5, Num Validation Error",
			input:   []string{"1", "+", "I"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Calculate(tt.input)

			if err != nil {
				if !tt.wantErr {
					t.Errorf("Unexpected error: %v", err)
					return
				}
			}

			if res != tt.want {
				t.Errorf("Wrong ansver. Want: %q. Got: %q", tt.want, res)
			}
		})
	}
}

func TestExpressionValidation(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		err     error
		wantErr bool
	}{
		{
			name:    "TestExpressionValidation_1, OK",
			input:   []string{"1", "+", "1"},
			err:     nil,
			wantErr: false,
		},
		{
			name:    "TestExpressionValidation_2, Error",
			input:   []string{"I"},
			err:     ErrNoOperation,
			wantErr: true,
		},
		{
			name:    "TestExpressionValidation_3, Error",
			input:   []string{"I", "^", "I"},
			err:     ErrInvalidOperator,
			wantErr: true,
		},
		{
			name:    "TestExpressionValidation_4, Error",
			input:   []string{"I", "+", "I", "+", "I"},
			err:     ErrInvalidExpression,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := expressionValidation(tt.input)

			if err != nil {
				if !tt.wantErr {
					t.Errorf("Unexpected error: %v", err)
					return
				} else {
					if !errors.Is(err, tt.err) {
						t.Errorf("Wrong error. Want: %v. Got: %v", tt.err, err)
					}
					return
				}
			}
		})
	}
}

func TestNumsValidation(t *testing.T) {
	type want struct {
		n1      int
		n2      int
		isRoman bool
		err     error
	}

	tests := []struct {
		name    string
		n1      string
		n2      string
		want    want
		wantErr bool
	}{
		{
			name: "TestNumsValidation_1, OK",
			n1:   "1",
			n2:   "1",
			want: want{
				n1:      1,
				n2:      1,
				isRoman: false,
				err:     nil,
			},
		},
		{
			name: "TestNumsValidation_2, Error",
			n1:   "1",
			n2:   "I",
			want: want{
				n1:      0,
				n2:      0,
				isRoman: false,
				err:     ErrDifferentSystems,
			},
			wantErr: true,
		},
		{
			name: "TestNumsValidation_3, Error",
			n1:   "1",
			n2:   "0",
			want: want{
				n1:      0,
				n2:      0,
				isRoman: false,
				err:     ErrInvalidNumber,
			},
			wantErr: true,
		},
		{
			name: "TestNumsValidation_4, Error",
			n1:   "0",
			n2:   "1",
			want: want{
				n1:      0,
				n2:      0,
				isRoman: false,
				err:     ErrInvalidNumber,
			},
			wantErr: true,
		},
		{
			name: "TestNumsValidation_5, OK",
			n1:   "I",
			n2:   "I",
			want: want{
				n1:      1,
				n2:      1,
				isRoman: true,
			},
		},
		{
			name: "TestNumsValidation_6, Error",
			n1:   "I",
			n2:   "IC",
			want: want{
				n1:      0,
				n2:      0,
				isRoman: true,
				err:     ErrInvalidNumber,
			},
			wantErr: true,
		},
		{
			name: "TestNumsValidation_7, Error",
			n1:   "IC",
			n2:   "I",
			want: want{
				n1:      0,
				n2:      0,
				isRoman: true,
				err:     ErrInvalidNumber,
			},
			wantErr: true,
		},
		{
			name: "TestNumsValidation_8, Error",
			n1:   "XX",
			n2:   "I",
			want: want{
				n1:      0,
				n2:      0,
				isRoman: true,
				err:     ErrInvalidNumber,
			},
			wantErr: true,
		},
		{
			name: "TestNumsValidation_9, Error",
			n1:   "I",
			n2:   "XX",
			want: want{
				n1:      0,
				n2:      0,
				isRoman: true,
				err:     ErrInvalidNumber,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n1, n2, isRoman, err := numsValidation(tt.n1, tt.n2)

			if err != nil {
				if !tt.wantErr {
					t.Errorf("Unexpected error: %v", err)
					return
				} else {
					if !errors.Is(err, tt.want.err) {
						t.Errorf("Wrong error. Want: %v. Got: %v", tt.want.err, err)
					}
					return
				}
			}

			if n1 != tt.want.n1 {
				t.Errorf("Wrong ansver. Want: %d. Got: %d", tt.want.n1, n1)
			}

			if n2 != tt.want.n2 {
				t.Errorf("Wrong ansver. Want: %d. Got: %d", tt.want.n2, n2)
			}

			if isRoman != tt.want.isRoman {
				t.Errorf("Wrong ansver. Want: %t. Got: %t", tt.want.isRoman, isRoman)
			}
		})
	}
}

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		err     error
		wantErr bool
	}{
		{
			name:  "TestRomanToInt_1, OK",
			input: "I",
			want:  1,
		},
		{
			name:    "TestRomanToInt_2, Error",
			input:   "IC",
			err:     ErrInvalidNumber,
			wantErr: true,
		},
		{
			name:  "TestRomanToInt_3, OK",
			input: "IV",
			want:  4,
		},
		{
			name:  "TestRomanToInt_4, OK",
			input: "VI",
			want:  6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := romanToInt(tt.input)

			if err != nil {
				if !tt.wantErr {
					t.Errorf("Unexpected error: %v", err)
					return
				} else {
					if !errors.Is(err, tt.err) {
						t.Errorf("Wrong error. Want: %v. Got: %v", tt.err, err)
					}
					return
				}
			}

			if res != tt.want {
				t.Errorf("Wrong ansver. Want: %q. Got: %q", tt.want, res)
			}
		})
	}
}

func TestArabCalc(t *testing.T) {
	tests := []struct {
		name string
		n1   int
		n2   int
		exp  string
		want string
	}{
		{
			name: "TestArabCalc_1, `+` OK",
			n1:   1,
			n2:   1,
			exp:  "+",
			want: "2",
		},
		{
			name: "TestArabCalc_2, `-` OK",
			n1:   1,
			n2:   1,
			exp:  "-",
			want: "0",
		},
		{
			name: "TestArabCalc_3, `*` OK",
			n1:   1,
			n2:   1,
			exp:  "*",
			want: "1",
		},
		{
			name: "TestArabCalc_4, `/` OK",
			n1:   1,
			n2:   1,
			exp:  "/",
			want: "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := arabCalc(tt.n1, tt.n2, tt.exp)

			if res != tt.want {
				t.Errorf("Wrong ansver. Want: %q. Got: %q", tt.want, res)
			}
		})
	}
}

func TestRomanCalc(t *testing.T) {
	tests := []struct {
		name    string
		n1      int
		n2      int
		exp     string
		want    string
		err     error
		wantErr bool
	}{
		{
			name: "TestRomanCalc_1, `+` OK",
			n1:   1,
			n2:   1,
			exp:  "+",
			want: "II",
		},
		{
			name: "TestRomanCalc_2, `-` OK",
			n1:   2,
			n2:   1,
			exp:  "-",
			want: "I",
		},
		{
			name: "TestRomanCalc_3, `*` OK",
			n1:   1,
			n2:   2,
			exp:  "*",
			want: "II",
		},
		{
			name: "TestRomanCalc_4, `/` OK",
			n1:   1,
			n2:   1,
			exp:  "/",
			want: "I",
		},
		{
			name:    "TestRomanCalc_5, Error",
			n1:      1,
			n2:      1,
			exp:     "-",
			err:     ErrInvalidRoman,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := romanCalc(tt.n1, tt.n2, tt.exp)

			if err != nil {
				if !tt.wantErr {
					t.Errorf("Unexpected error: %v", err)
					return
				} else {
					if !errors.Is(err, tt.err) {
						t.Errorf("Wrong error. Want: %v. Got: %v", tt.err, err)
					}
					return
				}
			}

			if res != tt.want {
				t.Errorf("Wrong ansver. Want: %q. Got: %q", tt.want, res)
			}
		})
	}
}
