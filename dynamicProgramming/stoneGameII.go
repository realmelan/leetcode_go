/*

1140. Stone Game II
Solved
Medium
Topics
Companies
Hint
Alice and Bob continue their games with piles of stones. There are a number of piles arranged in a row, and each pile has a positive integer number of stones piles[i]. The objective of the game is to end with the most stones.

Alice and Bob take turns, with Alice starting first.

On each player's turn, that player can take all the stones in the first X remaining piles, where 1 <= X <= 2M. Then, we set M = max(M, X). Initially, M = 1.

The game continues until all the stones have been taken.

Assuming Alice and Bob play optimally, return the maximum number of stones Alice can get.

 

Example 1:

Input: piles = [2,7,9,4,4]

Output: 10

Explanation:

If Alice takes one pile at the beginning, Bob takes two piles, then Alice takes 2 piles again. Alice can get 2 + 4 + 4 = 10 stones in total.
If Alice takes two piles at the beginning, then Bob can take all three piles left. In this case, Alice get 2 + 7 = 9 stones in total.
So we return 10 since it's larger.

Example 2:

Input: piles = [1,2,3,4,5,100]

Output: 104

 

Constraints:

1 <= piles.length <= 100
1 <= piles[i] <= 104

*/

func stoneGameII(piles []int) int {
    // dfs(m, stones, stone_idx) => maximum stones player can get
    n := len(piles)
    sums := make([]int, n+1)
    for i := n-1; i >= 0; i-- {
        sums[i] = sums[i+1]+piles[i]
    }

    memo := make([][]int, n+1)
    for i := range memo {
        memo[i] = make([]int, n+1)
    }

    return dfs(piles, sums, 1, 0, memo)
}

func dfs(piles, sums []int, m, start int, memo [][]int) int {
    n := len(piles)
    if start == n {
        return 0
    }
    if n-start <= m*2 {
        memo[m][start] = sums[start]
        return sums[start]
    }
    // lookup
    if memo[m][start] != 0 {
        return memo[m][start]
    }


    var res int
    for x := 1; x <= m*2 && start+x-1 < n; x++ {
        left := dfs(piles, sums, max(m, x), start+x, memo)
        res = max(res, sums[start]-left)
    }

    memo[m][start] = res
    return res
}
