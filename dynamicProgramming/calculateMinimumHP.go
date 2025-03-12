/*
174. Dungeon Game
Solved
Hard
Topics
Companies
The demons had captured the princess and imprisoned her in the bottom-right corner of a dungeon. The dungeon consists of m x n rooms laid out in a 2D grid. Our valiant knight was initially positioned in the top-left room and must fight his way through dungeon to rescue the princess.

The knight has an initial health point represented by a positive integer. If at any point his health point drops to 0 or below, he dies immediately.

Some of the rooms are guarded by demons (represented by negative integers), so the knight loses health upon entering these rooms; other rooms are either empty (represented as 0) or contain magic orbs that increase the knight's health (represented by positive integers).

To reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.

Return the knight's minimum initial health so that he can rescue the princess.

Note that any room can contain threats or power-ups, even the first room the knight enters and the bottom-right room where the princess is imprisoned.

 

Example 1:


Input: dungeon = [[-2,-3,3],[-5,-10,1],[10,30,-5]]
Output: 7
Explanation: The initial health of the knight must be at least 7 if he follows the optimal path: RIGHT-> RIGHT -> DOWN -> DOWN.
Example 2:

Input: dungeon = [[0]]
Output: 1
 

Constraints:

m == dungeon.length
n == dungeon[i].length
1 <= m, n <= 200
-1000 <= dungeon[i][j] <= 1000
*/

func calculateMinimumHP(dungeon [][]int) int {
    // let dp[x][y] = minimum hp needed to survive
    // starting from bottom right corner, expand till it covers the starting room
    m, n := len(dungeon), len(dungeon[0])
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    for i := 0; i <= m; i++ {
        for j := 0; j <= n; j++ {
            dp[i][j] = 50000000
        }
    }
    dp[m][n-1]=1
    dp[m-1][n]=1
    for x := m-1; x >= 0; x-- {
        for y := n-1; y >= 0; y-- {
            dx, dy := x+1, y
            rx, ry := x, y+1
            dp[x][y] = min(max(1, dp[dx][dy]-dungeon[x][y]), max(1, dp[rx][ry]-dungeon[x][y]))
        }
    }
    //fmt.Println("dp=%v", dp)
    return dp[0][0]
}
