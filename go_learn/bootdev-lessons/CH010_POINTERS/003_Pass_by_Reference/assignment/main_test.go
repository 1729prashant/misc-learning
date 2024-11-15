package main

import (
	"fmt"
	"testing"
)

func TestGetMessageText(t *testing.T) {
	type testCase struct {
		initialAnalytics Analytics
		newMessage       Message
		expected         Analytics
	}

	tests := []testCase{
		{
			initialAnalytics: Analytics{MessagesTotal: 0, MessagesFailed: 0, MessagesSucceeded: 0},
			newMessage:       Message{Recipient: "mickey", Success: true},
			expected:         Analytics{MessagesTotal: 1, MessagesFailed: 0, MessagesSucceeded: 1},
		},
		{
			initialAnalytics: Analytics{MessagesTotal: 1, MessagesFailed: 0, MessagesSucceeded: 1},
			newMessage:       Message{Recipient: "minnie", Success: false},
			expected:         Analytics{MessagesTotal: 2, MessagesFailed: 1, MessagesSucceeded: 1},
		},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{
				initialAnalytics: Analytics{MessagesTotal: 2, MessagesFailed: 1, MessagesSucceeded: 1},
				newMessage:       Message{Recipient: "goofy", Success: false},
				expected:         Analytics{MessagesTotal: 3, MessagesFailed: 2, MessagesSucceeded: 1},
			},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		a := test.initialAnalytics
		getMessageText(&a, test.newMessage)
		if a != test.expected {
			failCount++
			t.Errorf(`
---------------------------------
Test Failed:
  Initial Analytics:
    MessagesTotal=%d, MessagesFailed=%d, MessagesSucceeded=%d
  New Message:
    Recipient=%s, Success=%v
  Expected:
    MessagesTotal=%d, MessagesFailed=%d, MessagesSucceeded=%d
  Actual:
    MessagesTotal=%d, MessagesFailed=%d, MessagesSucceeded=%d
`, test.initialAnalytics.MessagesTotal, test.initialAnalytics.MessagesFailed, test.initialAnalytics.MessagesSucceeded,
				test.newMessage.Recipient, test.newMessage.Success,
				test.expected.MessagesTotal, test.expected.MessagesFailed, test.expected.MessagesSucceeded,
				a.MessagesTotal, a.MessagesFailed, a.MessagesSucceeded)
		} else {
			passCount++
			fmt.Printf(`
---------------------------------
Test Passed:
  Initial Analytics:
    MessagesTotal=%d, MessagesFailed=%d, MessagesSucceeded=%d
  New Message:
    Recipient=%s, Success=%v
  Expected:
    MessagesTotal=%d, MessagesFailed=%d, MessagesSucceeded=%d
  Actual:
    MessagesTotal=%d, MessagesFailed=%d, MessagesSucceeded=%d
`, test.initialAnalytics.MessagesTotal, test.initialAnalytics.MessagesFailed, test.initialAnalytics.MessagesSucceeded,
				test.newMessage.Recipient, test.newMessage.Success,
				test.expected.MessagesTotal, test.expected.MessagesFailed, test.expected.MessagesSucceeded,
				a.MessagesTotal, a.MessagesFailed, a.MessagesSucceeded)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
