/*

212. Word Search II
Solved
Hard
Topics
Companies
Hint
Given an m x n board of characters and a list of strings words, return all words on the board.

Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

 

Example 1:


Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]
Example 2:


Input: board = [["a","b"],["c","d"]], words = ["abcb"]
Output: []
 

Constraints:

m == board.length
n == board[i].length
1 <= m, n <= 12
board[i][j] is a lowercase English letter.
1 <= words.length <= 3 * 104
1 <= words[i].length <= 10
words[i] consists of lowercase English letters.
All the strings of words are unique.

*/

func findWords(board [][]byte, words []string) []string {
    mprefix := make(map[string]int)
    mword := make(map[string]int)
    mlen := 0
    for _, w := range words {
        l := len(w)
        mlen = max(mlen, l)
        for i := l; i > 0; i-- {
            mprefix[w[:i]]++
        }
        mword[w]=1
    }
    m, n := len(board), len(board[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            used := make(map[int]bool)
            dfs(board, mprefix, mword, used, "", i, j, mlen)
        }
    }
    var res []string
    for k, v := range mword {
        if v == 2 {
            res = append(res, k)
        }
    }
    return res
}

func dfs(board [][]byte, mprefix map[string]int, mword map[string]int, used map[int]bool, p string, x, y, mlen int) {
    m, n := len(board), len(board[0])
    id := x*n+y
    if x < 0 || x >= m || y < 0 || y >= n || used[id] {
        return
    }
    np := p + string([]byte{board[x][y]})
    if len(np) > mlen {
        return
    }
    if mprefix[np] <= 0 {
        return
    }
    if mword[np] == 1 {
        mword[np] = 2
        for i := len(np); i > 0; i-- {
            mprefix[np[:i]]--
        }
    }
    used[id]=true
    dfs(board, mprefix, mword, used, np, x+1, y, mlen)
    dfs(board, mprefix, mword, used, np, x-1, y, mlen)
    dfs(board, mprefix, mword, used, np, x, y+1, mlen)
    dfs(board, mprefix, mword, used, np, x, y-1, mlen)
    used[id]=false
}
