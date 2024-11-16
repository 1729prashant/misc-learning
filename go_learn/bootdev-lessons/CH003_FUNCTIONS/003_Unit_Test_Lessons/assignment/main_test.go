package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		tier     string
		expected int
	}
	tests := []testCase{
		{"basic", 10000},
		{"premium", 15000},
		{"enterprise", 50000},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{"invalid", 0},
			{"", 0},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := getMonthlyPrice(test.tier)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail`, test.tier, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, test.tier, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

