package main

import "unicode/utf8"

func getNameCounts(names []string) map[rune]map[string]int {
    result := make(map[rune]map[string]int)

    for _, name := range names {
        firstChar, _ := utf8.DecodeRuneInString(name)
        if _, ok := result[firstChar]; !ok {
            result[firstChar] = make(map[string]int)
        }
        result[firstChar][name]++
    }

    return result
}




/*
provided solution
package main

func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int)
	for _, name := range names {
		runes := []rune(name)
		firstLetter := runes[0]

		if _, ok := nameCounts[firstLetter]; !ok {
			nameCounts[firstLetter] = make(map[string]int)
		}
		nameCounts[firstLetter][name]++
	}
	return nameCounts
}


*/
