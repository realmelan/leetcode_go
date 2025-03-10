/*
2035. Partition Array Into Two Arrays to Minimize Sum Difference
Attempted
Hard
Topics
Companies
Hint
You are given an integer array nums of 2 * n integers. You need to partition nums into two arrays of length n to minimize the absolute difference of the sums of the arrays. To partition nums, put each element of nums into one of the two arrays.

Return the minimum possible absolute difference.

 

Example 1:

example-1
Input: nums = [3,9,7,3]
Output: 2
Explanation: One optimal partition is: [3,9] and [7,3].
The absolute difference between the sums of the arrays is abs((3 + 9) - (7 + 3)) = 2.
Example 2:

Input: nums = [-36,36]
Output: 72
Explanation: One optimal partition is: [-36] and [36].
The absolute difference between the sums of the arrays is abs((-36) - (36)) = 72.
Example 3:

example-3
Input: nums = [2,-1,0,4,-2,-9]
Output: 0
Explanation: One optimal partition is: [2,4,-9] and [-1,0,-2].
The absolute difference between the sums of the arrays is abs((2 + 4 + -9) - (-1 + 0 + -2)) = 0.
 

Constraints:

1 <= n <= 15
nums.length == 2 * n
-107 <= nums[i] <= 107
*/

func minimumDifference(nums []int) int {
    total, n := 0, len(nums)
    for _, num := range nums {
        total += num
    }

    target, res := total/2, math.MaxInt
    if total < 0 {
        target = (total-1)/2
    }
    left := subsetSums(nums[0:n/2])
    right := subsetSums(nums[n/2:n])
    for i, sums := range left {
        other := right[n/2-i]

        j, jn := 0, len(other)
        k := len(sums)-1
        for k >= 0 && j < jn {
            tsum := other[j] + sums[k]
            if tsum == target {
                if total & 1 == 0 {
                    return 0
                }
                res = min(res, total - tsum - tsum)
                j++
            } else if tsum > target {
                res = min(res, tsum + tsum - total)
                k--
            } else {
                res = min(res, total - tsum - tsum)
                j++
            }
        }
    }
    return res
}

func subsetSums(nums []int) [][]int {
    n := len(nums)
    res := make([][]int, n+1)
    for i := 0; i < 1<<n; i++ {
        cnt, sum := 0, 0
        for j := 0; j < n; j++ {
            if (i>>j) & 1 > 0 {
                sum += nums[j]
                cnt++
            }
        }
        res[cnt] = append(res[cnt], sum)
    }
    for i := range res {
        sort.Ints(res[i])
    }
   
