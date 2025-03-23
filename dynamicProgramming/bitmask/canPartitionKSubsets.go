/*

698. Partition to K Equal Sum Subsets
Solved
Medium
Topics
Companies
Hint
Given an integer array nums and an integer k, return true if it is possible to divide this array into k non-empty subsets whose sums are all equal.

 

Example 1:

Input: nums = [4,3,2,3,5,2,1], k = 4
Output: true
Explanation: It is possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal sums.
Example 2:

Input: nums = [1,2,3,4], k = 3
Output: false
 

Constraints:

1 <= k <= nums.length <= 16
1 <= nums[i] <= 104
The frequency of each element is in the range [1, 4].

*/

func canPartitionKSubsets(nums []int, k int) bool {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    if sum%k != 0 {
        return false
    }
    t := sum / k
    for _, v := range nums {
        if v > t {
            return false
        }
    }

    // let dp[s] = true if if s can be split into multiple subsets, with each subset's sum
    // <= t
    // then if dp[1<<n-1] is true, then the array can be split into k subset of t
    n := len(nums)
    dp := make([]bool, 1<<n)
    dp[0] = true
    sums := make([]int, 1<<n)

    for s := 0; s < 1<<n; s++ {
        if !dp[s] {
            continue
        }

        for i, num := range nums {
            if (s>>i) & 1 > 0 {
                continue
            }
            if sums[s]%t + num > t {
                continue
            }
            dp[s|1<<i] = true
            sums[s|1<<i] = sums[s] + num
        }
    }
    return dp[(1<<n)-1]
}
