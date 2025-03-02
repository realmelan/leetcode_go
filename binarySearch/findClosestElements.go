
/*
658. Find K Closest Elements
Solved
Medium
Topics
Companies
Given a sorted integer array arr, two integers k and x, return the k closest integers to x in the array. The result should also be sorted in ascending order.

An integer a is closer to x than an integer b if:

|a - x| < |b - x|, or
|a - x| == |b - x| and a < b
 

Example 1:

Input: arr = [1,2,3,4,5], k = 4, x = 3

Output: [1,2,3,4]

Example 2:

Input: arr = [1,1,2,3,4,5], k = 4, x = -1

Output: [1,1,2,3]

 

Constraints:

1 <= k <= arr.length
1 <= arr.length <= 104
arr is sorted in ascending order.
-104 <= arr[i], x <= 104
*/
func findClosestElements(arr []int, k int, x int) []int {
    // find position in arr that is closed to x
    // then use two pointers to move left or right
    n := len(arr)
    lo, up := 0, n-1
    for lo < up {
        m := (lo+up+1)/2
        if arr[m] <= x {
            lo = m
        } else {
            up=m-1
        }
    }
    // check whether lo == 0 or lo == n-1
    i, j := lo, lo+1
    for j-i-1<k {
        if i < 0 {
            j++
        } else if j >= n {
            i--
        } else {
            di, dj := x - arr[i], arr[j]-x
            if di <= dj {
                i--
            } else {
                j++
            }
        }
    }
    return arr[i+1:j]
}
