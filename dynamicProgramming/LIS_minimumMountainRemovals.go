/*
1671. Minimum Number of Removals to Make Mountain Array
Solved
Hard
Topics
Companies
Hint
You may recall that an array arr is a mountain array if and only if:

arr.length >= 3
There exists some index i (0-indexed) with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
Given an integer array nums​​​, return the minimum number of elements to remove to make nums​​​ a mountain array.

 

Example 1:

Input: nums = [1,3,1]
Output: 0
Explanation: The array itself is a mountain array so we do not need to remove any elements.
Example 2:

Input: nums = [2,1,1,5,6,2,3,1]
Output: 3
Explanation: One solution is to remove the elements at indices 0, 1, and 5, making the array nums = [1,5,6,3,1].
 

Constraints:

3 <= nums.length <= 1000
1 <= nums[i] <= 109
It is guaranteed that you can make a mountain array out of nums.
*/

func minimumMountainRemovals(nums []int) int {
    // for each element i, find LIS(ending at nums[i]) on the left and right
    // then the longest mountain array with i as the peak is ma[i]=1+lis[0:i]+lis[i+1:n]
    n := len(nums)
    lis := make([]int, n)
    a := make([]int, 0)
    for i, num := range nums {
        lo, up := 0, len(a)
        for lo < up {
            m := (lo+up)/2
            if a[m] < num {
                lo = m+1
            } else {
                up = m
            }
        }
        lis[i] = lo+1
        if lo == len(a) {
            a = append(a, num)
        } else {
            a[lo] = num
        }
    }

    var res int
    var b []int
    for i := n-1; i >= 0; i-- {
        num := nums[i]
        lo, up := 0, len(b)
        for lo < up {
            m := (lo+up)/2
            if b[m] < num {
                lo = m+1
            } else {
                up = m
            }
        }
        if lis[i] > 1 && lo >= 1 {
            res = max(res, lis[i] + lo)
        }
        if lo == len(b) {
            b = append(b, num)
        } else {
            b[lo] = num
        }
    }
    return n-res
}
