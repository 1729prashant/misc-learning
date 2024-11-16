package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		tier     string
		expected string
	}
	tests := []testCase{
		{"basic", "You get 1,000 texts per month for $30 per month."},
		{"premium", "You get 50,000 texts per month for $60 per month."},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{"enterprise", "You get unlimited texts per month for $100 per month."},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := getProductMessage(test.tier)
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
