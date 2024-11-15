package main

import (
	"fmt"
	"testing"
)

func TestSetMessage(t *testing.T) {
	type testCase struct {
		e          email
		newMessage string
		expected   string
	}
	tests := []testCase{
		{
			email{
				message:     "My name is Lt. Aldo Raine and I'm putting together a special team, and I need me eight soldiers.",
				fromAddress: "lt.aldo.raine@mailio.com",
				toAddress:   "army@mailio.com",
			},
			"You just say bingo.",
			"You just say bingo.",
		},
		{
			email{
				message:     "Now, if one were to determine what attribute the German people share with a beast, it would be the cunning and the predatory instinct of a hawk.",
				fromAddress: "col.hans.landa@mailio.com",
				toAddress:   "lapadite@mailio.com",
			},
			"What a tremendously hostile world that a rat must endure.",
			"What a tremendously hostile world that a rat must endure.",
		},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{
				email{
					message:     "Nazi ain't got no humanity. They're the foot soldiers of a Jew-hatin', mass murderin' maniac and they need to be dee-stroyed.",
					fromAddress: "lt.aldo.raine@mailio.com",
					toAddress:   "basterds@mailio.com",
				},
				"I think this just might be my masterpiece.",
				"I think this just might be my masterpiece.",
			},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		originalMessage := test.e.message
		test.e.setMessage(test.newMessage)
		if test.e.message != test.expected {
			failCount++
			t.Errorf(`
---------------------------------
Test Failed:
  inputs:
    * msg: %v
    * newMessage: %v
    * from: %v
    * to: %v
  expected: %v
  actual: %v
`, originalMessage, test.newMessage, test.e.fromAddress, test.e.toAddress, test.expected, test.e.message)
		} else {
			passCount++
			fmt.Printf(`
---------------------------------
Test Passed:
  inputs:
    * msg: %v
    * newMessage: %v
    * from: %v
    * to: %v
  expected: %v
  actual: %v
`, originalMessage, test.newMessage, test.e.fromAddress, test.e.toAddress, test.expected, test.e.message)
		}
	}

	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
