/*

473. Matchsticks to Square
Solved
Medium
Topics
Companies
Hint
You are given an integer array matchsticks where matchsticks[i] is the length of the ith matchstick. You want to use all the matchsticks to make one square. You should not break any stick, but you can link them up, and each matchstick must be used exactly one time.

Return true if you can make this square and false otherwise.

 

Example 1:


Input: matchsticks = [1,1,2,2,2]
Output: true
Explanation: You can form a square with length 2, one side of the square came two sticks with length 1.
Example 2:

Input: matchsticks = [3,3,3,3,4]
Output: false
Explanation: You cannot find a way to form a square with all the matchsticks.
 

Constraints:

1 <= matchsticks.length <= 15
1 <= matchsticks[i] <= 108

*/

func makesquare(matchsticks []int) bool {
    sum := 0
    m := matchsticks
    n := len(m)
    for _, l := range m {
        sum += l
    }
    if sum % 4 != 0 {
        return false
    }
    for _, l := range m {
        if l > sum/4 {
            return false
        }
    }

    t := sum/4
    dp := make([]int, 1<<n)
    for i := range dp {
        dp[i] = -1
    }
    dp[0] = 0
    for s := 0; s < 1<<n; s++ {
        if dp[s] < 0 {
            continue
        }
        for i, match := range matchsticks {
            if (s >> i) & 1 > 0 {
                continue
            }
            ns := s |1<<i
            if dp[s] + match > t {
                continue
            }

            dp[ns] = (dp[s] + match)%t
        }
    }
    return dp[(1<<n)-1] != -1
}
