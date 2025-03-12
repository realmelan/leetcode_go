/*
1406. Stone Game III
Solved
Hard
Topics
Companies
Hint
Alice and Bob continue their games with piles of stones. There are several stones arranged in a row, and each stone has an associated value which is an integer given in the array stoneValue.

Alice and Bob take turns, with Alice starting first. On each player's turn, that player can take 1, 2, or 3 stones from the first remaining stones in the row.

The score of each player is the sum of the values of the stones taken. The score of each player is 0 initially.

The objective of the game is to end with the highest score, and the winner is the player with the highest score and there could be a tie. The game continues until all the stones have been taken.

Assume Alice and Bob play optimally.

Return "Alice" if Alice will win, "Bob" if Bob will win, or "Tie" if they will end the game with the same score.

 

Example 1:

Input: stoneValue = [1,2,3,7]
Output: "Bob"
Explanation: Alice will always lose. Her best move will be to take three piles and the score become 6. Now the score of Bob is 7 and Bob wins.
Example 2:

Input: stoneValue = [1,2,3,-9]
Output: "Alice"
Explanation: Alice must choose all the three piles at the first move to win and leave Bob with negative score.
If Alice chooses one pile her score will be 1 and the next move Bob's score becomes 5. In the next move, Alice will take the pile with value = -9 and lose.
If Alice chooses two piles her score will be 3 and the next move Bob's score becomes 3. In the next move, Alice will take the pile with value = -9 and also lose.
Remember that both play optimally so here Alice will choose the scenario that makes her win.
Example 3:

Input: stoneValue = [1,2,3,6]
Output: "Tie"
Explanation: Alice cannot win this game. She can end the game in a draw if she decided to choose all the first three piles, otherwise she will lose.
 

Constraints:

1 <= stoneValue.length <= 5 * 104
-1000 <= stoneValue[i] <= 1000
*/

func stoneGameIII(stoneValues []int) string {
    // let dp[i] = max value a player gets for stoneValues[i:n]
    // then dp[i] = max{stoneValues[i]+sums[i+1]-dp[i+1], stoneValues[i:i+2]+sums[i+2] - dp[i+2], stoneValues[i:i+3]+sums[i+3]-dp[i+3]}
    // starting from n backward to 0
    n := len(stoneValues)
    sums := make([]int, n+1)
    for i := n-1; i >= 0; i-- {
        sums[i] = sums[i+1] + stoneValues[i]
    }
    dp := make([]int, n)
    for i := range dp {
        dp[i] = math.MinInt
    }
    dp[n-1] = sums[n-1]
    for i := n-2; i >= 0; i-- {
        sum := 0
        for j := 0; j < 3 && i+j<n; j++ {
            sum += stoneValues[i+j]
            if i+j+1 < n {
                dp[i] = max(dp[i], sum + sums[i+j+1] - dp[i+j+1])
            } else {
                dp[i] = max(dp[i], sum)
            }
        }
    }
    //fmt.Println("dp=%v", dp)
    if dp[0] > sums[0]-dp[0] {
        return "Alice"
    } else if dp[0] < sums[0] - dp[0] {
        return "Bob"
    } else {
        return "Tie"
    }
}

func stoneGameIII_dfs(stoneValues []int) string {
    n := len(stoneValues)
    sums := make([]int, n+1)
    for i := n-1; i >= 0; i-- {
        sums[i] = sums[i+1] + stoneValues[i]
    }

    memo := make([]int, n)
    for i := range memo {
        memo[i] = math.MinInt
    }
    first := dfs(stoneValues, sums, 0, &memo)
    if first > sums[0] - first {
        return "Alice"
    } else if first < sums[0] - first {
        return "Bob"
    } else {
        return "Tie"
    }
}

func dfs(vals, sums []int, start int, memo *[]int) int {
    n := len(vals)
    if start == n {
        return 0
    } else if start == n-1 {
        return vals[n-1]
    }
    if (*memo)[start] != math.MinInt {
        return (*memo)[start]
    }


    ires, res := 0, math.MinInt
    for i := 0; i < min(3, n-start); i++ {
        ires += vals[start+i]
        rival := dfs(vals, sums, start+i+1, memo)
        res = max(res, ires + sums[start+i+1]-rival)
    }
    (*memo)[start] = res
    return res
}
