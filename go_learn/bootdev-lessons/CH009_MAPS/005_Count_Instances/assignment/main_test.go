package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		messagedUsers []string
		validUsers    map[string]int
		expected      map[string]int
	}
	tests := []testCase{
		{
			[]string{"cersei", "jaime", "cersei"},
			map[string]int{"cersei": 0, "jaime": 0},
			map[string]int{"cersei": 2, "jaime": 1},
		},
		{
			[]string{"cersei", "tyrion", "jaime", "tyrion", "tyrion"},
			map[string]int{"cersei": 0, "jaime": 0, "tyrion": 0},
			map[string]int{"cersei": 1, "jaime": 1, "tyrion": 3},
		},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{
				[]string{},
				map[string]int{"tyrion": 0},
				map[string]int{"tyrion": 0},
			},
			{
				[]string{"cersei", "jaime", "tyrion"},
				map[string]int{"tywin": 0},
				map[string]int{"tywin": 0},
			},
			{
				[]string{"cersei", "cersei", "cersei", "tyrion"},
				map[string]int{"cersei": 0},
				map[string]int{"cersei": 3},
			},
			{
				[]string{"cersei", "tywin", "jaime", "cersei", "tyrion", "cersei", "jaime"},
				map[string]int{"cersei": 0, "jaime": 0, "tyrion": 0},
				map[string]int{"cersei": 3, "jaime": 2, "tyrion": 1},
			},
			{
				[]string{"cersei", "cersei", "jaime", "jaime", "tywin", "cersei", "tywin", "tyrion"},
				map[string]int{"cersei": 0, "jaime": 0, "tyrion": 0},
				map[string]int{"cersei": 3, "jaime": 2, "tyrion": 1},
			},
		}...)
	}

	passCount := 0
	failCount := 0

	for i, test := range tests {
		getCounts(test.messagedUsers, test.validUsers)
		if !compareMaps(test.validUsers, test.expected) {
			failCount++
			t.Errorf(`
Test #%v Failed:
  Messaged Users: %v
  Expected: %v
  Actual: %v
`, i, test.messagedUsers, test.expected, test.validUsers)
		} else {
			passCount++
			fmt.Printf(`
Test #%v Passed:
  Messaged Users: %v
  Expected: %v
  Actual: %v
`, i, test.messagedUsers, test.expected, test.validUsers)
		}
		fmt.Println("---------------------------------")
	}

	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func compareMaps(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if v2, ok := m2[k]; !ok || v != v2 {
			return false
		}
	}
	return true
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
