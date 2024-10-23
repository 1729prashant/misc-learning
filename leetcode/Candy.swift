//135. Candy
//There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.
//
//You are giving candies to these children subjected to the following requirements:
//
//Each child must have at least one candy.
//Children with a higher rating get more candies than their neighbors.
//Return the minimum number of candies you need to have to distribute the candies to the children.
//
// 
//
//Example 1:
//
//Input: ratings = [1,0,2]
//Output: 5
//Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.
//Example 2:
//
//Input: ratings = [1,2,2]
//Output: 4
//Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
//The third child gets 1 candy because it satisfies the above two conditions.
// 
//
//Constraints:
//
//n == ratings.length
//1 <= n <= 2 * 104
//0 <= ratings[i] <= 2 * 104

class Solution {
    func candy(_ ratings: [Int]) -> Int {
        let n = ratings.count
        if n <= 1 { return n }
        
        // Initialize a candies array with 1 candy for each child
        var candies = Array(repeating: 1, count: n)
        
        // Left-to-right pass
        for i in 1..<n {
            if ratings[i] > ratings[i - 1] {
                candies[i] = candies[i - 1] + 1
            }
        }
        
        // Right-to-left pass
        for i in (0..<n-1).reversed() {
            if ratings[i] > ratings[i + 1] {
                candies[i] = max(candies[i], candies[i + 1] + 1)
            }
        }
        
        // Sum the candies to get the minimum number
        return candies.reduce(0, +)
    }
}



import XCTest

class Tests: XCTestCase {

    private let solution = Solution()
    func testA() {
        let value = solution.candy([1,0,2])
        XCTAssertEqual(value, 5)
    }
    
    func testB() {
        let value = solution.candy([1,2,87,87,87,2,1])
        XCTAssertEqual(value, 13)
    }
}

Tests.defaultTestSuite.run()
