package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		messages []string
		expected int
	}{
		{
			[]string{"WTS Arcanite Bar! Cheaper than AH", "Do you need an Arcanite Bar!"},
			10,
		},
		{
			[]string{"Could you give me a number crunch real quick?", "Looks like we have a 32.33% (repeating of course) percentage of survival."},
			19,
		},
		{
			[]string{"LFG UBRS", "lfg ubrs", "LFG Ubrs"},
			2,
		},
	}

	if withSubmit {
		testCases = append(testCases, struct {
			messages []string
			expected int
		}{
			[]string{"Alright time's up! Let's do this.", "Leroy Jenkins!", "Damn it Leroy"},
			10,
		}, struct {
			messages []string
			expected int
		}{
			[]string{"I'm out of range", "I'm out of mana"},
			5,
		},
			struct {
				messages []string
				expected int
			}{
				[]string{
					"LF9M UBRS need all",
					"LF8M UBRS need all",
					"LF7M UBRS need all",
					"LF6M UBRS need tanks and heals",
					"LF5M UBRS need tanks and heals",
					"LF4M UBRS need tanks and heals",
					"LF3M UBRS need tanks and healer",
					"LF2M UBRS need tanks",
					"LF1M UBRS need tank",
					"Group is full thanks!",
				},
				21,
			},
			struct {
				messages []string
				expected int
			}{
				[]string{""},
				0,
			})
	}

	passCount := 0
	failCount := 0

	for _, tc := range testCases {
		result := countDistinctWords(tc.messages)
		formattedMessages := formatMessages(tc.messages)
		if result != tc.expected {
			failCount++
			t.Errorf(`---------------------------------
FAIL:
Messages: %v
Expecting: %d unique words
Actual:    %d unique words
Fail
`, formattedMessages, tc.expected, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed!
Messages: %v
Expecting: %d unique words
Actual:    %d unique words
Pass
`, formattedMessages, tc.expected, result)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func formatMessages(messages []string) string {
	var formattedMessages []string
	for _, message := range messages {
		words := strings.Fields(message)
		formattedMessage := strings.Join(words, " ")
		formattedMessages = append(formattedMessages, fmt.Sprintf("[%s]", formattedMessage))
	}
	return strings.Join(formattedMessages, ", ")
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
