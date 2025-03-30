/*
1970. Last Day Where You Can Still Cross
Solved
Hard
Topics
Companies
Hint
There is a 1-based binary matrix where 0 represents land and 1 represents water. You are given integers row and col representing the number of rows and columns in the matrix, respectively.

Initially on day 0, the entire matrix is land. However, each day a new cell becomes flooded with water. You are given a 1-based 2D array cells, where cells[i] = [ri, ci] represents that on the ith day, the cell on the rith row and cith column (1-based coordinates) will be covered with water (i.e., changed to 1).

You want to find the last day that it is possible to walk from the top to the bottom by only walking on land cells. You can start from any cell in the top row and end at any cell in the bottom row. You can only travel in the four cardinal directions (left, right, up, and down).

Return the last day where it is possible to walk from the top to the bottom by only walking on land cells.

 

Example 1:


Input: row = 2, col = 2, cells = [[1,1],[2,1],[1,2],[2,2]]
Output: 2
Explanation: The above image depicts how the matrix changes each day starting from day 0.
The last day where it is possible to cross from top to bottom is on day 2.
Example 2:


Input: row = 2, col = 2, cells = [[1,1],[1,2],[2,1],[2,2]]
Output: 1
Explanation: The above image depicts how the matrix changes each day starting from day 0.
The last day where it is possible to cross from top to bottom is on day 1.
Example 3:


Input: row = 3, col = 3, cells = [[1,2],[2,1],[3,3],[2,2],[1,1],[1,3],[2,3],[3,2],[3,1]]
Output: 3
Explanation: The above image depicts how the matrix changes each day starting from day 0.
The last day where it is possible to cross from top to bottom is on day 3.
 

Constraints:

2 <= row, col <= 2 * 104
4 <= row * col <= 2 * 104
cells.length == row * col
1 <= ri <= row
1 <= ci <= col
All the values of cells are unique.
*/

func latestDayToCross(row int, col int, cells [][]int) int {
    // reversely add cells back and see whether it is top-bottom land connected.
    grid := make([][]byte, row)
    for i := range grid {
        grid[i] = make([]byte, col)
    }
    p := make([]int, row*col+1)
    for i := range p {
        p[i] = i
    }
    for i := len(cells)-1; i >= 0; i-- {
        x, y := cells[i][0]-1, cells[i][1]-1
        id := x*col+y
        if x == 0 {
            uni(p, id, row*col)
        }
        grid[x][y] = 1
        for _, d := range [][]int{{0,1},{0,-1},{1,0},{-1,0}} {
            nx, ny := x+d[0], y+d[1]
            if nx < 0 || nx >= row || ny < 0 || ny >= col || grid[nx][ny] == 0 {
                continue
            }
            uni(p, id, nx*col+ny)
        }

        if connected(row, col, p) {
            return i
        }
    }
    return 0
}

func connected(row, col int, p []int) bool {
    id := row*col
    for i := 0; i < col; i++ {
        id--
        pi := find(p, id)
        if pi == row*col {
            return true
        }
    }
    return false
}

func uni(p []int, x, y int) {
	px := find(p, x)
	py := find(p, y)
	if px > py {
		px, py = py, px
	}
	p[px] = py // connected to larger
}

func find(p []int, x int) int {
	if p[x] != x {
		p[x] = find(p, p[x])
	}
	return p[x]
}
