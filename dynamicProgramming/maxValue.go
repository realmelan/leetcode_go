/*
*/

func maxValue(events [][]int, k int) int {
    n := len(events)
    sort.Slice(events, func(i, j int)bool {
        return events[i][1] < events[j][1]
    })
    dp := make([][]int, k+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    for i, e := range events {
        j := i-1
        for j>= 0 && events[j][1] >= e[0] {
            j--
        }
        for at := 1; at <= k; at++ {
            dp[at][i+1] = max(dp[at][i], dp[at-1][j+1]+e[2])
        }
    }
    return dp[k][n]
}
