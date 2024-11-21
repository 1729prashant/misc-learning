package main

import (
	"fmt"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		notification       notification
		expectedID         string
		expectedImportance int
	}{
		{
			directMessage{senderUsername: "Kaladin", messageContent: "Life before death", priorityLevel: 10, isUrgent: true},
			"Kaladin",
			50,
		},
		{
			groupMessage{groupName: "Bridge 4", messageContent: "Soups ready!", priorityLevel: 2},
			"Bridge 4",
			2,
		},
		{
			systemAlert{alertCode: "ALERT001", messageContent: "THIS IS NOT A TEST HIGH STORM COMING SOON"},
			"ALERT001",
			100,
		},
	}
	if withSubmit {
		tests = append(tests,
			struct {
				notification       notification
				expectedID         string
				expectedImportance int
			}{
				directMessage{senderUsername: "Shallan", messageContent: "I am that I am.", priorityLevel: 5, isUrgent: false},
				"Shallan",
				5,
			},
			struct {
				notification       notification
				expectedID         string
				expectedImportance int
			}{
				groupMessage{groupName: "Knights Radiant", messageContent: "For the greater good.", priorityLevel: 10},
				"Knights Radiant",
				10,
			},
			struct {
				notification       notification
				expectedID         string
				expectedImportance int
			}{
				directMessage{senderUsername: "Adolin", messageContent: "Duels are my favorite.", priorityLevel: 3, isUrgent: true},
				"Adolin",
				50,
			},
		)
	}

	passCount := 0
	failCount := 0

	for i, test := range tests {
		t.Run("TestProcessNotification_"+strconv.Itoa(i+1), func(t *testing.T) {
			id, importance := processNotification(test.notification)
			if id != test.expectedID || importance != test.expectedImportance {
				failCount++
				t.Errorf(`---------------------------------
Test Failed: TestProcessNotification_%d
Notification: %+v
Expecting:    %v/%d
Actual:       %v/%d
Fail`, i+1, test.notification, test.expectedID, test.expectedImportance, id, importance)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Test Passed: TestProcessNotification_%d
Notification: %+v
Expecting:    %v/%d
Actual:       %v/%d
Pass
`, i+1, test.notification, test.expectedID, test.expectedImportance, id, importance)
			}
		})
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
