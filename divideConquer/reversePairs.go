/*
493. Reverse Pairs
Solved
Hard
Topics
Companies
Hint
Given an integer array nums, return the number of reverse pairs in the array.

A reverse pair is a pair (i, j) where:

0 <= i < j < nums.length and
nums[i] > 2 * nums[j].
 

Example 1:

Input: nums = [1,3,2,3,1]
Output: 2
Explanation: The reverse pairs are:
(1, 4) --> nums[1] = 3, nums[4] = 1, 3 > 2 * 1
(3, 4) --> nums[3] = 3, nums[4] = 1, 3 > 2 * 1
Example 2:

Input: nums = [2,4,3,5,1]
Output: 3
Explanation: The reverse pairs are:
(1, 4) --> nums[1] = 4, nums[4] = 1, 4 > 2 * 1
(2, 4) --> nums[2] = 3, nums[4] = 1, 3 > 2 * 1
(3, 4) --> nums[3] = 5, nums[4] = 1, 5 > 2 * 1
 

Constraints:

1 <= nums.length <= 5 * 104
-231 <= nums[i] <= 231 - 1
*/

func reversePairs(nums []int) int {
    _, res := dcon(nums, 0, len(nums)-1)
    return res
}

func dcon(nums []int, l, r int) (arr []int64, cnt int) {
    if l >= r {
        return []int64{int64(nums[l])}, 0
    }
    m := (l+r)/2
    left, liv := dcon(nums, l, m)
    right, riv := dcon(nums, m+1, r)

    cnt = liv + riv
    i, ni := 0, len(left)
    j, nj := 0, len(right)
    for i < ni {
        for j < nj && right[j] * 2 < left[i] {
            j++
        }
        cnt += j
        i++
    }

    // now merge left and right
    i, j = 0, 0
    for i < ni {
        for j < nj && right[j] < left[i] {
            arr = append(arr, right[j])
            j++
        }
        arr = append(arr, left[i])
        i++
    }
    arr = append(arr, right[j:nj]...)

    //fmt.Println("arr=%v, cnt=%v", arr, cnt)
    return 
}
