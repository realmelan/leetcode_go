/*
47. Permutations II
Solved
Medium
Topics
Companies
Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.

 

Example 1:

Input: nums = [1,1,2]
Output:
[[1,1,2],
 [1,2,1],
 [2,1,1]]
Example 2:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 

Constraints:

1 <= nums.length <= 8
-10 <= nums[i] <= 10
*/
func permuteUnique(nums []int) [][]int {
    res := make([][]int, 0)
    backtrack(nums, 0, &res)
    return res
}

func backtrack(nums []int, start int, res *[][]int) {
    n := len(nums)
    if start >= n-1 {
        *res = append(*res, slices.Clone(nums))
        return
    }

    m := make(map[int]bool)
    for i := start; i < n; i++ {
        if m[nums[i]] {
            continue
        }
        m[nums[i]] = true
        nums[start], nums[i] = nums[i], nums[start]
        backtrack(nums, start+1, res)
        nums[start], nums[i] = nums[i], nums[start]
    }
}
