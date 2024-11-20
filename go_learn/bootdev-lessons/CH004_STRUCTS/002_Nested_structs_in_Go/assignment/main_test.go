package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		mToSend  messageToSend
		expected bool
	}
	tests := []testCase{
		{messageToSend{
			message:   "you have an appointment tomorrow",
			sender:    user{name: "Brenda Halafax", number: 16545550987},
			recipient: user{name: "Sally Sue", number: 19035558973},
		}, true},
		{messageToSend{
			message:   "you have an event tomorrow",
			sender:    user{number: 16545550987},
			recipient: user{name: "Suzie Sall", number: 19035558973},
		}, false},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{messageToSend{
				message:   "you have an birthday tomorrow",
				sender:    user{name: "Jason Bjorn", number: 16545550987},
				recipient: user{name: "Jim Bond"},
			}, false},
			{messageToSend{
				message:   "you have a party tomorrow",
				sender:    user{name: "Njorn Halafax"},
				recipient: user{name: "Becky Sue", number: 19035558973},
			}, false},
			{messageToSend{
				message:   "you have a birthday tomorrow",
				sender:    user{name: "Eli Halafax", number: 16545550987},
				recipient: user{number: 19035558973},
			}, false},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := canSendMessage(test.mToSend)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Test Failed. Inputs:
* message:          %s
* sender.name:      %s
* sender.number:    %d
* recipient.name:   %s
* recipient.number: %d
Expected:           %v
Actual:             %v
`,
				test.mToSend.message,
				test.mToSend.sender.name,
				test.mToSend.sender.number,
				test.mToSend.recipient.name,
				test.mToSend.recipient.number,
				test.expected,
				output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed. Inputs:
* message:          %s
* sender.name:      %s
* sender.number:    %d
* recipient.name:   %s
* recipient.number: %d
Expected:           %v
Actual:             %v
`,
				test.mToSend.message,
				test.mToSend.sender.name,
				test.mToSend.sender.number,
				test.mToSend.recipient.name,
				test.mToSend.recipient.number,
				test.expected,
				output)
		}
	}

	fmt.Printf("---------------------------------\n")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

