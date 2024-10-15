//189. Rotate Array
//Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
//
//
//
//Example 1:
//
//Input: nums = [1,2,3,4,5,6,7], k = 3
//Output: [5,6,7,1,2,3,4]
//Explanation:
//rotate 1 steps to the right: [7,1,2,3,4,5,6]
//rotate 2 steps to the right: [6,7,1,2,3,4,5]
//rotate 3 steps to the right: [5,6,7,1,2,3,4]
//Example 2:
//
//Input: nums = [-1,-100,3,99], k = 2
//Output: [3,99,-1,-100]
//Explanation:
//rotate 1 steps to the right: [99,-1,-100,3]
//rotate 2 steps to the right: [3,99,-1,-100]
//
//
//Constraints:
//
//1 <= nums.length <= 105
//-231 <= nums[i] <= 231 - 1
//0 <= k <= 105


//var nums = [1,2,3,4,5,6,7]
//var k = 3

var nums = [-1,-100,3,99]
var k = 2
class Solution {
    func rotate(_ nums: inout [Int], _ k: Int) {
        let n = nums.count
        
        // Handle edge cases
        if k == 0 || n == 0 || k % n == 0 {
            return
        }
        // Calculate the effective rotate value (in case k > n)
        let rotateVal = k % n
        
        // Step 1: Reverse the entire array
        reverse(&nums, 0, n - 1)
        
        // Step 2: Reverse the first `rotateVal` elements
        reverse(&nums, 0, rotateVal - 1)
        
        // Step 3: Reverse the remaining `n - rotateVal` elements
        reverse(&nums, rotateVal, n - 1)
    }
    
    // Helper function to reverse a portion of the array in place
    private func reverse(_ nums: inout [Int], _ start: Int, _ end: Int) {
        let count = (end - start + 1) / 2  // Number of swaps required to reverse the segment
        //This calculates how many swaps are required to reverse the subarray. Since we only need to swap the first half of the range with the second half, we divide the number of elements in the range by 2.

        
        for i in 0..<count {
            nums.swapAt(start + i, end - i)
        }
        //The loop iterates over the first half of the subarray, swapping elements at start + i (beginning) and end - i (end).
        //For example, in the first iteration, i = 0, so it swaps the element at start with the element at end.
        //In the next iteration, i = 1, so it swaps the element at start + 1 with the element at end - 1, and so on.
    }
}


var solutions = Solution()
solutions.rotate(&nums, k)
print(nums)
