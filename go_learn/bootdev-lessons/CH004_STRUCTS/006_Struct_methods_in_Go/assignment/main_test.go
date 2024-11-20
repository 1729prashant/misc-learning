package main

import (
	"fmt"
	"testing"
)

func TestGetBasicAuth(t *testing.T) {
	tests := []struct {
		auth     authenticationInfo
		expected string
	}{
		{authenticationInfo{"Google", "12345"}, "Authorization: Basic Google:12345"},
		{authenticationInfo{"Bing", "98765"}, "Authorization: Basic Bing:98765"},
	}
	if withSubmit {
		tests = append(tests, struct {
			auth     authenticationInfo
			expected string
		}{authenticationInfo{"DDG", "76921"}, "Authorization: Basic DDG:76921"})
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := test.auth.getBasicAuth()
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  %s
Actual:     %s
Fail`, test.auth, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  %s
Actual:     %s
Pass
`, test.auth, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

