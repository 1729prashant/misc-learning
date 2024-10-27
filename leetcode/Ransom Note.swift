//383. Ransom Note
//Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
//
//Each letter in magazine can only be used once in ransomNote.
//
//
//
//Example 1:
//
//Input: ransomNote = "a", magazine = "b"
//Output: false
//Example 2:
//
//Input: ransomNote = "aa", magazine = "ab"
//Output: false
//Example 3:
//
//Input: ransomNote = "aa", magazine = "aab"
//Output: true
//
//
//Constraints:
//
//1 <= ransomNote.length, magazine.length <= 105
//ransomNote and magazine consist of lowercase English letters.


class Solution {
    func canConstruct(_ ransomNote: String, _ magazine: String) -> Bool {
        var ransomDict: [Character: Int] = [:]
        var magazineDict: [Character: Int] = [:]
        
        for char in ransomNote {
            ransomDict[char, default: 0] += 1
        }
        
        for char in magazine {
            magazineDict[char, default: 0] += 1
        }
        
        for (char, count) in ransomDict {
            // Ensure `magazineDict` has enough of each character needed in `ransomDict`
            if magazineDict[char, default: 0] < count {
                return false
            }
        }
        return true
    }
}

