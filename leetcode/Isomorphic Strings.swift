//205. Isomorphic Strings
//Given two strings s and t, determine if they are isomorphic.
//
//Two strings s and t are isomorphic if the characters in s can be replaced to get t.
//
//All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.
//
// 
//
//Example 1:
//
//Input: s = "egg", t = "add"
//
//Output: true
//
//Explanation:
//
//The strings s and t can be made identical by:
//
//Mapping 'e' to 'a'.
//Mapping 'g' to 'd'.
//Example 2:
//
//Input: s = "foo", t = "bar"
//
//Output: false
//
//Explanation:
//
//The strings s and t can not be made identical as 'o' needs to be mapped to both 'a' and 'r'.
//
//Example 3:
//
//Input: s = "paper", t = "title"
//
//Output: true
//
// 
//
//Constraints:
//
//1 <= s.length <= 5 * 104
//t.length == s.length
//s and t consist of any valid ascii character.



class Solution {
    func isIsomorphic(_ s: String, _ t: String) -> Bool {
        if s.count != t.count {
            return false
        }
        
        var sToTMap: [Character: Character] = [:]
        var tToSMap: [Character: Character] = [:]
        
        for (sChar, tChar) in zip(s, t) {
            if let mappedChar = sToTMap[sChar], mappedChar != tChar {
                return false
            }
            if let mappedChar = tToSMap[tChar], mappedChar != sChar {
                return false
            }
            
            sToTMap[sChar] = tChar
            tToSMap[tChar] = sChar
        }
        
        return true
    }
}



var s = "paper"
var t = "title"

let sol = Solution().isIsomorphic(s, t)
print("\(s) and \(t) are \(sol)")


class Solution2 {
    func isIsomorphic(_ s: String, _ t: String) -> Bool {
        var sdc = [Character:String.Index](), tdc = sdc
        for i in s.indices {
            guard sdc[s[i]] == tdc[t[i]] else { return false }
            sdc[s[i]] = i
            tdc[t[i]] = i
        }
        return true
    }
}
