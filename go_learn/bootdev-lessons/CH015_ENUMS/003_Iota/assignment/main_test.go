package main

import (
	"fmt"
	"testing"
)

func TestEmailStatus(t *testing.T) {
	type testCase struct {
		status   emailStatus
		expected string
	}
	tests := []testCase{
		{emailBounced, "emailBounced"},
		{emailInvalid, "emailInvalid"},
		{emailDelivered, "emailDelivered"},
	}
	if withSubmit {
		tests = append(tests, testCase{emailOpened, "emailOpened"})
		tests = append(tests, testCase{17, "Unknown"})
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := getEmailStatusName(test.status)
		if output != test.expected {
			failCount++
			t.Errorf(`
---------------------------------
Test Failed:
  status:   %v
  expected: %v
  actual:   %v
`, test.status, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`
---------------------------------
Test Passed:
  status:   %v
  expected: %v
  actual:   %v
`, test.status, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func getEmailStatusName(status emailStatus) string {
	switch status {
	case emailBounced:
		return "emailBounced"
	case emailInvalid:
		return "emailInvalid"
	case emailDelivered:
		return "emailDelivered"
	case emailOpened:
		return "emailOpened"
	default:
		return "Unknown"
	}
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
