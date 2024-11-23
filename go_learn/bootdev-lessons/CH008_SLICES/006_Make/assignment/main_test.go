package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		messages    []string
		expected    []float64
		expectedCap int
	}
	tests := []testCase{
		{
			[]string{"Welcome to the movies!", "Enjoy your popcorn!"},
			[]float64{0.22, 0.19},
			2,
		},
		{
			[]string{"I don't want to be here anymore", "Can we go home?", "I'm hungry", "I'm bored"},
			[]float64{0.31, 0.15, 0.1, 0.09},
			4,
		},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{[]string{}, []float64{}, 0},
			{[]string{""}, []float64{0}, 1},
			{[]string{"Hello", "Hi", "Hey"}, []float64{0.05, 0.02, 0.03}, 3},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := getMessageCosts(test.messages)
		if !slicesEqual(output, test.expected) || cap(output) != test.expectedCap {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
%v
Expecting:
%v
expected cap: %v
Actual:
%v
actual cap: %v
Fail
`, sliceWithBullets(test.messages), sliceWithBullets(test.expected), test.expectedCap, sliceWithBullets(output), cap(output))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
%v
Expecting:
%v
expected cap: %v
Actual:
%v
actual cap: %v
Pass
`, sliceWithBullets(test.messages), sliceWithBullets(test.expected), test.expectedCap, sliceWithBullets(output), cap(output))
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func sliceWithBullets[T any](slice []T) string {
	if slice == nil {
		return "  <nil>"
	}
	if len(slice) == 0 {
		return "  []"
	}
	output := ""
	for i, item := range slice {
		form := "  - %#v\n"
		if i == (len(slice) - 1) {
			form = "  - %#v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}

func slicesEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
