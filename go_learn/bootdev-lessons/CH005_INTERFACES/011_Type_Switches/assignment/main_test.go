package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		expense      expense
		expectedTo   string
		expectedCost float64
	}
	tests := []testCase{
		{
			email{true, "hello there", "kit@boga.com"},
			"kit@boga.com",
			0.11,
		},
		{
			sms{false, "I'm a Nigerian prince, please send me your bank info so I can deposit $1000 dollars", "+155555509832"},
			"+155555509832",
			8.3,
		},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{invalid{}, "", 0},
			{
				email{false, "This meeting could have been an email", "jane@doe.com"},
				"jane@doe.com",
				1.85,
			},
			{
				sms{false, "Please sir/madam", "+155555504444"},
				"+155555504444",
				1.6,
			},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		to, cost := getExpenseReport(test.expense)
		if to != test.expectedTo || cost != test.expectedCost {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail`, test.expense, test.expectedTo, test.expectedCost, to, cost)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.expense, test.expectedTo, test.expectedCost, to, cost)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

