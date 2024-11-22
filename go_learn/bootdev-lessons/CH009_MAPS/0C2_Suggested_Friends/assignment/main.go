package main

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	// Map to store direct friends of the user
	directFriends := make(map[string]bool)
	for _, friend := range friendships[username] {
		directFriends[friend] = true
	}

	// Map to store suggested friends, ensuring each suggested friend is unique
	suggestedFriendsMap := make(map[string]bool)

	// Loop through each direct friend of the user
	for _, friend := range friendships[username] {
		// Find the friends of the user's direct friends
		for _, friendsFriend := range friendships[friend] {
			// Add to suggested friends if they are not the user and not a direct friend
			if friendsFriend != username && !directFriends[friendsFriend] {
				suggestedFriendsMap[friendsFriend] = true
			}
		}
	}

	// Convert the map of suggested friends to a slice
	suggestedFriends := []string{}
	for suggestedFriend := range suggestedFriendsMap {
		suggestedFriends = append(suggestedFriends, suggestedFriend)
	}

	// Return nil if there are no suggested friends
	if len(suggestedFriends) == 0 {
		return nil
	}
	return suggestedFriends
}


/*
provided solution

package main

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	directFriends := make(map[string]bool)
	for _, friend := range friendships[username] {
		directFriends[friend] = true
	}

	mutualFriends := make(map[string]bool)
	for _, friend := range friendships[username] {
		for _, friendOfFriend := range friendships[friend] {
			if friendOfFriend != username && !directFriends[friendOfFriend] {
				mutualFriends[friendOfFriend] = true
			}
		}
	}

	var suggestions []string
	for friend := range mutualFriends {
		suggestions = append(suggestions, friend)
	}

	return suggestions
}


*/
