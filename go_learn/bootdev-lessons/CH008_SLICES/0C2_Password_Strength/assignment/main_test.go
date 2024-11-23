package main

import (
	"fmt"
	"testing"
)

func TestIsValidPassword(t *testing.T) {
	type testCase struct {
		password string
		isValid  bool
	}
	testCases := []testCase{
		{"Pass123", true},
		{"pas", false},
		{"Password", false},
		{"123456", false},
	}
	if withSubmit {
		testCases = append(testCases,
			[]testCase{
				{"VeryLongPassword1", false},
				{"Short", false},
				{"1234short", false},
				{"Short5", true},
				{"P4ssword", true},
			}...,
		)
	}

	passCount := 0
	failCount := 0

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("TestCase%d", i+1), func(t *testing.T) {
			result := isValidPassword(tc.password)
			if result != tc.isValid {
				failCount++
				t.Errorf(`---------------------------------
Password:  "%s"
Expecting: %v
Actual:    %v
Fail`, tc.password, tc.isValid, result)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Password:  "%s"
Expecting: %v
Actual:    %v
Pass
`, tc.password, tc.isValid, result)
			}
		})
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true

