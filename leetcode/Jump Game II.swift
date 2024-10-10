//45. Jump Game II
//You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].
//
//Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:
//
//0 <= j <= nums[i] and
//i + j < n
//Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].
//
// 
//
//Example 1:
//
//Input: nums = [2,3,1,1,4]
//Output: 2
//Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
//Example 2:
//
//Input: nums = [2,3,0,1,4]
//Output: 2
// 
//
//Constraints:
//
//1 <= nums.length <= 104
//0 <= nums[i] <= 1000
//It's guaranteed that you can reach nums[n - 1].




class Solution {
    func jump(_ nums: [Int]) -> Int {
        let n = nums.count
        if n == 1 { return 0 }
        
        var maxReach = 0      // The farthest index that can be reached
        var currentEnd = 0    // The farthest index we can jump to within the current jump
        var minJumps = 0      // The number of jumps made

        for i in 0..<n-1 {    // We don't need to check the last element
            maxReach = max(maxReach, i + nums[i])

            // If we've reached the end of the current jump range
            if i == currentEnd {
                minJumps += 1
                currentEnd = maxReach  // Set the new end for the next jump

                // If we can already reach the last index, break out of the loop
                if currentEnd >= n - 1 {
                    break
                }
            }
        }
        
        return minJumps
    }
}


var jumpArray = [2,3,1,1,4]
var solution = Solution()
print(solution.jump(jumpArray))




//Imagine standing on each index of the array:
//Every number in the array tells you how far you can jump forward from that index.

//Goal:
//You need to figure out the fewest number of jumps it takes to get to the last index in the array.

//How the approach works:
//Start at the first index. Your first "jump" takes you to some range of positions that you can reach (based on the value at nums[0]).
//As you move through the array, always keep track of the farthest index you can jump to from any position in your current jump.
//When you reach the farthest point in your current jump (i.e., the edge of the range), you must make another jump. Increase the jump count and update the new farthest point you can reach from that jump.

//When do we stop:
//If at any point, the farthest index you can reach is beyond or exactly at the last index of the array, you're done. The number of jumps you've counted is your answer.

//Example:
//For nums = [2, 3, 1, 1, 4]:
//Start at index 0 with nums[0] = 2. You can jump to index 1 or 2. The farthest you can go is index 2.
//Jump to index 1. From nums[1] = 3, you can jump to index 2, 3, or 4. The farthest you can go is index 4 (the end).
//Since you're already at or beyond the last index, you're done. It took 2 jumps to reach the end.
//Key Ideas:
//Each jump takes you to the farthest reachable point in the current jump range.
//You count jumps only when you reach the end of the current range, and then extend your reach for the next jump.
