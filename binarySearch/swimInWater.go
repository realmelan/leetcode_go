/*
778. Swim in Rising Water
Solved
Hard
Topics
Companies
Hint
You are given an n x n integer matrix grid where each value grid[i][j] represents the elevation at that point (i, j).

The rain starts to fall. At time t, the depth of the water everywhere is t. You can swim from a square to another 4-directionally adjacent square if and only if the elevation of both squares individually are at most t. You can swim infinite distances in zero time. Of course, you must stay within the boundaries of the grid during your swim.

Return the least time until you can reach the bottom right square (n - 1, n - 1) if you start at the top left square (0, 0).

 

Example 1:


Input: grid = [[0,2],[1,3]]
Output: 3
Explanation:
At time 0, you are in grid location (0, 0).
You cannot go anywhere else because 4-directionally adjacent neighbors have a higher elevation than t = 0.
You cannot reach point (1, 1) until time 3.
When the depth of water is 3, we can swim anywhere inside the grid.
Example 2:


Input: grid = [[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]
Output: 16
Explanation: The final route is shown.
We need to wait until time 16 so that (0, 0) and (4, 4) are connected.
 

Constraints:

n == grid.length
n == grid[i].length
1 <= n <= 50
0 <= grid[i][j] < n2
Each value grid[i][j] is unique.
*/

func swimInWater(grid [][]int) int {
    lo, up := 0, 0
    for _, g := range grid {
        for _, e := range g {
            up = max(up, e)
        }
    }
    for lo < up {
        mid := (lo + up) / 2
        //if bfs(grid, mid) {
        visited := make(map[int]bool)
        if dfs(grid, visited, 0, 0, max(grid[0][0], mid), mid) {
            up = mid
        } else {
            lo = mid+1
        }
    }
    return up
}

func id(x, y, n int) int {
    return x*n+y
}
func coord(id, n int) (x, y int) {
    x = id / n
    y = id % n
    return
}

func dfs(grid [][]int, visited map[int]bool, x, y, e0, t int) bool {
    n := len(grid)
    if x == n-1 && y == n-1 {
        return true
    }

    dir := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
    for _, d:= range dir {
        nx, ny := x + d[0], y + d[1]
        if nx < 0 || nx >= n || ny < 0 || ny >= n {
            continue
        }
        nid := id(nx, ny, n)
        if visited[nid] {
            continue
        }
        visited[nid] = true
        ne := max(grid[nx][ny], t)
        if e0 == ne {
            if dfs(grid, visited, nx, ny, e0, t) {
                return true
            }
        }
    }
    return false
}

func bfs(grid [][]int, t int) bool {
    dir := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
    n := len(grid)
    i := 0
    reach := make(map[int]bool)
    queue := make([]int, 0)
    reach[id(0,0, n)] = true
    queue = append(queue, id(0,0,n))
    for len(queue) > i {
        j := len(queue)
        for ; i < j; i++ {
            x, y := coord(queue[i], n)
            e := max(grid[x][y], t)
            // check the adjacent cells
            for _, d := range dir {
                nx, ny := x + d[0], y + d[1]
                if nx < 0 || nx >= n || ny < 0 || ny >= n {
                    continue
                }
                nid := id(nx, ny, n)
                if reach[nid] {
                    continue
                }
                ne := max(grid[nx][ny], t)
                if e == ne {
                    if nx == n-1 && ny == n-1 {
                        return true
                    }
                    queue = append(queue, nid)
                    reach[nid] = true
                }
            }
        }
    }
    return false
}
