/*

1879. Minimum XOR Sum of Two Arrays
Solved
Hard
Topics
Companies
Hint
You are given two integer arrays nums1 and nums2 of length n.

The XOR sum of the two integer arrays is (nums1[0] XOR nums2[0]) + (nums1[1] XOR nums2[1]) + ... + (nums1[n - 1] XOR nums2[n - 1]) (0-indexed).

For example, the XOR sum of [1,2,3] and [3,2,1] is equal to (1 XOR 3) + (2 XOR 2) + (3 XOR 1) = 2 + 0 + 2 = 4.
Rearrange the elements of nums2 such that the resulting XOR sum is minimized.

Return the XOR sum after the rearrangement.

 

Example 1:

Input: nums1 = [1,2], nums2 = [2,3]
Output: 2
Explanation: Rearrange nums2 so that it becomes [3,2].
The XOR sum is (1 XOR 3) + (2 XOR 2) = 2 + 0 = 2.
Example 2:

Input: nums1 = [1,0,3], nums2 = [5,3,4]
Output: 8
Explanation: Rearrange nums2 so that it becomes [5,4,3]. 
The XOR sum is (1 XOR 5) + (0 XOR 4) + (3 XOR 3) = 4 + 4 + 0 = 8.
 

Constraints:

n == nums1.length
n == nums2.length
1 <= n <= 14
0 <= nums1[i], nums2[i] <= 107

*/

func minimumXORSum(nums1 []int, nums2 []int) int {
    // it is travel sales man problem
    // let dp[s][i] = xor sum for  state s  plus nums2[i], with nums[i] at the end
    // then we can extend s by adding j
    n := len(nums1)
    dp := make([][]int, 1<<n)
    for i := range dp {
        dp[i] = make([]int, n)
        for j := range dp[i] {
            dp[i][j] = math.MaxInt
        }
    }
    for i, num := range nums2 {
        dp[1<<i][i] = nums1[0] ^ num
    }
    for s := 1; s < 1<<n; s++ {
        cnt := 0
        for i := 0; i < n; i++ {
            if (s >> i) & 1 > 0 {
                cnt++
            }
        }
        for _, sum := range dp[s] {
            if sum == math.MaxInt {
                continue
            }
            for j, num := range nums2 {
                if (s>>j) & 1 > 0 {
                    continue
                }

                ns := s | 1<<j
                nsum := sum + (nums1[cnt] ^ num)
                dp[ns][j] = min(dp[ns][j], nsum)
                //fmt.Println("s=%v, j=%v, cnt=%v, num=%v, sum=%v, nsum=%v", s, j, cnt, num, sum, nsum)
            }
        }
    }
    res := math.MaxInt
    for _, v := range dp[(1<<n)-1] {
        res = min(res, v)
    }
    return res
}
