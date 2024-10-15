//274. H-Index
//Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper, return the researcher's h-index.
//
//According to the definition of h-index on Wikipedia: The h-index is defined as the maximum value of h such that the given researcher has published at least h papers that have each been cited at least h times.
//
// 
//
//Example 1:
//
//Input: citations = [3,0,6,1,5]
//Output: 3
//Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each of them had received 3, 0, 6, 1, 5 citations respectively.
//Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, their h-index is 3.
//Example 2:
//
//Input: citations = [1,3,1]
//Output: 1
// 
//
//Constraints:
//
//n == citations.length
//1 <= n <= 5000
//0 <= citations[i] <= 1000

class Solution {
    func hIndex(_ citations: [Int]) -> Int {
        // Sort citations in descending order
        let sortedCitations = citations.sorted(by: >)
        
        var h = 0
        for (index, citation) in sortedCitations.enumerated() {
            if citation >= index + 1 {
                h = index + 1
            } else {
                break
            }
        }
        
        return h
    }
}
//Sorting:
//sorted(by: >) sorts the citations array in descending order.

//Iterating Through Sorted Citations:
//Using enumerated(), we get both the index and citation.
//Condition: If citation >= index + 1, it means there are at least index + 1 papers with index + 1 citations each.
//Updating h: We set h to index + 1 whenever the condition is met.
//Breaking Early: If the condition fails, we break out of the loop as further citations will be smaller.

//Returning the h-Index:
//After the loop, h holds the maximum value satisfying the h-index condition.


class Solution2 {
    func hIndex(_ citations: [Int]) -> Int {
        let n = citations.count
        var counts = Array(repeating: 0, count: n + 1)
        
        // Count the number of papers for each citation count
        for citation in citations {
            if citation >= n {
                counts[n] += 1
            } else {
                counts[citation] += 1
            }
        }
        
        var total = 0
        // Iterate from highest possible h to 0
        for i in stride(from: n, through: 0, by: -1) {
            total += counts[i]
            if total >= i {
                return i
            }
        }
        
        return 0
    }
}


var citations = [3,0,6,1,5]
var hSolution = Solution()
hSolution.hIndex(citations)




//Sorting:
//sorted(by: >) sorts the citations array in descending order.

//Iterating Through Sorted Citations:
//Using enumerated(), we get both the index and citation.
//Condition: If citation >= index + 1, it means there are at least index + 1 papers with index + 1 citations each.
//Updating h: We set h to index + 1 whenever the condition is met.
//Breaking Early: If the condition fails, we break out of the loop as further citations will be smaller.

//Returning the h-Index:
//After the loop, h holds the maximum value satisfying the h-index condition.
