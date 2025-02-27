/*
668. Kth Smallest Number in Multiplication Table
Solved
Hard
Topics
Companies
Nearly everyone has used the Multiplication Table. The multiplication table of size m x n is an integer matrix mat where mat[i][j] == i * j (1-indexed).

Given three integers m, n, and k, return the kth smallest element in the m x n multiplication table.

 

Example 1:


Input: m = 3, n = 3, k = 5
Output: 3
Explanation: The 5th smallest number is 3.
Example 2:


Input: m = 2, n = 3, k = 6
Output: 6
Explanation: The 6th smallest number is 6.
 

Constraints:

1 <= m, n <= 3 * 104
1 <= k <= m * n
*/

func findKthNumber(m int, n int, k int) int {
    lo, up := 1, m*n
    for lo < up {
        mid := (lo + up) / 2
        le := check(m, n, mid)
        //fmt.Println("mid=%v, le=%v, closed=%v", mid, le, closed)
        if le < k {
            lo = mid+1
        } else {
            up = mid
        }
    }

    return up
}

func check(m, n, t int) (le int) {
    for i := 1; i <= m; i++ {
        v := min(n, t/i)
        le += v
    }
    return
}

func findKthNumber(m int, n int, k int) int {
    lo, up := 1, m*n
    res := up
    for lo <= up {
        mid := (lo + up) / 2
        le, closed := check(m, n, mid)
        //fmt.Println("mid=%v, le=%v, closed=%v", mid, le, closed)
        if le < k {
            lo = mid+1
        } else {
            res = min(res, closed)
            up = mid-1
        }
    }

    return res
}

func check(m, n, t int) (le, closed int) {
    for i := 1; i <= m; i++ {
        v := min(n, t/i)
        le += v
        closed = max(closed, v*i)
    }
    return
}
