package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		recipient string
		text      string
		expected  string
	}
	tests := []testCase{
		{
			recipient: "Honey Bunny",
			text:      "I love you, Pumpkin.",
			expected: `
To: Honey Bunny
Message: I love you, Pumpkin.
`,
		},
		{
			recipient: "Pumpkin",
			text:      "I love you, Honey Bunny.",
			expected: `
To: Pumpkin
Message: I love you, Honey Bunny.
`,
		},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{
				recipient: "poor sap 1",
				text:      "And you will know My name is the Lord when I lay My vengeance upon thee.",
				expected: `
To: poor sap 1
Message: And you will know My name is the Lord when I lay My vengeance upon thee.
`,
			},
			{
				recipient: "Fabienne",
				text:      "Zed's dead, baby. Zed's dead.",
				expected: `
To: Fabienne
Message: Zed's dead, baby. Zed's dead.
`,
			},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		m := Message{Recipient: test.recipient, Text: test.text}
		messageText := getMessageText(m)
		if strings.TrimSpace(messageText) != strings.TrimSpace(test.expected) {
			failCount++
			t.Errorf(`Test Failed:
input:
* Recipient: %v
* Text: %v
expected:
%v
actual:
%v
`,
				m.Recipient,
				m.Text,
				test.expected,
				messageText,
			)
		} else {
			passCount++
			fmt.Printf(`Test Passed:
input:
* Recipient: %v
* Text: %v
expected:
%v
actual:
%v
`,
				m.Recipient,
				m.Text,
				test.expected,
				messageText,
			)
		}
		fmt.Println("------------------------------")
	}

	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

