package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	test(
		t,
		[]int{1, 2, 3, 4},
		4,
	)
	test(
		t,
		[]string{"a", "b", "c", "d"},
		"d",
	)
	if withSubmit {
		test(
			t,
			[]int{},
			0,
		)
		test(
			t,
			[]bool{true, false, true, true, false},
			false,
		)
	}

}

func test[T comparable](t *testing.T, s []T, expected T) {
	if output := getLast(s); output != expected {
		t.Errorf(`
---------------------------------
Test Failed:
  input:    %v
  expected: %v
  actual:   %v
`,
			s,
			expected,
			output,
		)
	} else {
		fmt.Printf(`
---------------------------------
Test Passed:
  input:    %v
  expected: %v
  actual:   %v
`,
			s,
			expected,
			output,
		)
	}
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
