/*
803. Bricks Falling When Hit
Solved
Hard
Topics
Companies
You are given an m x n binary grid, where each 1 represents a brick and 0 represents an empty space. A brick is stable if:

It is directly connected to the top of the grid, or
At least one other brick in its four adjacent cells is stable.
You are also given an array hits, which is a sequence of erasures we want to apply. Each time we want to erase the brick at the location hits[i] = (rowi, coli). The brick on that location (if it exists) will disappear. Some other bricks may no longer be stable because of that erasure and will fall. Once a brick falls, it is immediately erased from the grid (i.e., it does not land on other stable bricks).

Return an array result, where each result[i] is the number of bricks that will fall after the ith erasure is applied.

Note that an erasure may refer to a location with no brick, and if it does, no bricks drop.

 

Example 1:

Input: grid = [[1,0,0,0],[1,1,1,0]], hits = [[1,0]]
Output: [2]
Explanation: Starting with the grid:
[[1,0,0,0],
 [1,1,1,0]]
We erase the underlined brick at (1,0), resulting in the grid:
[[1,0,0,0],
 [0,1,1,0]]
The two underlined bricks are no longer stable as they are no longer connected to the top nor adjacent to another stable brick, so they will fall. The resulting grid is:
[[1,0,0,0],
 [0,0,0,0]]
Hence the result is [2].
Example 2:

Input: grid = [[1,0,0,0],[1,1,0,0]], hits = [[1,1],[1,0]]
Output: [0,0]
Explanation: Starting with the grid:
[[1,0,0,0],
 [1,1,0,0]]
We erase the underlined brick at (1,1), resulting in the grid:
[[1,0,0,0],
 [1,0,0,0]]
All remaining bricks are still stable, so no bricks fall. The grid remains the same:
[[1,0,0,0],
 [1,0,0,0]]
Next, we erase the underlined brick at (1,0), resulting in the grid:
[[1,0,0,0],
 [0,0,0,0]]
Once again, all remaining bricks are still stable, so no bricks fall.
Hence the result is [0,0].
 

Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 200
grid[i][j] is 0 or 1.
1 <= hits.length <= 4 * 104
hits[i].length == 2
0 <= xi <= m - 1
0 <= yi <= n - 1
All (xi, yi) are unique.
*/

func hitBricks(grid [][]int, hits [][]int) []int {
	// if a brick is connected to the top row directly or indirectly, then it is stable
	// we can't start from the original grid and build parent relation and remove connections
	// but we first remove all hits and build parent relations.
	// all bricks with parent at top row is stable.
	// then starting from the last of hits, and check its neighbors:
	//  If it has a stable neighbor, then hit is stable, and it may make other neighbors stable
	//.   (change from unstable to stable), and these neighbors will fall if hit is removed.
	res := make([]int, len(hits))
	m := len(grid)
	n := len(grid[0])
	for i, h := range hits {
		if grid[h[0]][h[1]] == 0 {
			res[i] = -1
		} else {
			grid[h[0]][h[1]] = 0
		}
	}
    p := make([]int, m*n+1)
    size := make([]int, m*n+1)
	for i := range p {
        p[i] = i
        size[i] = 1
    }
	// connect
	dir := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	for x, row := range grid {
		for y, b := range row {
			if b != 1 {
				continue
			}
			id := x*n + y
			for _, d := range dir {
				nx, ny := x+d[0], y+d[1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
					uni(p, size, id, nx*n+ny)
				}
			}
            if x == 0 {
                uni(p, size, id, m*n)
            }
		}
	}
	//fmt.Println("p=%v, size=%v", p, size)
	// starting from hits
	stableCnt := size[m*n]
	for i := len(hits) - 1; i >= 0; i-- {
		if res[i] == -1 {
			res[i] = 0
			continue
		}
		x, y := hits[i][0], hits[i][1]
		grid[x][y] = 1
        size[x*n+y] = 1
		if x == 0 {
            uni(p, size, x*n+y, m*n)
		}
		for _, d := range dir {
			nx, ny := x+d[0], y+d[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] == 0 {
				continue
			}
			uni(p, size, x*n+y, nx*n+ny)
		}

        res[i] = max(0, size[m*n]-stableCnt-1)
        stableCnt = size[m*n]
	}
	return res
}
func uni(p, size []int, x, y int) {
	px := find(p, x)
	py := find(p, y)
	if px > py {
		px, py = py, px
	}
	p[px] = py
	if px != py {
		size[py] += size[px]
	}
}
func hitBricks2(grid [][]int, hits [][]int) []int {
	// if a brick is connected to the top row directly or indirectly, then it is stable
	// we can't start from the original grid and build parent relation and remove connections
	// but we first remove all hits and build parent relations.
	// all bricks with parent at top row is stable.
	// then starting from the last of hits, and check its neighbors:
	//  If it has a stable neighbor, then hit is stable, and it may make other neighbors stable
	//.   (change from unstable to stable), and these neighbors will fall if hit is removed.
	res := make([]int, len(hits))
	m := len(grid)
	n := len(grid[0])
	p := make([]int, m*n+1)
	p[m*n] = m * n
	for i, h := range hits {
		if grid[h[0]][h[1]] == 0 {
			res[i] = -1
		} else {
			grid[h[0]][h[1]] = 0
		}
	}
	for r, row := range grid {
		for c, b := range row {
			if r == 0 && b == 1 {
				p[r*n+c] = m * n // stable
			} else {
				p[r*n+c] = r*n + c
			}
		}
	}
	// connect
	dir := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	for x, row := range grid {
		for y, b := range row {
			if b != 1 {
				continue
			}
			id := x*n + y
			for _, d := range dir {
				nx, ny := x+d[0], y+d[1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
					uni2(p, id, nx*n+ny)
				}
			}
		}
	}
	for x, row := range grid {
		for y, b := range row {
			if b != 1 {
				continue
			}
			find(p, x*n+y)
		}
	}
	//fmt.Println("p=%v", p)
	// starting from hits
	for i := len(hits) - 1; i >= 0; i-- {
		if res[i] == -1 {
			res[i] = 0
			continue
		}
		x, y := hits[i][0], hits[i][1]
		grid[x][y] = 1
		stable := false
		if x == 0 {
			stable = true
		} else {
			for _, d := range dir {
				nx, ny := x+d[0], y+d[1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] == 0 {
					continue
				}
				if find(p, nx*n+ny) == m*n {
					stable = true
					break
				}
			}
		}
		if stable {
			//p[x*n+y]=m*n
			fmt.Println("x,y=%v,%v", x, y)
			res[i] = dfs(grid, x, y, p) - 1
		}
	}
	return res
}

func dfs(grid [][]int, x, y int, p []int) int {
	m := len(grid)
	n := len(grid[0])
	if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 || p[x*n+y] == m*n {
		return 0
	}
	//fmt.Println("test x,y=%v,%v, p=%v", x, y, p[x*n+y])
	//if pxy == m*n {
	//    return 0
	//}
	p[x*n+y] = m * n // from unstable to stable
	res := 1
	dir := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	for _, d := range dir {
		nx, ny := x+d[0], y+d[1]
		res += dfs(grid, nx, ny, p)
	}
	return res
}

func uni2(p []int, x, y int) {
	px := find(p, x)
	py := find(p, y)
	if px > py {
		px, py = py, px
	}
	p[px] = py
}

func find(p []int, x int) int {
	if p[x] != x {
		p[x] = find(p, p[x])
	}
	return p[x]
}
