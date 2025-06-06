/*
719. Find K-th Smallest Pair Distance
Solved
Hard
Topics
Companies
Hint
The distance of a pair of integers a and b is defined as the absolute difference between a and b.

Given an integer array nums and an integer k, return the kth smallest distance among all the pairs nums[i] and nums[j] where 0 <= i < j < nums.length.

 

Example 1:

Input: nums = [1,3,1], k = 1
Output: 0
Explanation: Here are all the pairs:
(1,3) -> 2
(1,1) -> 0
(3,1) -> 2
Then the 1st smallest distance pair is (1,1), and its distance is 0.
Example 2:

Input: nums = [1,1,1], k = 2
Output: 0
Example 3:

Input: nums = [1,6,1], k = 3
Output: 5
 

Constraints:

n == nums.length
2 <= n <= 104
0 <= nums[i] <= 106
1 <= k <= n * (n - 1) / 2
*/
func smallestDistancePair(nums []int, k int) int {
    sort.Ints(nums)
    n := len(nums)
    lo, up := 0, nums[n-1]-nums[0]
    for lo < up {
        mid := (lo+up)/2
        le := check(nums, mid)
        if le < k {
            lo = mid+1
        } else {
            up = mid
        }
    }
    return up
}

func check(nums []int, di int) (le int) {
    n := len(nums)
    j := 0
    for i := 0; i < n-1; i++ {
        for j < n && nums[j] - nums[i] <= di {
            j++
        }
        le += j-i-1
    }
    return
}

func smallestDistancePair(nums []int, k int) int {
    sort.Ints(nums)
    n := len(nums)
    lo, up := 0, nums[n-1]-nums[0]
    res := up
    for lo <= up {
        mid := (lo+up)/2
        le, lec := check(nums, mid)
        if le < k {
            lo = mid+1
        } else {
            up = mid-1
            res = min(res, lec)
        }
    }
    return res
}

func check(nums []int, di int) (le, lec int) {
    n := len(nums)
    for i := 0; i < n-1; i++ {
        t := nums[i]+di
        x, _ := slices.BinarySearch(nums, t+1)
        le += x-i-1
        lec = max(lec, nums[x-1]-nums[i])
    }
    return
}
