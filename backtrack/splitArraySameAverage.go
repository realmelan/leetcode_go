/*

805. Split Array With Same Average
Solved
Hard
Topics
Companies
You are given an integer array nums.

You should move each element of nums into one of the two arrays A and B such that A and B are non-empty, and average(A) == average(B).

Return true if it is possible to achieve that and false otherwise.

Note that for an array arr, average(arr) is the sum of all the elements of arr over the length of arr.

 

Example 1:

Input: nums = [1,2,3,4,5,6,7,8]
Output: true
Explanation: We can split the array into [1,4,5,8] and [2,3,6,7], and both of them have an average of 4.5.
Example 2:

Input: nums = [3,1]
Output: false
 

Constraints:

1 <= nums.length <= 30
0 <= nums[i] <= 104

*/

func splitArraySameAverage(nums []int) bool {
    total, n := 0, len(nums)
    if n <= 1 {
        return false
    }
    for _, num := range nums {
        total += num
    }
    if total == 0 {
        return true
    }

    memo := make([][]map[int]bool, n/2+1)
    for i := range memo {
        memo[i] = make([]map[int]bool, n)
        for j := range memo[i] {
            memo[i][j] = make(map[int]bool)
        }
    }

    for k := 1; k <= n/2; k++ {
        if total*k%n != 0 {
            continue
        }
        if dfs(k, 0, total*k/n, nums, memo) {
            return true
        }
    }
    return false
}

func dfs(k, start, sum int, nums []int, memo [][]map[int]bool) bool {
    if k == 0 {
        return sum == 0
    }
    n := len(nums)
    if start >= n {
        return false
    }

    if v, ok := memo[k][start][sum]; ok {
        return v
    }

    for i := start; i < n; i++ {
       if sum < nums[i] {
            continue
        }
        if dfs(k-1, i+1, sum-nums[i], nums, memo) {
            memo[k][start][sum] = true
            return true
        }
    }
    memo[k][start][sum] = false
    return false
}
