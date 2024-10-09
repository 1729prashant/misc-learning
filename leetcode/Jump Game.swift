//55. Jump Game
//You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.
//
//Return true if you can reach the last index, or false otherwise.
//
//
//
//Example 1:
//
//Input: nums = [2,3,1,1,4]
//Output: true
//Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
//Example 2:
//
//Input: nums = [3,2,1,0,4]
//Output: false
//Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
//
//
//Constraints:
//
//1 <= nums.length <= 104
//0 <= nums[i] <= 105




class Solution {
    func canJump(_ nums: [Int]) -> Bool {
        
        var count = nums.count - 1
        var startValue = nums[0]
        if (startValue == 0 || startValue == 1) && nums.count == 1 {return true}
        if startValue == count {return true}

        
        if startValue >= 1 && nums.count > 1{
            for i in 1...startValue {
                var jumpBy = nums[i] % nums.count
                
                if i + jumpBy == count {
                    return true
                }
            }
        }
        
        return false
    }
}



var nums1 = [2,0,2]
var sol = Solution()
print(sol.canJump(nums1))



//Imagine you are at the start of a path:
//The array tells you how far you can jump from each position.
//Your goal is to move forward as far as possible with each jump.

//Track your maximum reach:
//As you move along the path, keep track of how far you can go. This is called maxReach.
//Initially, you're at index 0, so your reach starts from there.

//Step-by-step update:
//For each position you're on, check how far you can jump from that position (i.e., current index + nums[current index]).
//Update maxReach with the farthest position you can reach from the current position.

//Early exit if stuck:
//If at any point, you're standing on a position that you couldn't have reached (i.e., your current index is beyond maxReach), you can't proceed further, and it's impossible to reach the last index.
//Check if you can reach or pass the last index:

//As you move forward, if maxReach becomes greater than or equal to the last index of the array, you know you can reach the end, so you return true.
