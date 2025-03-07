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
    m := make(map[int64]bool)
    var res []int64
    for _, num := range nums {
        v := int64(num)
        if !m[v] {
            m[v]=true
            res = append(res, v)
        }

        if v > 0 {
            v = (v-1)/2
        } else {
            v = (v-2)/2
        }
        if !m[v] {
            m[v]=true
            res = append(res, v)
        }
    }
    slices.Sort(res)
    idm := make(map[int64]int)
    for i, v := range res {
        idm[v]=i+1
    }

    var cnt int
    bit := Bit{
        N: len(res)+1,
        count: make([]int, len(res)+2),
    }

    n := len(nums)
    bit.update(idm[int64(nums[n-1])])
    for i := n-2; i >= 0; i-- {
        v := int64(nums[i])
        t := (v - 1)/2
        if v <= 0 {
            t = (v-2)/2
        }
        cnt += bit.query(idm[t])
        bit.update(idm[v])
    }
    return cnt
}

type Bit struct {
    N int
    count []int
}

func lsb(x int) int {
    return x & (-x)
}

func (b *Bit) update(id int) {
    for id < b.N {
        b.count[id]++
        id += lsb(id)
    }
}

func (b *Bit) query(id int) int {
    var res int
    for id > 0 {
        res += b.count[id]
        id -= lsb(id)
    }
    return res
}

func reversePairs_dc(nums []int) int {
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
