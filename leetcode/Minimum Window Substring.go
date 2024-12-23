/*
Given two strings s and t of lengths m and n respectively, return the minimum window 
substring
 of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.

 

Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.
 

Constraints:

m == s.length
n == t.length
1 <= m, n <= 105
s and t consist of uppercase and lowercase English letters.
 

Follow up: Could you find an algorithm that runs in O(m + n) time?

*/

package main

import (
    "fmt"
)

func minWindow(s string, t string) string {
    if len(t) == 0 || len(s) == 0 {
        return ""
    }

    charsNeeded := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        charsNeeded[t[i]]++
    }

    uniqueCharsToFind := len(charsNeeded)
    var left, right, formed int
    charsInWindow := make(map[byte]int)
    bestWindow := [3]int{1<<31 - 1, 0, 0}

    for right < len(s) {
        character := s[right]
        charsInWindow[character]++

        if count, exists := charsNeeded[character]; exists && charsInWindow[character] == count {
            formed++
        }

        for formed == uniqueCharsToFind {
            if right-left+1 < bestWindow[0] {
                bestWindow = [3]int{right - left + 1, left, right}
            }

            character = s[left]
            charsInWindow[character]--
            if count, exists := charsNeeded[character]; exists && charsInWindow[character] < count {
                formed--
            }
            left++
        }

        right++
    }

    if bestWindow[0] == 1<<31-1 {
        return ""
    }
    return s[bestWindow[1]:bestWindow[2]+1]
}

func main() {
    fmt.Println(minWindow("ADOBECODEBANC", "ABC"))  // Output: "BANC"
    fmt.Println(minWindow("a", "a"))                 // Output: "a"
    fmt.Println(minWindow("a", "aa"))                // Output: ""
    fmt.Println(minWindow("cabwefgewcwaefgcf", "cae")) // Output: "cwae"
    fmt.Println(minWindow("ab", "b"))                // Output: "b"
}
