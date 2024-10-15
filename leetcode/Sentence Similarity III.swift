//1813. Sentence Similarity III
//You are given two strings sentence1 and sentence2, each representing a sentence composed of words. A sentence is a list of words that are separated by a single space with no leading or trailing spaces. Each word consists of only uppercase and lowercase English characters.
//
//Two sentences s1 and s2 are considered similar if it is possible to insert an arbitrary sentence (possibly empty) inside one of these sentences such that the two sentences become equal. Note that the inserted sentence must be separated from existing words by spaces.
//
//For example,
//
//s1 = "Hello Jane" and s2 = "Hello my name is Jane" can be made equal by inserting "my name is" between "Hello" and "Jane" in s1.
//s1 = "Frog cool" and s2 = "Frogs are cool" are not similar, since although there is a sentence "s are" inserted into s1, it is not separated from "Frog" by a space.
//Given two sentences sentence1 and sentence2, return true if sentence1 and sentence2 are similar. Otherwise, return false.
//
//
//
//Example 1:
//
//Input: sentence1 = "My name is Haley", sentence2 = "My Haley"
//
//Output: true
//
//Explanation:
//
//sentence2 can be turned to sentence1 by inserting "name is" between "My" and "Haley".
//
//Example 2:
//
//Input: sentence1 = "of", sentence2 = "A lot of words"
//
//Output: false
//
//Explanation:
//
//No single sentence can be inserted inside one of the sentences to make it equal to the other.
//
//Example 3:
//
//Input: sentence1 = "Eating right now", sentence2 = "Eating"
//
//Output: true
//
//Explanation:
//
//sentence2 can be turned to sentence1 by inserting "right now" at the end of the sentence.
//
//
//
//Constraints:
//
//1 <= sentence1.length, sentence2.length <= 100
//sentence1 and sentence2 consist of lowercase and uppercase English letters and spaces.
//The words in sentence1 and sentence2 are separated by a single space.

import Foundation

class Solution {
    func areSentencesSimilar(_ sentence1: String, _ sentence2: String) -> Bool {
        let sentence1Array = sentence1.components(separatedBy: " ")
        let sentence2Array = sentence2.components(separatedBy: " ")
        var copyArray: [String] = []
        var sentencePermutation: [String] = []
        
        if sentence1 == sentence2 {return true}
        
        if sentence1Array.count < sentence2Array.count {
            copyArray = sentence2Array
            copyArray.removeAll(where: { sentence1Array.contains($0) })
            
            var possible1 = (sentence1Array + copyArray).joined(separator: " ")
            var possible2 = (copyArray + sentence1Array).joined(separator: " ")
            
            if possible1 == sentence2 ||  possible2 == sentence2{
                return true
            }
            
            
            
            for i in 1..<sentence1Array.count {
                sentencePermutation = Array(sentence1Array[0..<i]) + copyArray + Array(sentence1Array[i..<sentence1Array.count])
                print("debug: \(sentencePermutation)")
                if sentencePermutation.joined(separator: " ") == sentence2 {
                    return true
                }
            }
            
            
        } else if sentence1Array.count > sentence2Array.count  {
            copyArray = sentence1Array
            copyArray.removeAll(where: { sentence2Array.contains($0) })
            
            var possible1 = (sentence2Array + copyArray).joined(separator: " ")
            var possible2 = (copyArray + sentence2Array).joined(separator: " ")
            if possible1 == sentence1 ||  possible2 == sentence1{
                return true
            }
            
            
            for i in 1..<sentence2Array.count {
                sentencePermutation = Array(sentence2Array[0..<i]) + copyArray + Array(sentence2Array[i..<sentence2Array.count])
                if sentencePermutation.joined(separator: " ") == sentence1 {
                    return true
                }
            }
        }
        
        return false
        
    }
}

var sentence1 = "I love programming"
var sentence2 = "I love playing football"
var checkSolution = Solution().areSentencesSimilar(sentence1,sentence2)
print(checkSolution)




//import Foundation

class Solution2 {
    func areSentencesSimilar(_ sentence1: String, _ sentence2: String) -> Bool {
        let sentence1Array = sentence1.split(separator: " ")
        let sentence2Array = sentence2.split(separator: " ")
        
        // Ensure sentence1Array is the shorter sentence
        if sentence1Array.count > sentence2Array.count {
            return areSentencesSimilar(sentence2, sentence1)
        }
        
        let sentence1WordCount = sentence1Array.count
        let sentence2WordCount = sentence2Array.count
        
        var prefixIndex = 0
        // Find the common prefixIndex
        while prefixIndex < sentence1WordCount && sentence1Array[prefixIndex] == sentence2Array[prefixIndex] {
            prefixIndex += 1
        }
        
        // If entire sentence1Array is a prefixIndex of sentence2Array
        if prefixIndex == sentence1WordCount {
            return true
        }
        
        var suffixIndex = 0
        // Find the common suffixIndex
        while suffixIndex < (sentence1WordCount - prefixIndex) && sentence1Array[sentence1WordCount - 1 - suffixIndex] == sentence2Array[sentence2WordCount - 1 - suffixIndex] {
            suffixIndex += 1
        }
        
        // Check if the remaining middle part in sentence1Array is empty
        return prefixIndex + suffixIndex == sentence1WordCount
    }
}


var checkSolution2 = Solution2().areSentencesSimilar("jiRNY fW rN S bpL r T S nl vndZkF om oUm ilsf pvyNJObW F Uj QNJUek", "jiRNY fW rN S bpL r T Uj QNJUek")
print(checkSolution2)
