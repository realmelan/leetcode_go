/*

1463. Cherry Pickup II
Solved
Hard
Topics
Companies
Hint
You are given a rows x cols matrix grid representing a field of cherries where grid[i][j] represents the number of cherries that you can collect from the (i, j) cell.

You have two robots that can collect cherries for you:

Robot #1 is located at the top-left corner (0, 0), and
Robot #2 is located at the top-right corner (0, cols - 1).
Return the maximum number of cherries collection using both robots by following the rules below:

From a cell (i, j), robots can move to cell (i + 1, j - 1), (i + 1, j), or (i + 1, j + 1).
When any robot passes through a cell, It picks up all cherries, and the cell becomes an empty cell.
When both robots stay in the same cell, only one takes the cherries.
Both robots cannot move outside of the grid at any moment.
Both robots should reach the bottom row in grid.
 

Example 1:


Input: grid = [[3,1,1],[2,5,1],[1,5,5],[2,1,1]]
Output: 24
Explanation: Path of robot #1 and #2 are described in color green and blue respectively.
Cherries taken by Robot #1, (3 + 2 + 5 + 2) = 12.
Cherries taken by Robot #2, (1 + 5 + 5 + 1) = 12.
Total of cherries: 12 + 12 = 24.
Example 2:


Input: grid = [[1,0,0,0,0,0,1],[2,0,0,0,0,3,0],[2,0,9,0,0,0,0],[0,3,0,5,4,0,0],[1,0,2,3,0,0,6]]
Output: 28
Explanation: Path of robot #1 and #2 are described in color green and blue respectively.
Cherries taken by Robot #1, (1 + 9 + 5 + 2) = 17.
Cherries taken by Robot #2, (1 + 3 + 4 + 3) = 11.
Total of cherries: 17 + 11 = 28.
 

Constraints:

rows == grid.length
cols == grid[i].length
2 <= rows, cols <= 70
0 <= grid[i][j] <= 100
*/

func cherryPickup(grid [][]int) int {
    // let dp[r][x][y] = # of cherries can be picked when one at (r,x) 
    // and the other at (r,y)
    // then dp[r][x][y] = max(dp[r-1][x-1,x,x+1][y-1,y,y+1]) + grid[r][y] + grid[r][y](or 0 if x=y)
    r, c := len(grid), len(grid[0])
    dp := make([][][]int, r)
    for i := range dp {
        dp[i] = make([][]int, c)
        for j := range dp[i] {
            dp[i][j] = make([]int, c)
        }
    }

    res := 0
    dp[0][0][c-1] = grid[0][0] + grid[0][c-1]
    dir := []int{-1,0,1}
    for i := 1; i < r; i++ {
        for x := 0; x < c && x <= i; x++ {
            for y := c-1; y >= 0 && y >= c-i-1; y-- {
                for _, dx := range dir {
                    for _, dy := range dir {
                        nx, ny := x+dx, y+dy
                        if nx >= 0 && nx < c && nx < i && ny >= 0 && ny >= c-i && ny < c  {
                            dp[i][x][y] = max(dp[i][x][y], dp[i-1][nx][ny])
                        }
                    }
                }
                dp[i][x][y] += grid[i][x]
                if x != y {
                    dp[i][x][y] += grid[i][y]
                }
                //fmt.Println("dp[%v][%v][%v]=%v", i, x, y, dp[i][x][y])
                if i == r-1 {
                    res = max(res, dp[i][x][y])
                }
            }
        }
    }
    //fmt.Println("dp=%v", dp)
    return res
}
