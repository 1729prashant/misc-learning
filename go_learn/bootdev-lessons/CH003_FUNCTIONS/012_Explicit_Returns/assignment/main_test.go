package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		age                   int
		exYearsUntilAdult     int
		exYearsUntilDrinking  int
		exYearsUntilCarRental int
	}
	tests := []testCase{
		{4, 14, 17, 21},
		{18, 0, 3, 7},
		{22, 0, 0, 3},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{25, 0, 0, 0},
			{35, 0, 0, 0},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(test.age)
		if yearsUntilAdult != test.exYearsUntilAdult ||
			yearsUntilDrinking != test.exYearsUntilDrinking ||
			yearsUntilCarRental != test.exYearsUntilCarRental {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  (%v, %v, %v)
Actual:     (%v, %v, %v)
Fail`, test.age, test.exYearsUntilAdult, test.exYearsUntilDrinking, test.exYearsUntilCarRental, yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  (%v, %v, %v)
Actual:     (%v, %v, %v)
Pass
`, test.age, test.exYearsUntilAdult, test.exYearsUntilDrinking, test.exYearsUntilCarRental, yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
