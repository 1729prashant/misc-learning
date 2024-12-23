"""
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
"""

from collections import Counter

class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if not t or not s:
            return ""

        # Count how many of each character we need from t
        chars_needed = Counter(t)
        # Number of unique characters we need to find in s
        unique_chars_to_find = len(chars_needed)

        # Window boundaries
        window_start = 0
        window_end = 0
        # Count of characters in our current window
        chars_in_window = {}
        
        # Best solution so far: (length, start index, end index)
        best_window = float("inf"), None, None

        while window_end < len(s):
            # Add one character from s to our window
            character = s[window_end]
            chars_in_window[character] = chars_in_window.get(character, 0) + 1

            # If we've found all instances of this character we need
            if character in chars_needed and chars_in_window[character] == chars_needed[character]:
                unique_chars_to_find -= 1

            # Try to contract the window until it's no longer valid
            while unique_chars_to_find == 0:
                character = s[window_start]

                # Update if this window is the smallest so far
                if window_end - window_start + 1 < best_window[0]:
                    best_window = (window_end - window_start + 1, window_start, window_end)

                # Remove the character from the start of our window
                chars_in_window[character] -= 1
                if character in chars_needed and chars_in_window[character] < chars_needed[character]:
                    unique_chars_to_find += 1

                window_start += 1    

            # Expand the window by moving the right pointer
            window_end += 1    

        # If no valid window was found
        return "" if best_window[0] == float("inf") else s[best_window[1] : best_window[2] + 1]



# Test Cases
sol = Solution()

print(sol.minWindow("ADOBECODEBANC", "ABC"))  # Output: "BANC"
print(sol.minWindow("a", "a"))                # Output: "a"
print(sol.minWindow("a", "aa"))               # Output: ""
print(sol.minWindow("cabwefgewcwaefgcf", "cae"))  # Output: "cwae"
print(sol.minWindow("ab", "b"))               # Output: "b"
