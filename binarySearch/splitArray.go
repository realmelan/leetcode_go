/*
410. Split Array Largest Sum
Solved
Hard
Topics
Companies
Given an integer array nums and an integer k, split nums into k non-empty subarrays such that the largest sum of any subarray is minimized.

Return the minimized largest sum of the split.

A subarray is a contiguous part of the array.

 

Example 1:

Input: nums = [7,2,5,10,8], k = 2
Output: 18
Explanation: There are four ways to split nums into two subarrays.
The best way is to split it into [7,2,5] and [10,8], where the largest sum among the two subarrays is only 18.
Example 2:

Input: nums = [1,2,3,4,5], k = 2
Output: 9
Explanation: There are four ways to split nums into two subarrays.
The best way is to split it into [1,2,3] and [4,5], where the largest sum among the two subarrays is only 9.
 

Constraints:

1 <= nums.length <= 1000
0 <= nums[i] <= 106
1 <= k <= min(50, nums.length)
*/
func splitArray(nums []int, k int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    lo, up := 0, sum
    res := sum
    for lo <= up {
        mid := (lo + up) / 2
        if found(nums, mid, k) {
            res = min(res, mid)
            up = mid-1
        } else {
            lo = mid+1
        }
    }
    return res
}

func found(nums []int, t, k int) bool {
    sum := 0
    for _, num := range nums {
        if num > t {
            return false
        }
        if sum + num <= t {
            sum += num
        } else {
            sum = num
            k--
        }
    }
    return k > 0
}
