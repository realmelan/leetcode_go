/*
215. Kth Largest Element in an Array
Solved
Medium
Topics
Companies
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?

 

Example 1:

Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
Example 2:

Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4
 

Constraints:

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/
func findKthLargest(nums []int, k int) int {
    k = len(nums) - k
    p, q := 0, len(nums)-1
    for p < q {
        j := partition(nums, p, q)
        //fmt.Println("j=%v, p=%v, q=%v, nums=%v", j, p, q, nums)
        if j == k {
            break
        } else if j < k {
            p = j+1
        } else {
            q = j-1
        }
    }
    return nums[k]
}

// Hoare's partition algorithm
func partition(nums []int, p, q int) int {
    pivot := nums[p]
    i, j := p+1, q
    for {
        if i < q && nums[i] < pivot {
            i++
        } else if j > p && nums[j] > pivot { // has to be >, not >=
            j--
        } else if i>=j {
            break
        } else {
            nums[i], nums[j] = nums[j], nums[i]
            i++
            j--
        }
    }
    nums[p], nums[j] = nums[j], nums[p]
    //fmt.Println("partition: pivot=%v, j=%v, p=%v, q=%v, nums=%v", pivot, j, p, q, nums)
    return j
}
