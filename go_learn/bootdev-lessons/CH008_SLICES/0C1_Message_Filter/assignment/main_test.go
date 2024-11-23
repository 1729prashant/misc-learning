package main

import (
	"fmt"
	"testing"
)

func TestFilterMessages(t *testing.T) {
	messages := []Message{
		TextMessage{"Alice", "Hello, World!"},
		MediaMessage{"Bob", "image", "A beautiful sunset"},
		LinkMessage{"Charlie", "http://example.com", "Example Domain"},
		TextMessage{"Dave", "Another text message"},
		MediaMessage{"Eve", "video", "Cute cat video"},
		LinkMessage{"Frank", "https://boot.dev", "Learn Coding Online"},
	}

	testCases := []struct {
		filterType    string
		expectedCount int
		expectedType  string
	}{
		{"text", 2, "text"},
		{"media", 2, "media"},
		{"link", 2, "link"},
	}

	if withSubmit {
		testCases = append(testCases,
			struct {
				filterType    string
				expectedCount int
				expectedType  string
			}{"media", 2, "media"},
			struct {
				filterType    string
				expectedCount int
				expectedType  string
			}{"text", 2, "text"},
		)
	}

	passCount := 0
	failCount := 0

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("TestCase%d", i+1), func(t *testing.T) {
			filtered := filterMessages(messages, tc.filterType)
			if len(filtered) != tc.expectedCount {
				failCount++
				t.Errorf(`---------------------------------
Test Case %d - Filtering for %s
Expecting:  %d messages
Actual:     %d messages
Fail`, i+1, tc.filterType, tc.expectedCount, len(filtered))
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Test Case %d - Filtering for %s
Expecting:  %d messages
Actual:     %d messages
Pass
`, i+1, tc.filterType, tc.expectedCount, len(filtered))
			}

			for _, m := range filtered {
				if m.Type() != tc.expectedType {
					failCount++
					t.Errorf(`---------------------------------
Test Case %d - Message Type Check
Expecting:  %s message
Actual:     %s message
Fail`, i+1, tc.expectedType, m.Type())
				} else {
					passCount++
					fmt.Printf(`---------------------------------
Test Case %d - Message Type Check
Expecting:  %s message
Actual:     %s message
Pass
`, i+1, tc.expectedType, m.Type())
				}
			}
		})
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true
