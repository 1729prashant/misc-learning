package main

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for i := 0; i < len(messagedUsers); i++ {
		user := messagedUsers[i]
		if _, ok := validUsers[user]; ok { // Check if the user is valid
			validUsers[user] += 1
		}
	}
}


/*
provided solution

package main

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		if _, ok := validUsers[user]; ok {
			validUsers[user]++
		}
	}
}


*/



