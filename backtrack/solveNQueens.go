/*
51. N-Queens
Solved
Hard
Topics
Companies
The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.

 

Example 1:


Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above
Example 2:

Input: n = 1
Output: [["Q"]]
 

Constraints:

1 <= n <= 9
*/
func solveNQueens(n int) [][]string {
    var res [][]string
    q := make([]int, n)
    for i := 0; i < n; i++ {
        q[0] = i
        backtrack(q, n, 1, &res)
    }
    return res
}

func backtrack(q []int, n, row int, res *[][]string) {
    if row >= n {
        *res = append(*res, format(n, q))
        return
    }

    // check coloumn i of row
    for i := 0; i < n; i++ {
        j := 0
        for j < row { // check each previous row
            if !valid(j, q[j], row, i) {
                break
            }
            j++
        }

        if j >= row {
            q[row] = i
            backtrack(q, n, row+1, res)
        }
    }
}

func valid(qi, qj, pi, pj int) bool {
    return !(qi == pi || qj == pj || pi - qi == pj - qj || pi - qi == qj - pj)
}

func format(n int, q []int) []string {
    var res []string
    for i := 0; i < n; i++ {
        var buf strings.Builder
        for j := 0; j < q[i]; j++ {
            buf.WriteByte('.')
        }
        buf.WriteByte('Q')
        for j := q[i]+1; j < n; j++ {
            buf.WriteByte('.')
        }
        res = append(res, buf.String())
    }
    return res
}
