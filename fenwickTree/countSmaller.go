/*
315. Count of Smaller Numbers After Self
Solved
Hard
Topics
Companies
Given an integer array nums, return an integer array counts where counts[i] is the number of smaller elements to the right of nums[i].

 

Example 1:

Input: nums = [5,2,6,1]
Output: [2,1,1,0]
Explanation:
To the right of 5 there are 2 smaller elements (2 and 1).
To the right of 2 there is only 1 smaller element (1).
To the right of 6 there is 1 smaller element (1).
To the right of 1 there is 0 smaller element.
Example 2:

Input: nums = [-1]
Output: [0]
Example 3:

Input: nums = [-1,-1]
Output: [0,0]
 

Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104
*/

func countSmaller(nums []int) []int {
    // Binary indexed tree
    n := len(nums)
    mi, ma := 100000, -100000
    for _, num := range nums {
        mi = min(mi, num)
    }
    for i := range nums {
        nums[i] -= mi-1
        ma = max(ma, nums[i])
    }

    res := []int{0}
    bit := Bit{
        N: ma+1,
        Counts: make([]int, ma+1),
    }
    bit.update(nums[n-1])
    for i := n-2; i >= 0; i-- {
        res = append(res, bit.query(nums[i]-1))
        bit.update(nums[i])
    }
    slices.Reverse(res)
    return res
}

type Bit struct {
    N int
    Counts []int
}

func lsb(i int) int {
    return i & (-i)
}
func (b *Bit) query(num int) int {
    res := 0
    for num > 0 {
        res += b.Counts[num]
        num -= lsb(num)
    }
    return res
}
func (b * Bit) update(num int) {
    for num < b.N {
        b.Counts[num]++
        num += lsb(num)
    }
}
