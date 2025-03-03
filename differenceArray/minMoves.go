/*
1674. Minimum Moves to Make Array Complementary
Solved
Medium
Topics
Companies
Hint
You are given an integer array nums of even length n and an integer limit. In one move, you can replace any integer from nums with another integer between 1 and limit, inclusive.

The array nums is complementary if for all indices i (0-indexed), nums[i] + nums[n - 1 - i] equals the same number. For example, the array [1,2,3,4] is complementary because for all indices i, nums[i] + nums[n - 1 - i] = 5.

Return the minimum number of moves required to make nums complementary.

 

Example 1:

Input: nums = [1,2,4,3], limit = 4
Output: 1
Explanation: In 1 move, you can change nums to [1,2,2,3] (underlined elements are changed).
nums[0] + nums[3] = 1 + 3 = 4.
nums[1] + nums[2] = 2 + 2 = 4.
nums[2] + nums[1] = 2 + 2 = 4.
nums[3] + nums[0] = 3 + 1 = 4.
Therefore, nums[i] + nums[n-1-i] = 4 for every i, so nums is complementary.
Example 2:

Input: nums = [1,2,2,1], limit = 2
Output: 2
Explanation: In 2 moves, you can change nums to [2,2,2,2]. You cannot change any number to 3 since 3 > limit.
Example 3:

Input: nums = [1,2,1,2], limit = 2
Output: 0
Explanation: nums is already complementary.
 

Constraints:

n == nums.length
2 <= n <= 105
1 <= nums[i] <= limit <= 105
n is even.
*/
func minMoves(nums []int, limit int) int {
    n := len(nums)
    maxStop, minStart := 0, math.MaxInt
    m := make(map[int]int)
    for i := 0; i < n/2; i++ {
        sum := nums[i] + nums[n-1-i]
        m[1+min(nums[i],nums[n-1-i])]--
        minStart = min(minStart, 1+min(nums[i],nums[n-1-i]))
        m[sum]--
        m[sum+1]++
        m[limit+max(nums[i],nums[n-1-i])+1]++
        maxStop = max(maxStop, limit+max(nums[i],nums[n-1-i])+1)
    }

    res := n
    mod := n
    for i := minStart; i <= maxStop; i++ {
        mod += m[i]
        res = min(res, mod)
    }
    return res
}
