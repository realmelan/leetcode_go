/*
1473. Paint House III
Solved
Hard
Topics
Companies
Hint
There is a row of m houses in a small city, each house must be painted with one of the n colors (labeled from 1 to n), some houses that have been painted last summer should not be painted again.

A neighborhood is a maximal group of continuous houses that are painted with the same color.

For example: houses = [1,2,2,3,3,2,1,1] contains 5 neighborhoods [{1}, {2,2}, {3,3}, {2}, {1,1}].
Given an array houses, an m x n matrix cost and an integer target where:

houses[i]: is the color of the house i, and 0 if the house is not painted yet.
cost[i][j]: is the cost of paint the house i with the color j + 1.
Return the minimum cost of painting all the remaining houses in such a way that there are exactly target neighborhoods. If it is not possible, return -1.

 

Example 1:

Input: houses = [0,0,0,0,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3
Output: 9
Explanation: Paint houses of this way [1,2,2,1,1]
This array contains target = 3 neighborhoods, [{1}, {2,2}, {1,1}].
Cost of paint all houses (1 + 1 + 1 + 1 + 5) = 9.
Example 2:

Input: houses = [0,2,1,2,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3
Output: 11
Explanation: Some houses are already painted, Paint the houses of this way [2,2,1,2,2]
This array contains target = 3 neighborhoods, [{2,2}, {1}, {2,2}]. 
Cost of paint the first and last house (10 + 1) = 11.
Example 3:

Input: houses = [3,1,2,3], cost = [[1,1,1],[1,1,1],[1,1,1],[1,1,1]], m = 4, n = 3, target = 3
Output: -1
Explanation: Houses are already painted with a total of 4 neighborhoods [{3},{1},{2},{3}] different of target = 3.
 

Constraints:

m == houses.length == cost.length
n == cost[i].length
1 <= m <= 100
1 <= n <= 20
1 <= target <= m
0 <= houses[i] <= n
1 <= cost[i][j] <= 104
*/

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
    // let dp[t][i][c] = cost of paint t neighborhoods for house[0:i] with house[i] painted in color c.
    // dp[t][i][c] = min(dp[t][i-1][c], dp[t-1][i-1][nc]) +cost(i,c)
    dp := make([][][]int, target+1)
    for t := range dp {
        dp[t] = make([][]int, m+1)
        for h := range dp[t] {
            dp[t][h] = make([]int, n+1) // 0 for no color
            for c := range dp[t][h] {
                dp[t][h][c] = math.MaxInt
            }
        }
    }
    for c := 1; c <= n; c++ {
        dp[0][0][c] = 0
        dp[1][0][c] = 0
    }
    for t := 1; t <= target; t++ {
        for h := t; h <= m; h++ {
            if houses[h-1] != 0 {
                c := houses[h-1]
                for pc := 1; pc <= n; pc++ { // color of previous house
                    if pc == c {
                        dp[t][h][c] = min(dp[t][h][c], dp[t][h-1][c])
                    } else {
                        dp[t][h][c] = min(dp[t][h][c], dp[t-1][h-1][pc])
                    }
                }
                continue
            }

            for c := 1; c <= n; c++ {
                dp[t][h][c] = math.MaxInt
                for pc := 1; pc <= n; pc++ { // color of previous house
                    if pc == c {
                        dp[t][h][c] = min(dp[t][h][c], dp[t][h-1][c])
                    } else {
                        dp[t][h][c] = min(dp[t][h][c], dp[t-1][h-1][pc])
                    }
                }
                if dp[t][h][c] != math.MaxInt {
                    if houses[h-1] != c {
                        dp[t][h][c] += cost[h-1][c-1]
                    }
                }
            }
        }
    }

    res := math.MaxInt
    for c := 1; c <= n; c++ {
        res = min(res, dp[target][m][c])
    }
    if res == math.MaxInt {
        return -1
    }
    return res
}
