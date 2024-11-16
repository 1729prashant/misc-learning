package main

import (
	"fmt"
	"testing"
)

func TestSplitEmail(t *testing.T) {
	type testCase struct {
		email    string
		username string
		domain   string
	}
	tests := []testCase{
		{"drogon@dragonstone.com", "drogon", "dragonstone.com"},
		{"rhaenyra@targaryen.com", "rhaenyra", "targaryen.com"},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{"viserys@kingslanding.com", "viserys", "kingslanding.com"},
			{"aegon@stormsend.com", "aegon", "stormsend.com"},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		username, domain := splitEmail(test.email)
		if username != test.username || domain != test.domain {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail`, test.email, test.username, test.domain, username, domain)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.email, test.username, test.domain, username, domain)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
