package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSendMessage(t *testing.T) {
	tests := []struct {
		formatter Formatter
		expected  string
	}{
		{PlainText{message: "Hello, World!"}, "Hello, World!"},
		{Bold{message: "Bold Message"}, "**Bold Message**"},
		{Code{message: "Code Message"}, "`Code Message`"},
	}

	if withSubmit {
		tests = append(tests,
			struct {
				formatter Formatter
				expected  string
			}{Code{message: ""}, "``"},
			struct {
				formatter Formatter
				expected  string
			}{Bold{message: ""}, "****"},
			struct {
				formatter Formatter
				expected  string
			}{PlainText{message: ""}, ""},
		)
	}

	passCount := 0
	failCount := 0

	for i, test := range tests {
		testName := "Test Case " + strconv.Itoa(i+1)
		t.Run(testName, func(t *testing.T) {
			formattedMessage := SendMessage(test.formatter)
			if formattedMessage != test.expected {
				failCount++
				t.Errorf(`---------------------------------
%s
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail`, testName, test.formatter, test.expected, formattedMessage)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
%s
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, testName, test.formatter, test.expected, formattedMessage)
			}
		})
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

