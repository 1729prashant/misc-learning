"""
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
"""
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        char_index = {}
        max_length = 0
        start = 0

        for i, char in enumerate(s):
            if char in char_index and char_index[char] >= start:
                # If char is already in the current window, move start to right after the last occurrence
                start = char_index[char] + 1
            else:
                # If char is not in the window or outside of it, update max_length
                max_length = max(max_length, i - start + 1)
            
            char_index[char] = i  # Update the last seen index of char

        return max_length

import unittest

class TestLengthOfLongestSubstring(unittest.TestCase):

    def test_lengthOfLongestSubstring(self):
        solution = Solution()
        test_cases = [
            ("abcabcbb", 3),  # "abc" is the longest substring without repeating characters
            ("bbbbb", 1),     # The longest substring without repeating characters is "b"
            ("pwwkew", 3),    # The longest substring is "wke"
            ("", 0),          # Empty string
            ("dvdf", 3),      # "vdf" is the longest substring without repeating characters
            ("tmmzuxt", 5),   # "mzuxt" is the longest substring without repeating characters
        ]

        for s, expected in test_cases:
            with self.subTest(s=s):
                result = solution.lengthOfLongestSubstring(s)
                self.assertEqual(result, expected, f"Failed for '{s}': expected {expected}, got {result}")

if __name__ == '__main__':
    unittest.main()