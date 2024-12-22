"""
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
"""


from typing import List

class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        n = len(nums)
        left = 0
        current_sum = 0
        min_length = float('inf')
        
        for right in range(n):
            current_sum += nums[right]  # Expand the window to the right
            
            # Shrink the window from the left as long as the condition is met
            while current_sum >= target:
                min_length = min(min_length, right - left + 1)
                current_sum -= nums[left]
                left += 1
        
        return min_length if min_length != float('inf') else 0

# Test cases
solution = Solution()
test_cases = [
    (7, [2, 3, 1, 2, 4, 3], 2),
    (15, [1, 2, 3, 4, 5], 0),
    (5, [5, 1, 2, 3], 1),
    (15, [1, 2, 3, 4, 5], 5),
    (11, [1, 2, 3, 4, 5], 2),
    (7, [], 0),
    (10, [1], 0),
    (1, [1, 2, 3, 4, 5], 1),
    (8, [4, 4, 4, 4], 2)
]

for i, (target, nums, expected) in enumerate(test_cases):
    result = solution.minSubArrayLen(target, nums)
    print(f"Test Case {i+1}: {'Pass' if result == expected else 'Fail'}, Output: {result}, Expected: {expected}")
