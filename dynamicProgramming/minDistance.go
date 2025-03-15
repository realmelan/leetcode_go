/*

1478. Allocate Mailboxes
Solved
Hard
Topics
Companies
Hint
Given the array houses where houses[i] is the location of the ith house along a street and an integer k, allocate k mailboxes in the street.

Return the minimum total distance between each house and its nearest mailbox.

The test cases are generated so that the answer fits in a 32-bit integer.

 

Example 1:


Input: houses = [1,4,8,10,20], k = 3
Output: 5
Explanation: Allocate mailboxes in position 3, 9 and 20.
Minimum total distance from each houses to nearest mailboxes is |3-1| + |4-3| + |9-8| + |10-9| + |20-20| = 5 
Example 2:


Input: houses = [2,3,5,12,18], k = 2
Output: 9
Explanation: Allocate mailboxes in position 3 and 14.
Minimum total distance from each houses to nearest mailboxes is |2-3| + |3-3| + |5-3| + |12-14| + |18-14| = 9.
 

Constraints:

1 <= k <= houses.length <= 100
1 <= houses[i] <= 104
All the integers of houses are unique.

*/

func minDistance(houses []int, k int) int {
    // let dp[i][k] = distance of setting up k mailboxes for the first i houses, ignoring the houses after the ith
    // then dp[i][k] = min(dp[i-j][k-1] + distance(j+1:i))
    n := len(houses)
    sort.Ints(houses)
    //fmt.Println("houses=%v", houses)
    dp := make([][]int, k+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    // x = 1
    for i := 0; i < n; i++ {
        dp[1][i] = dist(houses, 0, i)
    }

    for x := 2; x <= k; x++ {
        for i := x; i < n; i++ { // dp[x][0:x-1] = 0
            dp[x][i] = math.MaxInt
            for j := x-2; j < i; j++ {
                dp[x][i] = min(dp[x][i], dp[x-1][j]+dist(houses, j+1, i))
            }
        }
        //fmt.Println("dp[%v]=%v", x, dp[x])
    }
    return dp[k][n-1]
}

func dist(houses []int, i, j int) int {
    res := 0
    for i < j {
        res += houses[j]-houses[i]
        i++
        j--
    }
    return res
}
