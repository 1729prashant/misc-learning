/*
209. Minimum Size Subarray Sum

Given an array of positive integers nums and a positive integer target, return the minimal length of a 
subarray
 whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

 

Example 1:

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.
Example 2:

Input: target = 4, nums = [1,4,4]
Output: 1
Example 3:

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
 

Constraints:

1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 104
 
*/

package main

import (
    "fmt"
)

func minSubArrayLen(target int, nums []int) int {
    n := len(nums)
    left := 0
    currentSum := 0
    minLength := n + 1 // Use n+1 to represent "infinity" (impossible length)
    
    for right := 0; right < n; right++ {
        currentSum += nums[right] // Expand the window to the right

        // Shrink the window from the left as long as the condition is met
        for currentSum >= target {
            if minLength > right-left+1 {
                minLength = right - left + 1
            }
            currentSum -= nums[left]
            left++
        }
    }

    // If no valid window was found, return 0
    if minLength == n+1 {
        return 0
    }
    return minLength
}

func main() {
    nums := []int{2, 3, 1, 2, 4, 3}
    target := 7
    fmt.Println(minSubArrayLen(target, nums)) // Output: 2
}
