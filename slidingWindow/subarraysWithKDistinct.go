/*
992. Subarrays with K Different Integers
Solved
Hard
Topics
Companies
Hint
Given an integer array nums and an integer k, return the number of good subarrays of nums.

A good array is an array where the number of different integers in that array is exactly k.

For example, [1,2,3,1,2] has 3 different integers: 1, 2, and 3.
A subarray is a contiguous part of an array.

 

Example 1:

Input: nums = [1,2,1,2,3], k = 2
Output: 7
Explanation: Subarrays formed with exactly 2 different integers: [1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2]
Example 2:

Input: nums = [1,2,1,3,4], k = 3
Output: 3
Explanation: Subarrays formed with exactly 3 different integers: [1,2,1,3], [2,1,3], [1,3,4].
 

Constraints:

1 <= nums.length <= 2 * 104
1 <= nums[i], k <= nums.length
*/
func subarraysWithKDistinct(nums []int, k int) int {
    i, j, p, n, res := 0, 0, 0, len(nums), 0
    mi := make(map[int]int) // count for i
    mj := make(map[int]int)
    for p < n {
        mi[nums[p]]++
        mj[nums[p]]++
        p++
        for len(mi) > k {
            mi[nums[i]]--
            if mi[nums[i]] == 0 {
                delete(mi, nums[i])
            }
            i++
        }
        for len(mj) >= k {
            mj[nums[j]]--
            if mj[nums[j]] == 0 {
                delete(mj, nums[j])
            }
            j++
        }
        res = res + j-i
    }
    return res
}
