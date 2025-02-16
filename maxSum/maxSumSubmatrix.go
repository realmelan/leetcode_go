/*
363. Max Sum of Rectangle No Larger Than K
Solved
Hard
Topics
Companies

Given an m x n matrix matrix and an integer k, return the max sum of a rectangle in the matrix such that its sum is no larger than k.

It is guaranteed that there will be a rectangle with a sum no larger than k.
*/
func maxSumSubmatrix(matrix [][]int, k int) int {
    m, n := len(matrix), len(matrix[0])
    res := math.MinInt

    for i := 0; i < m; i++ {
        sums := make([]int, n)
        for j := i; j < m; j++ {
            for l := 0; l < n; l++ {
                sums[l] += matrix[j][l]
            }

            // use sort.Search to find
            sorted := []int{0}
            curSum := 0
            for _, sum := range sums {
                curSum += sum
                idx := sort.Search(len(sorted), func(p int) bool {
                    return curSum - sorted[p] <= k
                })
                if idx < len(sorted) {
                    res = max(res, curSum - sorted[idx])
                }
                sorted = append(sorted, curSum)
                sort.Ints(sorted)
            }
        }
    }
    return res
}
