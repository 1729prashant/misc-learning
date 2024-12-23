"""
30. Substring with Concatenation of All Words

You are given a string s and an array of strings words. All the strings of words are of the same length.

A concatenated string is a string that exactly contains all the strings of any permutation of words concatenated.

For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are all concatenated strings. "acdbef" is not a concatenated string because it is not the concatenation of any permutation of words.
Return an array of the starting indices of all the concatenated substrings in s. You can return the answer in any order.

 

Example 1:

Input: s = "barfoothefoobarman", words = ["foo","bar"]

Output: [0,9]

Explanation:

The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.

Example 2:

Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]

Output: []

Explanation:

There is no concatenated substring.

Example 3:

Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]

Output: [6,9,12]

Explanation:

The substring starting at 6 is "foobarthe". It is the concatenation of ["foo","bar","the"].
The substring starting at 9 is "barthefoo". It is the concatenation of ["bar","the","foo"].
The substring starting at 12 is "thefoobar". It is the concatenation of ["the","foo","bar"].

 

Constraints:

1 <= s.length <= 104
1 <= words.length <= 5000
1 <= words[i].length <= 30
s and words[i] consist of lowercase English letters.
"""


from collections import Counter, defaultdict

class Solution:
    def findSubstring(self, s: str, words: List[str]) -> List[int]:
        if not s or not words:
            return []
        
        # Initialize key variables
        word_len = len(words[0])
        word_count = len(words)
        total_len = word_len * word_count
        result = []
        
        # If the total length is greater than string length, return empty
        if total_len > len(s):
            return []
        
        # Create a frequency map of words
        word_freq = Counter(words)
        
        # Check each possible starting position
        for i in range(len(s) - total_len + 1):
            seen = defaultdict(int)
            valid = True
            
            # Check each word position
            for j in range(word_count):
                # Extract the current word
                start_pos = i + j * word_len
                current_word = s[start_pos:start_pos + word_len]
                
                # If this word isn't in our word list, this position is invalid
                if current_word not in word_freq:
                    valid = False
                    break
                
                # Count the occurrence of this word
                seen[current_word] += 1
                
                # If we've seen this word more times than it appears in words, position is invalid
                if seen[current_word] > word_freq[current_word]:
                    valid = False
                    break
            
            # If we've found a valid concatenation, add the starting position
            if valid:
                result.append(i)
        
        return result
