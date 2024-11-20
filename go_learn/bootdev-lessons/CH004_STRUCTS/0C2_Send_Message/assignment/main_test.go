package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name           string
		membershipType string
		message        string
		expectResult   string
		expectSuccess  bool
	}{
		{"Syl", "standard", "Hello, Kaladin!", "Hello, Kaladin!", true},
		{"Pattern", "premium", "You are not as good with patterns... You are abstract. You think in lies and tell them to yourselves. That is fascinating, but it is not good for patterns.", "You are not as good with patterns... You are abstract. You think in lies and tell them to yourselves. That is fascinating, but it is not good for patterns.", true},
		{"Dalinar", "standard", "I will take responsibility for what I have done. If I must fall, I will rise each time a better man.", "I will take responsibility for what I have done. If I must fall, I will rise each time a better man.", true},
	}
	if withSubmit {
		submitCases := []struct {
			name           string
			membershipType string
			message        string
			expectResult   string
			expectSuccess  bool
		}{
			{"Pattern", "standard", "Humans can see the world as it is not. It is why your lies can be so strong. You are able to not admit that they are lies.", "", false},
			{"Dabbid", "premium", ".........................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................", "", false},
		}
		tests = append(tests, submitCases...)
	}

	passCount := 0
	failCount := 0

	for _, tc := range tests {
		user := newUser(tc.name, tc.membershipType)
		result, pass := user.SendMessage(tc.message, len(tc.message))
		if tc.expectSuccess != pass || result != tc.expectResult {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
* user:               %s
* membership type:    %s
* message:            %s
* expected result:    %s
* expected success:   %v
* actual result:      %s
* actual success:     %v
`, tc.name, tc.membershipType, tc.message, tc.expectResult, tc.expectSuccess, result, pass)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
* user:               %s
* membership type:    %s
* message:            %s
* expected result:    %s
* expected success:   %v
* actual result:      %s
* actual success:     %v
`, tc.name, tc.membershipType, tc.message, tc.expectResult, tc.expectSuccess, result, pass)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
