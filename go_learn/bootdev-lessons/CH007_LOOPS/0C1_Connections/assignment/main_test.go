package main

import (
	"fmt"
	"testing"
)

func TestCountGroupConnections(t *testing.T) {
	type testCase struct {
		groupSize int
		expected  int
	}
	testCases := []testCase{
		{1, 0},
		{2, 1},
		{3, 3},
		{4, 6},
	}
	if withSubmit {
		testCases = append(testCases,
			[]testCase{
				{0, 0},
				{10, 45},
				{100, 4950},
				{1000, 499500},
			}...,
		)
	}

	passCount := 0
	failCount := 0

	for _, tc := range testCases {
		result := countConnections(tc.groupSize)
		if result != tc.expected {
			failCount++
			t.Errorf(`---------------------------------
Group Size: %v
Expecting: %v
Actual:    %v
Fail`, tc.groupSize, tc.expected, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Group Size: %v
Expecting: %v
Actual:    %v
Pass
`, tc.groupSize, tc.expected, result)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true

