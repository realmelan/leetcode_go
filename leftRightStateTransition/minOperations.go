/*
2033. Minimum Operations to Make a Uni-Value Grid
Solved
Medium
Topics
Companies
Hint
You are given a 2D integer grid of size m x n and an integer x. In one operation, you can add x to or subtract x from any element in the grid.

A uni-value grid is a grid where all the elements of it are equal.

Return the minimum number of operations to make the grid uni-value. If it is not possible, return -1.

 

Example 1:


Input: grid = [[2,4],[6,8]], x = 2
Output: 4
Explanation: We can make every element equal to 4 by doing the following: 
- Add x to 2 once.
- Subtract x from 6 once.
- Subtract x from 8 twice.
A total of 4 operations were used.
Example 2:


Input: grid = [[1,5],[2,3]], x = 1
Output: 5
Explanation: We can make every element equal to 3.
Example 3:


Input: grid = [[1,2],[3,4]], x = 2
Output: -1
Explanation: It is impossible to make every element equal.
 

Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 105
1 <= m * n <= 105
1 <= x, grid[i][j] <= 104
*/
func minOperations(grid [][]int, x int) int {
    m := make(map[int]int)
    for _, row := range grid {
        for _, v := range row {
            m[v]++
        }
    }
    type pair struct {
        val int
        cnt int
    }
    nums := make([]pair, 0)
    for k, v := range m {
        nums = append(nums, pair{k,v})
    }
    sort.Slice(nums, func(i,j int)bool{
        return nums[i].val < nums[j].val
    })
    n := len(nums)
    left, right := make([]int, n), make([]int, n)
    cnt := 0
    for i := 0; i < n; i++ {
        left[i] = cnt
        cnt += nums[i].cnt
    }
    cnt = 0
    for i := n-1; i >= 0; i-- {
        right[i] = cnt
        cnt += nums[i].cnt
    }

    // calculate num of operations to make grid to nums[i].val
    cur := 0
    for i := 1; i < n; i++ {
        if (nums[i].val - nums[0].val) % x != 0 {
            return -1
        }
        cur += nums[i].cnt * (nums[i].val - nums[0].val) / x
    }

    res := cur
    for i := 1; i < n; i++ {
        cur += (left[i] - right[i-1]) * (nums[i].val - nums[i-1].val) / x
        res = min(res, cur)
    }
    return res
}
