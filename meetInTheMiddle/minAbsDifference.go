/*

1755. Closest Subsequence Sum
Solved
Hard
Topics
Companies
Hint
You are given an integer array nums and an integer goal.

You want to choose a subsequence of nums such that the sum of its elements is the closest possible to goal. That is, if the sum of the subsequence's elements is sum, then you want to minimize the absolute difference abs(sum - goal).

Return the minimum possible value of abs(sum - goal).

Note that a subsequence of an array is an array formed by removing some elements (possibly all or none) of the original array.

 

Example 1:

Input: nums = [5,-7,3,5], goal = 6
Output: 0
Explanation: Choose the whole array as a subsequence, with a sum of 6.
This is equal to the goal, so the absolute difference is 0.
Example 2:

Input: nums = [7,-9,15,-2], goal = -5
Output: 1
Explanation: Choose the subsequence [7,-9,-2], with a sum of -4.
The absolute difference is abs(-4 - (-5)) = abs(1) = 1, which is the minimum.
Example 3:

Input: nums = [1,2,3], goal = -7
Output: 7
 

Constraints:

1 <= nums.length <= 40
-107 <= nums[i] <= 107
-109 <= goal <= 109

*/

func minAbsDifference(nums []int, goal int) int {
    n := len(nums)
    m := n/2
    left := subsetSums(nums[0:m])
    right := subsetSums(nums[m:n])
    sort.Ints(left)
    sort.Ints(right)

    res := math.MaxInt
    i, in := 0, len(left)
    j := len(right)-1
    for i < in && j >= 0 {
        sum := left[i] + right[j]
        if sum == goal {
            return 0
        } else if sum > goal {
            res = min(res, sum-goal)
            j--
        } else {
            res = min(res, goal-sum)
            i++
        }
    }
    return res
}

func subsetSums(nums []int) []int {
    var res []int
    n := len(nums)
    for i:=0; i < 1<<n; i++ {
        sum := 0
        for j:=0; j < n; j++ {
            if (i>>j)&1 > 0 {
                sum += nums[j]
            }
        }
        res = append(res, sum)
    }
    //fmt.Println("nums=%v, res=%v", nums, res)
    return res
}
