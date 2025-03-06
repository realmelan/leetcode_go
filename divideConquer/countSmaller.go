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
    n := len(nums)
    res := make([]int, n)
    o := mergeSort(nums, 0, n-1)
    for _, e := range o {
        res[e[1]] = e[2]
    }
    return res
}

// return values contains a tuple for each elements in range [l, r]
// tuple: value, index in nums, # of smaller elements on the right
func mergeSort(nums []int, l, r int) [][]int {
    if l == r {
        return [][]int{
            {nums[l], l, 0},
        }
    }

    m := (l+r)/2
    left := mergeSort(nums, l, m)
    right := mergeSort(nums, m+1, r)

    // now merge left and right
    var res [][]int
    i, ni := 0, len(left)
    j, nj := 0, len(right)
    for ; i < ni; i++ {
        for j < nj && right[j][0] < left[i][0] {
            res = append(res, right[j])
            j++
        }
        left[i][2]+=j
        res = append(res, left[i])
    }
    res = append(res, right[j:nj]...)

    //fmt.Println("l=%v, r=%v, res=%v", l, r, res)
    return res
}
func countSmaller_TLE(nums []int) []int {
    // method 1: use another array to store ordered numbers, starting from the right
    // use binary search to find out # of smaller elements.
    n := len(nums)
    ordered := []int{nums[n-1]}
    res := []int{0}
    for i := n-2; i >= 0; i-- {
        idx, _ := slices.BinarySearch(ordered, nums[i])
        res = append(res, idx)

        ordered = append(ordered, nums[i])
        // shift numbers backward in ordered
        for j := len(ordered)-1; j > idx; j-- {
            ordered[j] = ordered[j-1]
        }
        ordered[idx] = nums[i]
    }
    slices.Reverse(res)
    return res
}
