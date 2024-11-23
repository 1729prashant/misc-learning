package main

import (
	"fmt"
	"testing"
)

func TestHandleEmailBounce(t *testing.T) {
	type testCase struct {
		email           email
		expectedError   string
		expectedStatus  string
		expectedBounces int
	}
	tests := []testCase{
		{
			email: email{
				status:    "email_bounced",
				recipient: &user{email: "bugs@acme.inc"},
			},
			expectedError:   "<nil>",
			expectedStatus:  "email_bounced",
			expectedBounces: 1,
		},
		{
			email: email{
				status:    "email_sent",
				recipient: &user{email: "daffy@acme.inc"},
			},
			expectedError:   "error updating user status: invalid status: email_sent",
			expectedStatus:  "",
			expectedBounces: 0,
		},
	}
	if withSubmit {
		tests = append(tests, testCase{
			email: email{
				status:    "email_failed",
				recipient: &user{email: "porky@acme.inc"},
			},
			expectedError:   "error updating user status: invalid status: email_failed",
			expectedStatus:  "",
			expectedBounces: 0,
		})
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		a := &analytics{}
		err := a.handleEmailBounce(test.email)
		actualError := fmt.Sprintf("%v", err)
		if actualError != test.expectedError {
			failCount++
			t.Errorf(`
---------------------------------
Test Failed:
  status:    %v
  recipient: %v
  expected error:   %v
  actual error:     %v
`, test.email.status, test.email.recipient.email, test.expectedError, actualError)
		} else {
			passCount++
			fmt.Printf(`
---------------------------------
Test Passed:
  status:    %v
  recipient: %v
  expected error:   %v
  actual error:     %v
`, test.email.status, test.email.recipient.email, test.expectedError, actualError)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
