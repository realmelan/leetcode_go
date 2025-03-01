/*
81. Search in Rotated Sorted Array II
Solved
Medium
Topics
Companies
There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).

Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].

Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.

You must decrease the overall operation steps as much as possible.

 

Example 1:

Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true
Example 2:

Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false
 

Constraints:

1 <= nums.length <= 5000
-104 <= nums[i] <= 104
nums is guaranteed to be rotated at some pivot.
-104 <= target <= 104
 

Follow up: This problem is similar to Search in Rotated Sorted Array, but nums may contain duplicates. Would this affect the runtime complexity? How and why?
*/
func search(nums []int, target int) bool {
    n := len(nums)
    lo, up := 0, n-1
    for lo <= up {
        m := (lo+up+1)/2
        if nums[m] == target {
            return true
        }
        // invariant: nums[up] <= nums[lo]
        if nums[lo] == nums[m] && nums[m] == nums[up] {
            lo++
            up--
        } else if nums[lo] <= nums[m] {
            if nums[lo] <= target && target < nums[m] {
                up = m-1
            } else {
                lo = m+1
            }
        } else {
            if nums[m] < target && target <= nums[up] {
                lo = m+1
            } else {
                up = m-1
            }
        }
    }
    return false
}
func search2(nums []int, target int) bool {
    n := len(nums)
    lo, up := 0, n-1
    for lo < up {
        m := (lo+up)/2
        if nums[m] == target {
            return true
        }
        if nums[lo] < nums[up] {
            break
        } else if nums[lo] > nums[m] {
            up = m
        } else if nums[lo] < nums[m] {
            lo = m+1
        } else {
            lo++
        }
    }
    //fmt.Println("lo=%v, val=%v", lo, nums[lo])
    if target < nums[lo] {
        return false
    }

    bottom := lo
    lo, up = bottom, n-1
    for lo <= up {
        m := (lo + up)/2
        if nums[m] == target {
            return true
        } else if nums[m] < target {
            lo=m+1
        } else {
            up=m-1
        }
    }
    lo, up = 0, bottom-1
    for lo <= up {
        m := (lo + up)/2
        if nums[m] == target {
            return true
        } else if nums[m] < target {
            lo=m+1
        } else {
            up=m-1
        }
    }
    return false
}
