package main

import (
	"fmt"
	"slices"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		plan             string
		messages         [3]string
		expectedMessages []string
		expectedErr      string
	}
	tests := []testCase{
		{
			planFree,
			[3]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[]string{"Hello sir/madam can I interest you in a yacht?", "Please I'll even give you an Amazon gift card?"},
			"",
		},
		{
			planPro,
			[3]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			"",
		},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{
				planFree,
				[3]string{
					"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
					"Wouldn't you?",
					"Wouldn't you???",
				},
				[]string{
					"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
					"Wouldn't you?",
				},
				"",
			},
			{
				planPro,
				[3]string{
					"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
					"Wouldn't you?",
					"Wouldn't you???",
				},
				[]string{
					"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
					"Wouldn't you?",
					"Wouldn't you???",
				},
				"",
			},
			{
				"invalid plan",
				[3]string{
					"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
					"Wouldn't you?",
					"Wouldn't you???",
				},
				nil,
				"unsupported plan",
			},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		actualMessages, err := getMessageWithRetriesForPlan(test.plan, test.messages)
		errString := ""
		if err != nil {
			errString = err.Error()
		}
		if !slices.Equal(actualMessages, test.expectedMessages) || errString != test.expectedErr {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
Plan: %v
Messages:
%v
Expecting:
%v
errString:  %v
Actual:
%v
errString:  %v
Fail
`, test.plan, sliceWithBullets(test.messages[:]), sliceWithBullets(test.expectedMessages), test.expectedErr, sliceWithBullets(actualMessages), errString)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
Plan: %v
Messages:
%v
Expecting:
%v
errString:  %v
Actual:
%v
errString:  %v
Pass
`, test.plan, sliceWithBullets(test.messages[:]), sliceWithBullets(test.expectedMessages), test.expectedErr, sliceWithBullets(actualMessages), errString)
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

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
