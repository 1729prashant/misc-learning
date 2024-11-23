package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		costs    []cost
		expected []float64
	}
	tests := []testCase{
		{
			costs: []cost{
				{0, 1.0},
				{1, 2.0},
				{1, 3.1},
				{5, 2.5},
				{2, 3.6},
				{1, 2.7},
				{1, 3.3},
			},
			expected: []float64{
				1.0,
				11.1,
				3.6,
				0.0,
				0.0,
				2.5,
			},
		},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{
				costs: []cost{
					{0, 1.0},
					{1, 2.0},
					{1, 3.1},
					{2, 2.5},
					{3, 3.1},
					{3, 2.6},
					{4, 3.34},
				},
				expected: []float64{
					1.0,
					5.1,
					2.5,
					5.7,
					3.34,
				},
			},
			{
				costs: []cost{
					{0, 1.0},
					{10, 2.0},
					{3, 3.1},
					{2, 2.5},
					{1, 3.6},
					{2, 2.7},
					{4, 56.34},
					{13, 2.34},
					{28, 1.34},
					{25, 2.34},
					{30, 4.34},
				},
				expected: []float64{
					1.0,
					3.6,
					5.2,
					3.1,
					56.34,
					0.0,
					0.0,
					0.0,
					0.0,
					0.0,
					2.0,
					0.0,
					0.0,
					2.34,
					0.0,
					0.0,
					0.0,
					0.0,
					0.0,
					0.0,
					0.0,
					0.0,
					0.0,
					0.0,
					0.0,
					2.34,
					0.0,
					0.0,
					1.34,
					0.0,
					4.34,
				},
			},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := getCostsByDay(test.costs)
		if !reflect.DeepEqual(output, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:
%v
Expecting:
%v
Actual:
%v
Fail`, sliceWithBullets(test.costs), sliceWithBullets(test.expected), sliceWithBullets(output))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %v
Expecting:
%v
Actual:
%v
Pass
`, sliceWithBullets(test.costs), sliceWithBullets(test.expected), sliceWithBullets(output))
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
		form := "  - %v\n"
		if i == (len(slice) - 1) {
			form = "  - %v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
