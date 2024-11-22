package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {
	wordMap := make(map[string]bool)

	for _, message := range messages {
		words := strings.Fields(message)
		for _, word := range words {
			lowercaseWord := strings.ToLower(word)
			wordMap[lowercaseWord] = true
		}
	}

	return len(wordMap)
}



/*
provided solution
package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {
	distinctWords := make(map[string]struct{})
	for _, message := range messages {
		for _, word := range strings.Fields(strings.ToLower(message)) {
			distinctWords[word] = struct{}{}
		}
	}
	return len(distinctWords)
}


*/
