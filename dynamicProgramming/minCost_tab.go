/*

1547. Minimum Cost to Cut a Stick
Solved
Hard
Topics
Companies
Hint
Given a wooden stick of length n units. The stick is labelled from 0 to n. For example, a stick of length 6 is labelled as follows:


Given an integer array cuts where cuts[i] denotes a position you should perform a cut at.

You should perform the cuts in order, you can change the order of the cuts as you wish.

The cost of one cut is the length of the stick to be cut, the total cost is the sum of costs of all cuts. When you cut a stick, it will be split into two smaller sticks (i.e. the sum of their lengths is the length of the stick before the cut). Please refer to the first example for a better explanation.

Return the minimum total cost of the cuts.

 

Example 1:


Input: n = 7, cuts = [1,3,4,5]
Output: 16
Explanation: Using cuts order = [1, 3, 4, 5] as in the input leads to the following scenario:

The first cut is done to a rod of length 7 so the cost is 7. The second cut is done to a rod of length 6 (i.e. the second part of the first cut), the third is done to a rod of length 4 and the last cut is to a rod of length 3. The total cost is 7 + 6 + 4 + 3 = 20.
Rearranging the cuts to be [3, 5, 1, 4] for example will lead to a scenario with total cost = 16 (as shown in the example photo 7 + 4 + 3 + 2 = 16).
Example 2:

Input: n = 9, cuts = [5,6,1,4,2]
Output: 22
Explanation: If you try the given cuts ordering the cost will be 25.
There are much ordering with total cost <= 25, for example, the order [4, 6, 5, 2, 1] has total cost = 22 which is the minimum possible.
 

Constraints:

2 <= n <= 106
1 <= cuts.length <= min(n - 1, 100)
1 <= cuts[i] <= n - 1
All the integers in cuts array are distinct.

*/

func minCost(n int, cuts []int) int {
    // suppose their are k cuts, then there are k+1 sticks
    // now imagine you glue them up for a stick including 4 sticks, you can select
    // glue the first two, and then glue the last one; or glue the last two and then
    // glue the first one
    // let dp[i][j] = cost to glue stick i to stick j
    k := len(cuts)
    sort.Ints(cuts)
    sticks := make([]int, k)
    sums := make([]int, k+1)
    prevCut := 0
    for i, c := range cuts {
        sticks[i] = c - prevCut
        prevCut = c
        sums[i+1] = c
    }
    sticks = append(sticks, n-prevCut)
    sums = append(sums, n)
    //fmt.Println("sticks=%v", sticks)
    //fmt.Println("sums=%v", sums)

    n = len(sticks)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    for k := 2; k <= n; k++ { // k=#of sticks
        for i := 0; i+k-1<n; i++ {
            // for each possible glue
            j := i+k-1
            dp[i][j] = math.MaxInt
            for c := i+1; c <= j; c++ {
                dp[i][j] = min(dp[i][j], dp[i][c-1]+dp[c][j]+sums[j+1]-sums[i])
            }
        }
        //fmt.Println("dp=%v", dp)
    }
    return dp[0][n-1]
}
