package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		name     string
		expected uintptr
	}

	tests := []testCase{
		{"contact", uintptr(24)},
	}
	if withSubmit {
		tests = append(tests, testCase{"perms", uintptr(16)})
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		var typ reflect.Type
		if test.name == "contact" {
			typ = reflect.TypeOf(contact{})
		} else if test.name == "perms" {
			typ = reflect.TypeOf(perms{})
		}

		size := typ.Size()

		if size != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %v bytes
Actual:     %v bytes
Fail`, test.name, test.expected, size)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v bytes
Actual:     %v bytes
Pass
`, test.name, test.expected, size)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

