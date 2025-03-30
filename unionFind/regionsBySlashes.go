/*

959. Regions Cut By Slashes
Solved
Medium
Topics
Companies
An n x n grid is composed of 1 x 1 squares where each 1 x 1 square consists of a '/', '\', or blank space ' '. These characters divide the square into contiguous regions.

Given the grid grid represented as a string array, return the number of regions.

Note that backslash characters are escaped, so a '\' is represented as '\\'.

 

Example 1:


Input: grid = [" /","/ "]
Output: 2
Example 2:


Input: grid = [" /","  "]
Output: 1
Example 3:


Input: grid = ["/\\","\\/"]
Output: 5
Explanation: Recall that because \ characters are escaped, "\\/" refers to \/, and "/\\" refers to /\.
 

Constraints:

n == grid.length == grid[i].length
1 <= n <= 30
grid[i][j] is either '/', '\', or ' '.

*/

func regionsBySlashes(grid []string) int {
    n := len(grid)
    p := make([]int, n*n*4)
    for i := range p {
        p[i] = i
    }
    for r := 0; r < n; r++ {
        for c := 0; c < n; c++ {
            b := grid[r][c]
            id := (r*n+c)*4
            if c < n-1 {
                uni(p, id+2, id+4)
            }
            if r < n-1 {
                uni(p, id+3, id+n*4+1)
            }
            if b == ' ' {
                uni(p, id, id+1)
                uni(p, id+1, id+2)
                uni(p, id+2, id+3)
            } else if b == '/' {
                uni(p, id, id+1)
                uni(p, id+2, id+3)
            } else {
                uni(p, id, id+3)
                uni(p, id+1, id+2)
            }
        }
    }
    root := make(map[int]bool)
    for i := range p {
        root[find(p, i)]=true
    }
    return len(root)
}

func uni(p []int, x, y int) {
    p[find(p, x)] = find(p, y)
}
func find(p []int, x int) int {
    if p[x] != x {
        p[x] = find(p, p[x])
    }
    return p[x]
}
