package main

import "testing"

func TestCalculateFullCoverage(t *testing.T) {
	// Defining the test cases (8 paths)
	tests := []struct {
		name       string
		expression string
		want       float64
		wantErr    bool
	}{
		{"Addition", "10 + 5", 15, false},
		{"Subtraction", "20 - 8", 12, false},
		{"Multiplication", "4 * 3", 12, false},
		{"Division", "10 / 2", 5, false},
		{"Division by Zero", "10 / 0", 0, true},
		{"Invalid Format", "5 +", 0, true},
		{"Unknown Operator", "5 % 2", 0, true},
		{"Non-numeric Input", "abc + 2", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.expression)

			// Checking if we expected an error and if we actually got one
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Checking if the numeric result is correct
			if got != tt.want {
				t.Errorf("Calculate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
