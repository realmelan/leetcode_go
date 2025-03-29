/*

130. Surrounded Regions
Solved
Medium
Topics
Companies
You are given an m x n matrix board containing letters 'X' and 'O', capture regions that are surrounded:

Connect: A cell is connected to adjacent cells horizontally or vertically.
Region: To form a region connect every 'O' cell.
Surround: The region is surrounded with 'X' cells if you can connect the region with 'X' cells and none of the region cells are on the edge of the board.
To capture a surrounded region, replace all 'O's with 'X's in-place within the original board. You do not need to return anything.

 

Example 1:

Input: board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]

Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]

Explanation:


In the above diagram, the bottom region is not captured because it is on the edge of the board and cannot be surrounded.

Example 2:

Input: board = [["X"]]

Output: [["X"]]

 

Constraints:

m == board.length
n == board[i].length
1 <= m, n <= 200
board[i][j] is 'X' or 'O'.

*/

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	p := make([]int, m*n+1)
	p[m*n] = m * n
	id := 0
	for r, row := range board {
		for c, b := range row {
			p[id] = id
			if b == 'O' && (r == 0 || r == m-1 || c == 0 || c == n-1) {
				p[id] = m * n
			}
			id++
		}
	}

	id = 0
	dir := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	for r, row := range board {
		for c, b := range row {
			if b != 'O' {
				id++
				continue
			}
			for _, d := range dir {
				x, y := r+d[0], c+d[1]
				if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
					continue
				}
				nid := id + d[0]*n + d[1]
				p1 := find(p, id)
				p2 := find(p, nid)
				if p1 > p2 {
					p1, p2 = p2, p1
				}
				p[p1] = p2
			}
			id++
		}
	}
	id = 0
	for r, row := range board {
		for c, b := range row {
			if b != 'O' {
				id++
				continue
			}

			p1 := find(p, id)
			if p1 != m*n {
				board[r][c] = 'X'
			}
			id++

		}
	}

}
func find(p []int, id int) int {
	if p[id] == id {
		return id
	}
	p[id] = find(p, p[id])
	return p[id]
}
