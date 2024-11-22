package main

import (
	"fmt"
	"testing"
)

func TestLogAndDelete(t *testing.T) {
	type testCase struct {
		users         map[string]user
		name          string
		expectedUsers map[string]user
		expectedLog   string
	}
	tests := []testCase{
		{
			users: map[string]user{
				"laura": {name: "laura", number: 4355556023, admin: false},
				"dale":  {name: "dale", number: 8015558937, admin: true},
			},
			name: "laura",
			expectedUsers: map[string]user{
				"dale": {name: "dale", number: 8015558937, admin: true},
			},
			expectedLog: logDeleted,
		},
		{
			users: map[string]user{
				"audrey": {name: "audrey", number: 4355556024, admin: false},
				"bob":    {name: "bob", number: 8015558938, admin: true},
				"bobby":  {name: "bobby", number: 8015558939, admin: true},
			},
			name: "bobby",
			expectedUsers: map[string]user{
				"audrey": {name: "audrey", number: 4355556024, admin: false},
				"bob":    {name: "bob", number: 8015558938, admin: true},
			},
			expectedLog: logAdmin,
		},
	}

	if withSubmit {
		tests = append(tests, testCase{
			users: map[string]user{
				"log lady": {name: "log lady", number: 4355556025, admin: false},
				"shelly":   {name: "shelly", number: 8015558940, admin: false},
				"leland":   {name: "leland", number: 8015558941, admin: false},
			},
			name: "laura",
			expectedUsers: map[string]user{
				"log lady": {name: "log lady", number: 4355556025, admin: false},
				"shelly":   {name: "shelly", number: 8015558940, admin: false},
				"leland":   {name: "leland", number: 8015558941, admin: false},
			},
			expectedLog: logNotFound,
		})
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		originalMap := getMapCopy(test.users)
		log := logAndDelete(test.users, test.name)
		if log != test.expectedLog || !compareMaps(test.users, test.expectedUsers) {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  Users: 
%s
  Name: %v
  Expected Users: 
%s
  Actual Users: 
%s
  Expected Log: %s
  Actual Log: %s
`, formatMap(originalMap), test.name, formatMap(test.expectedUsers), formatMap(test.users), test.expectedLog, log)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  Users:
%s
  Name: %v
  Expected Users:
%s
  Actual Users:
%s
  Expected Log: %s
  Actual Log: %s
`, formatMap(originalMap), test.name, formatMap(test.expectedUsers), formatMap(test.users), test.expectedLog, log)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func getMapCopy(m map[string]user) map[string]user {
	copy := make(map[string]user)
	for key, value := range m {
		copy[key] = value
	}
	return copy
}

func formatMap(m map[string]user) string {
	var str string
	for key, value := range m {
		str += fmt.Sprintf("  * %s: %v\n", key, value)
	}
	return str
}

func compareMaps(map1, map2 map[string]user) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key, value1 := range map1 {
		if value2, exist := map2[key]; !exist || value1 != value2 {
			return false
		}
	}
	return true
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
