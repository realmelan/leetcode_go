/*
300. Longest Increasing Subsequence
Solved
Medium
Topics
Companies
Given an integer array nums, return the length of the longest strictly increasing subsequence.

 

Example 1:

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Example 2:

Input: nums = [0,1,0,3,2,3]
Output: 4
Example 3:

Input: nums = [7,7,7,7,7,7,7]
Output: 1
 

Constraints:

1 <= nums.length <= 2500
-104 <= nums[i] <= 104
 

Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?
*/

func lengthOfLIS(nums []int) int {
    // dp[i] = length of LIS
    // then dp[i] = max(dp[j]+1 if nums[j]<nums[i])
    // O(n^2)
    // to improve max(dp[j]), maintain an array of LIS from 1 to k, such LIS[j] = smallest num that
    // has LIS of length j
    lis := []int{-100000} // lis is sorted by nature
    for _, num := range nums {
        lo, up := 0, len(lis)-1
        // find largest index such that lis[j] < num
        for lo < up { 
            m := (lo+up+1)/2
            if lis[m] >= num {
                up = m-1
            } else {
                lo = m
            }
        }
        if lo == len(lis)-1 {
            lis = append(lis, num)
        } else {
            lis[lo+1] = num
        }
    }
    return len(lis)-1
}
