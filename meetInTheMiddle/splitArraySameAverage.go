/*

805. Split Array With Same Average
Solved
Hard
Topics
Companies
You are given an integer array nums.

You should move each element of nums into one of the two arrays A and B such that A and B are non-empty, and average(A) == average(B).

Return true if it is possible to achieve that and false otherwise.

Note that for an array arr, average(arr) is the sum of all the elements of arr over the length of arr.

 

Example 1:

Input: nums = [1,2,3,4,5,6,7,8]
Output: true
Explanation: We can split the array into [1,4,5,8] and [2,3,6,7], and both of them have an average of 4.5.
Example 2:

Input: nums = [3,1]
Output: false
 

Constraints:

1 <= nums.length <= 30
0 <= nums[i] <= 104

*/

func splitArraySameAverage(nums []int) bool {
    total, n := 0, len(nums)
    if n <= 1 {
        return false
    }
    for _, num := range nums {
        total += num
    }
    if total == 0 {
        return true
    }

    //fmt.Println("total=%v, tn=%v, avg=%v", total, n, float64(total)/float64(n))
    m := n/2
    left, exist1 := subsetSums(nums[0:m], total, n)
    if exist1 {
        return true
    }
    right, exist2 := subsetSums(nums[m:n], total, n)
    if exist2 {
        return true
    }

    //fmt.Println("left=%v", left)
    //fmt.Println("right=%v", right)
    for _, i := range left {
        for _, j := range right {
            if compare(i, j, total, n) == 0 {
                return true
            }
        }
    }
    return false
}

func subsetSums(nums []int, total, tn int) (res [][]int, exist bool) {
    n := len(nums)
    m := make(map[int][]bool)
    res = make([][]int, 0, 1<<n)
    for i := 1; i < 1<<n; i++ {
        sum, cnt := 0, 0
        for j := 0; j < n; j++ {
            if (i>>j)&1 > 0 {
                cnt++
                sum += nums[j]
            }
        }
        if cnt * total == tn * sum {
            return nil, true
        }
        if cnt != n {
            _, ok := m[sum]
            if !ok {
                m[sum] = make([]bool, n+1)
            }
            if !m[sum][cnt] {
                res = append(res, []int{sum, cnt})
                m[sum][cnt] = true
            }
        }
    }
    return
}

func compare(left, right []int, total, tn int) int {
    sum := left[0] + right[0]
    cnt := left[1] + right[1]
    return sum * tn - total * cnt
}
