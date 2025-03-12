/*
576. Out of Boundary Paths
Solved
Medium
Topics
Companies
Hint
There is an m x n grid with a ball. The ball is initially at the position [startRow, startColumn]. You are allowed to move the ball to one of the four adjacent cells in the grid (possibly out of the grid crossing the grid boundary). You can apply at most maxMove moves to the ball.

Given the five integers m, n, maxMove, startRow, startColumn, return the number of paths to move the ball out of the grid boundary. Since the answer can be very large, return it modulo 109 + 7.

 

Example 1:


Input: m = 2, n = 2, maxMove = 2, startRow = 0, startColumn = 0
Output: 6
Example 2:


Input: m = 1, n = 3, maxMove = 3, startRow = 0, startColumn = 1
Output: 12
 

Constraints:

1 <= m, n <= 50
0 <= maxMove <= 50
0 <= startRow < m
0 <= startColumn < n
*/

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
    // let dp[k][x][y] = # of paths for starting at [x,y] for k moves
    // then dp[k+1][i,j] = dp[k][x,y for adjacent to i, j]
    if maxMove <= 0 {
        return 0
    }
    const mod = 1000000007
    dp := make([][][]int, maxMove+1)
    for k := range dp {
        dp[k] = make([][]int, m)
        for j := range dp[k] {
            dp[k][j] = make([]int, n)
        }
    }

    dirs := [][]int{
        {1,0},{-1,0},{0,1},{0,-1},
    }
    for x := 0; x < m; x++ {
        dp[1][x][0] += 1
        dp[1][x][n-1] += 1
    }
    for y := 0; y < n; y++ {
        dp[1][0][y] += 1
        dp[1][m-1][y] += 1
    }
    for k := 2; k <= maxMove; k++ {
        for x := 0; x < m; x++ {
            for y := 0; y < n; y++ {
                for _, d := range dirs {
                    nx, ny := x + d[0], y + d[1]
                    if nx < 0 || nx >= m || ny < 0 || ny >= n {
                        continue
                    }
                    dp[k][x][y] += dp[k-1][nx][ny]
                    dp[k][x][y] %= mod
                }
            }
        }
    }

    var res int
    for k := 1; k <= maxMove; k++ {
        res += dp[k][startRow][startColumn]
        res %= mod
    }
    return res
}
