/*
327. Count of Range Sum
Solved
Hard
Topics
Companies
Given an integer array nums and two integers lower and upper, return the number of range sums that lie in [lower, upper] inclusive.

Range sum S(i, j) is defined as the sum of the elements in nums between indices i and j inclusive, where i <= j.

 

Example 1:

Input: nums = [-2,5,-1], lower = -2, upper = 2
Output: 3
Explanation: The three ranges are: [0,0], [2,2], and [0,2] and their respective sums are: -2, -1, 2.
Example 2:

Input: nums = [0], lower = 0, upper = 0
Output: 1
 

Constraints:

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
-105 <= lower <= upper <= 105
The answer is guaranteed to fit in a 32-bit integer.
*/
func countRangeSum(nums []int, lower int, upper int) int {
    n := len(nums)
    sums := make([]int64, n+1)
    for i, num := range nums {
        sums[i+1] = sums[i] + int64(num)
    }
    _, res := merge(sums, 0, n, int64(lower), int64(upper))
    return res
}

func merge(nums []int64, l, r int, lo, up int64) (arr []int64, cnt int) {
    if l >= r {
        return []int64{nums[l]}, 0
    }

    m := (l+r)/2
    left, lcnt := merge(nums, l, m, lo, up)
    right, rcnt := merge(nums, m+1, r, lo, up)
    cnt += lcnt + rcnt

    i, ni := 0, len(left)
    j, k, nj := 0, 0, len(right)
    for i < ni {
        for j < nj && right[j] - left[i] < lo {
            j++
        }
        for k < nj && right[k] - left[i] <= up {
            k++
        }
        cnt += k-j
        i++
    }

    //fmt.Println("l=%v,r=%v, left=%v, right=%v", l, r, left, right)
    i, j = 0, 0
    for i < ni {
        for j < nj && right[j] < left[i] {
            arr = append(arr, right[j])
            j++
        }
        arr = append(arr, left[i])
        i++
    }
    arr = append(arr, left[i:ni]...)
    arr = append(arr, right[j:nj]...)
    //fmt.Println("l=%v, r=%v, arr=%v, cnt=%v", l, r, arr, cnt)
    return
}
