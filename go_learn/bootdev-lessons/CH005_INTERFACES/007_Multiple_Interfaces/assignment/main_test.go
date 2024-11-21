package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		body           string
		isSubscribed   bool
		expectedCost   int
		expectedFormat string
	}
	tests := []testCase{
		{"hello there", true, 22, "'hello there' | Subscribed"},
		{"general kenobi", false, 70, "'general kenobi' | Not Subscribed"},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{"i hate sand", true, 22, "'i hate sand' | Subscribed"},
			{"it's coarse and rough and irritating", false, 180, "'it's coarse and rough and irritating' | Not Subscribed"},
			{"and it gets everywhere", true, 44, "'and it gets everywhere' | Subscribed"},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		e := email{
			body:         test.body,
			isSubscribed: test.isSubscribed,
		}
		cost := e.cost()
		format := e.format()
		if format != test.expectedFormat || cost != test.expectedCost {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail`, test.body, test.isSubscribed, test.expectedCost, test.expectedFormat, cost, format)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.body, test.isSubscribed, test.expectedCost, test.expectedFormat, cost, format)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
