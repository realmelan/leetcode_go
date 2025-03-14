/*

1866. Number of Ways to Rearrange Sticks With K Sticks Visible
Solved
Hard
Topics
Companies
Hint
There are n uniquely-sized sticks whose lengths are integers from 1 to n. You want to arrange the sticks such that exactly k sticks are visible from the left. A stick is visible from the left if there are no longer sticks to the left of it.

For example, if the sticks are arranged [1,3,2,5,4], then the sticks with lengths 1, 3, and 5 are visible from the left.
Given n and k, return the number of such arrangements. Since the answer may be large, return it modulo 109 + 7.

 

Example 1:

Input: n = 3, k = 2
Output: 3
Explanation: [1,3,2], [2,3,1], and [2,1,3] are the only arrangements such that exactly 2 sticks are visible.
The visible sticks are underlined.
Example 2:

Input: n = 5, k = 5
Output: 1
Explanation: [1,2,3,4,5] is the only arrangement such that all 5 sticks are visible.
The visible sticks are underlined.
Example 3:

Input: n = 20, k = 11
Output: 647427950
Explanation: There are 647427950 (mod 109 + 7) ways to rearrange the sticks such that exactly 11 sticks are visible.
 

Constraints:

1 <= n <= 1000
1 <= k <= n

*/

func rearrangeSticks(n int, k int) int {
    // let dp[j][i] = # of arrangements of 1 to i, to have j visible sticks
    // then dp[j][i] = dp[j-1][i-1](stick-i at last position) + (i-1)*dp[j][i-1] (other stick at last position, i-1 choices)
    // initial set: dp[1][1..n] = 1
    const mod = int64(1000000007)
    dp := make([][]int64, k+1)
    for i := range dp {
        dp[i] = make([]int64, n+1)
    }
    dp[1][1] = 1
    for i := 2; i <= n; i++ {
        dp[1][i] = dp[1][i-1]*int64(i-1)%mod
    }
    for j := 2; j <= k; j++ {
        for i := j; i <= n; i++ {
            dp[j][i] += dp[j-1][i-1] + int64(i-1)*dp[j][i-1]
            dp[j][i] %= mod
        }
    }
    return int(dp[k][n])
}
