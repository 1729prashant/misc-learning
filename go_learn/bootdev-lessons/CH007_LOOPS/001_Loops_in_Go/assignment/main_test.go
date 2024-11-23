package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		numMessages int
		expected    float64
	}
	tests := []testCase{
		{10, 10.45},
		{20, 21.9},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{0, 0.0},
			{1, 1.0},
			{5, 5.10},
			{30, 34.35},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := bulkSend(test.numMessages)
		if fmt.Sprintf("%.2f", output) != fmt.Sprintf("%.2f", test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %.2f
Actual:     %.2f
Fail`, test.numMessages, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %.2f
Actual:     %.2f
Pass
`, test.numMessages, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
