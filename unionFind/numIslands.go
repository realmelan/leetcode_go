/*

200. Number of Islands
Solved
Medium
Topics
Companies
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

 

Example 1:

Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
Example 2:

Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
 

Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.

*/

func numIslands(grid [][]byte) int {
    m, n := len(grid), len(grid[0])
    p := make([]int, m*n)
    for id := 0; id < m*n; id++ {
        p[id] = id
    }
    id := 0
    dir := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
    for r, row := range grid {
        for c, b :=range row {
            if b != '1' {
                continue
            }
            id := r*n+c
            for _, d := range dir {
                x, y := r+d[0],c+d[1]
                if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != '1' {
                    continue
                }
                nid := x*n+y
                p1 := find(p, id)
                p2 := find(p, nid)
                if p1>p2 {
                    p1, p2 = p2, p1
                }
                p[p1]=p2
            }
        }
    }
    id = 0
    res := 0
    for _, row := range grid {
        for _, b :=range row {
            if b != '1' {
                id++
                continue
            }
            p1 := find(p, id)
            if p1 == id {
                res++
            }
            id++
        }
    }
    return res
}

func find(p []int, x int) int {
    if p[x] == x {
        return x
    }
    p[x] = find(p, p[x])
    return p[x]
}
