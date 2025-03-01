/*
34. Find First and Last Position of Element in Sorted Array
Solved
Medium
Topics
Companies
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

 

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
Example 3:

Input: nums = [], target = 0
Output: [-1,-1]
 

Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109
*/
func searchRange(nums []int, target int) []int {
    n := len(nums)
    lo, up := 0, n-1
    for lo < up {
        m := (lo + up) / 2
        if nums[m] < target {
            lo = m+1
        } else {
            up = m
        }
    }
    if n == 0 || nums[lo] != target {
        return []int{-1,-1}
    }
    res := []int{lo, -1}
    up = n-1
    for lo < up {
        m := (lo+up+1)/2
        if nums[m] > target {
            up = m-1
        } else {
            lo = m
        }
    }
    res[1] = lo
    return res
}
