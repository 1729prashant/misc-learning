package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		costMultiplier   float64
		maxCostInPennies int
		expected         int
	}
	tests := []testCase{
		{1.1, 5, 4},
		{1.3, 10, 5},
		{1.35, 25, 7},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{1.2, 1, 1},
			{1.2, 15, 7},
			{1.3, 20, 7},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := getMaxMessagesToSend(test.costMultiplier, test.maxCostInPennies)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Fail`, test.costMultiplier, test.maxCostInPennies, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Pass
`, test.costMultiplier, test.maxCostInPennies, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
