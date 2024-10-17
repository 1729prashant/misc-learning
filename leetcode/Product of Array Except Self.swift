//238. Product of Array Except Self
//
//Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
//
//The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
//You must write an algorithm that runs in O(n) time and without using the division operation.
//
// 
//
//Example 1:
//
//Input: nums = [1,2,3,4]
//Output: [24,12,8,6]
//Example 2:
//
//Input: nums = [-1,1,0,-3,3]
//Output: [0,0,9,0,0]
// 
//
//Constraints:
//
//2 <= nums.length <= 105
//-30 <= nums[i] <= 30
//The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// 
//
//Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)



class Solution {
    func productExceptSelf(_ nums: [Int]) -> [Int] {
        var answer: [Int] = []
        
        for i in 0..<nums.count {
            var temp = nums
            temp.remove(at: i)
            var product = temp.reduce(1) { x, y in  x * y }
            answer.append(product)
            
        }
        return answer
    }
}


var nums = [1,2,3,4]
var OP = Solution().productExceptSelf(nums)


//faster
class Solution2 {
    func productExceptSelf(_ nums: [Int]) -> [Int] {
        let n = nums.count
        var answer = Array(repeating: 1, count: n)
        
        // Calculate prefix products
        var prefix = 1
        for i in 0..<n {
            answer[i] = prefix
            prefix *= nums[i]
        }
        
        // Calculate suffix products and multiply with prefix products
        var suffix = 1
        for i in (0..<n).reversed() {
            answer[i] *= suffix
            suffix *= nums[i]
        }
        
        return answer
    }
}
