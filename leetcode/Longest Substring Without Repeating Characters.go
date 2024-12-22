/*
3. Longest Substring Without Repeating Characters

Given a string s, find the length of the longest 
substring
 without repeating characters.

 

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
 

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

package main

import (
    "testing"
)

func lengthOfLongestSubstring(s string) int {
    charIndex := make(map[byte]int)
    maxLength := 0
    start := 0

    for i, char := range []byte(s) {
        if _, ok := charIndex[char]; ok && charIndex[char] >= start {
            // If char is already in the current window, move start to right after the last occurrence
            start = charIndex[char] + 1
        } else {
            // If char is not in the window or outside of it, update maxLength
            if i-start+1 > maxLength {
                maxLength = i - start + 1
            }
        }
        charIndex[char] = i // Update the last seen index of char
    }

    return maxLength
}

// Test cases
func TestLengthOfLongestSubstring(t *testing.T) {
    testCases := []struct {
        s        string
        expected int
    }{
        {"abcabcbb", 3}, // "abc" is the longest substring without repeating characters
        {"bbbbb", 1},    // The longest substring without repeating characters is "b"
        {"pwwkew", 3},   // The longest substring is "wke"
        {"", 0},         // Empty string
        {"dvdf", 3},     // "vdf" is the longest substring without repeating characters
        {"tmmzuxt", 5},  // "mzuxt" is the longest substring without repeating characters
    }

    for _, tc := range testCases {
        result := lengthOfLongestSubstring(tc.s)
        if result != tc.expected {
            t.Errorf("For string '%s', expected %d but got %d", tc.s, tc.expected, result)
        }
    }
}
