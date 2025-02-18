/*
31. Next Permutation
Solved
Medium
Topics
Companies
A permutation of an array of integers is an arrangement of its members into a sequence or linear order.

For example, for arr = [1,2,3], the following are all the permutations of arr: [1,2,3], [1,3,2], [2, 1, 3], [2, 3, 1], [3,1,2], [3,2,1].
The next permutation of an array of integers is the next lexicographically greater permutation of its integer. More formally, if all the permutations of the array are sorted in one container according to their lexicographical order, then the next permutation of that array is the permutation that follows it in the sorted container. If such arrangement is not possible, the array must be rearranged as the lowest possible order (i.e., sorted in ascending order).

For example, the next permutation of arr = [1,2,3] is [1,3,2].
Similarly, the next permutation of arr = [2,3,1] is [3,1,2].
While the next permutation of arr = [3,2,1] is [1,2,3] because [3,2,1] does not have a lexicographical larger rearrangement.
Given an array of integers nums, find the next permutation of nums.

The replacement must be in place and use only constant extra memory.

 

Example 1:

Input: nums = [1,2,3]
Output: [1,3,2]
Example 2:

Input: nums = [3,2,1]
Output: [1,2,3]
Example 3:

Input: nums = [1,1,5]
Output: [1,5,1]
 

Constraints:

1 <= nums.length <= 100
0 <= nums[i] <= 100
*/
func nextPermutation(nums []int)  {
    n := len(nums)
    k := n-1
    for k > 0 && nums[k] <= nums[k-1] {
        k--
    }
    if k == 0 {
        // swap elements
        swap(nums, 0, n-1)
        return
    }

    j := n-1
    for j >= k && nums[j] <= nums[k-1] {
        j--
    }
    target := nums[k-1]
    nums[k-1] = nums[j]
    // nums[j] = target
    j++
    for j < n && nums[j] > target {
        nums[j-1] = nums[j]
    }
    nums[j-1] = target
    swap(nums, k, n-1)
}

func swap(nums []int, i,j int) {
    for i < j {
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
}
