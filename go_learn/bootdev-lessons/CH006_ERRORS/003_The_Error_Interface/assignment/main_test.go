package main

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	type testCase struct {
		dividend, divisor, expected float64
		expectedError               string
	}
	tests := []testCase{
		{10, 2, 5, ""},
		{15, 3, 5, ""},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{10, 0, 0, "can not divide 10 by zero"},
			{15, 0, 0, "can not divide 15 by zero"},
			{100, 10, 10, ""},
			{16, 4, 4, ""},
			{30, 6, 5, ""},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output, err := divide(test.dividend, test.divisor)
		var errString string
		if err != nil {
			errString = err.Error()
		}
		if output != test.expected || errString != test.expectedError {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail`, test.dividend, test.divisor, test.expected, test.expectedError, output, errString)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.dividend, test.divisor, test.expected, test.expectedError, output, errString)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

