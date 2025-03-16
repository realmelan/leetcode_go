/*

1563. Stone Game V
Solved
Hard
Topics
Companies
Hint
There are several stones arranged in a row, and each stone has an associated value which is an integer given in the array stoneValue.

In each round of the game, Alice divides the row into two non-empty rows (i.e. left row and right row), then Bob calculates the value of each row which is the sum of the values of all the stones in this row. Bob throws away the row which has the maximum value, and Alice's score increases by the value of the remaining row. If the value of the two rows are equal, Bob lets Alice decide which row will be thrown away. The next round starts with the remaining row.

The game ends when there is only one stone remaining. Alice's is initially zero.

Return the maximum score that Alice can obtain.

 

Example 1:

Input: stoneValue = [6,2,3,4,5,5]
Output: 18
Explanation: In the first round, Alice divides the row to [6,2,3], [4,5,5]. The left row has the value 11 and the right row has value 14. Bob throws away the right row and Alice's score is now 11.
In the second round Alice divides the row to [6], [2,3]. This time Bob throws away the left row and Alice's score becomes 16 (11 + 5).
The last round Alice has only one choice to divide the row which is [2], [3]. Bob throws away the right row and Alice's score is now 18 (16 + 2). The game ends because only one stone is remaining in the row.
Example 2:

Input: stoneValue = [7,7,7,7,7,7,7]
Output: 28
Example 3:

Input: stoneValue = [4]
Output: 0
 

Constraints:

1 <= stoneValue.length <= 500
1 <= stoneValue[i] <= 106

*/

func stoneGameV(stoneValue []int) int {
    n := len(stoneValue)
    sums := make([]int, n+1)
    for i, v := range stoneValue {
        sums[i+1] = sums[i] + v
    }

    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    return dfs(sums, 0, n-1, memo)
}

func dfs(sums []int, l, r int, memo [][]int) int {
    if l>=r {
        return 0
    }
    if memo[l][r] >= 0 {
        return memo[l][r]
    }

    // two rows: l:i, i:r
    res := math.MinInt
    for i := l; i < r; i++ {
        suml := sums[i+1]-sums[l]
        sumr := sums[r+1]-sums[i+1]

        if suml < sumr {
            res = max(res, suml + dfs(sums, l, i, memo))
        } else if suml > sumr {
            res = max(res, sumr + dfs(sums, i+1, r, memo))
        } else {
            t1 := dfs(sums, l, i, memo)
            t2 := dfs(sums, i+1, r, memo)
            res = max(res, suml + max(t1, t2))
        }
    }
    memo[l][r] = res
    return res
}
