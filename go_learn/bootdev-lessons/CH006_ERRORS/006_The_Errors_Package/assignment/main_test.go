package main

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	type testCase struct {
		x, y, expected float64
		expectedErr    string
	}

	testCases := []testCase{
		{10, 0, 0, "no dividing by 0"},
		{10, 2, 5, ""},
		{15, 30, 0.5, ""},
		{6, 3, 2, ""},
	}
	if withSubmit {
		testCases = append(testCases,
			testCase{0, 10, 0, ""},
			testCase{100, 0, 0, "no dividing by 0"},
			testCase{-10, -2, 5, ""},
			testCase{-10, 2, -5, ""},
		)
	}

	passCount := 0
	failCount := 0

	for _, tc := range testCases {
		result, err := divide(tc.x, tc.y)
		errString := ""
		if err != nil {
			errString = err.Error()
		}
		if result != tc.expected || errString != tc.expectedErr {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail`, tc.x, tc.y, tc.expected, tc.expectedErr, result, errString)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, tc.x, tc.y, tc.expected, tc.expectedErr, result, errString)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true
