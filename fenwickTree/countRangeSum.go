/*
327. Count of Range Sum
Solved
Hard
Topics
Companies
Given an integer array nums and two integers lower and upper, return the number of range sums that lie in [lower, upper] inclusive.

Range sum S(i, j) is defined as the sum of the elements in nums between indices i and j inclusive, where i <= j.

 

Example 1:

Input: nums = [-2,5,-1], lower = -2, upper = 2
Output: 3
Explanation: The three ranges are: [0,0], [2,2], and [0,2] and their respective sums are: -2, -1, 2.
Example 2:

Input: nums = [0], lower = 0, upper = 0
Output: 1
 

Constraints:

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
-105 <= lower <= upper <= 105
The answer is guaranteed to fit in a 32-bit integer.
*/

func countRangeSum_FenwickTree(nums []int, lower int, upper int) int {
    sums := []int{0}
    sum := 0
    for _, num := range nums {
        sum += num
        sums = append(sums, sum)
        sums = append(sums, sum-upper-1)
        sums = append(sums, sum-lower)
    }
    sort.Ints(sums)
    m := make(map[int]int)
    for _, sum := range sums {
        _, found := m[sum]
        if !found {
            m[sum] = len(m)+1
        }
    }

    bit := BIT{
        n: len(m)+2,
        F: make([]int, len(m)+2),
    }
    res := 0
    sum = 0
    bit.update(m[0], 1)
    for _, num := range nums {
        sum += num
        idx := m[sum]
        // find idx of sum-upper-1 and sum-lower
        idup := m[sum-upper-1]
        idlo := m[sum-lower]
        // range sum of (idup, idlo]
        res += bit.query(idlo) - bit.query(idup)
        bit.update(idx, 1)
    }
    return res
}

type BIT struct {
    n int
    F []int
}

func lsb(x int) int {
    return x & -x
}
func (b *BIT) query(id int) int {
    res := 0
    for id > 0 {
        res += b.F[id]
        id -= lsb(id)
    }
    return res
}

func (b *BIT) update(id int, val int) {
    for id < b.n {
        b.F[id] += val
        id += lsb(id)
    }
}

func countRangeSum_TLE(nums []int, lower int, upper int) int {
    pos := make([][]int, 1)
    pos[0] = []int{0,-1}
    m := make(map[int]int)
    m[0] = 0
    sum := 0
    for i, num := range nums {
        sum += num
        idx, found := m[sum]
        if !found {
            m[sum] = len(pos)
            pos = append(pos, []int{sum, i})
        } else {
            pos[idx] = append(pos[idx], i)
        }
    }

    sort.Slice(pos, func(i,j int)bool{
        return pos[i][0] < pos[j][0]
    })
    //fmt.Println("pos=%v", pos)

    n := len(pos)
    res, k := 0, 0
    for _, p := range pos {
        // find j, k such that m[]
        for k < n && pos[k][0] < p[0]+lower {
            k++
        }
        if k >= n {
            break
        }
        
        j := k
        for j < n && pos[j][0] <= p[0]+upper {
            res += count(p, pos[j])
            j++
        }
    }
    return res
}

func count(p1, p2 []int) int {
    //fmt.Println("p1=%v, p2=%v", p1, p2)
    n1, n2 := len(p1), len(p2)
    j, res := 1, 0
    for i := 1; i < n1; i++ {
        for j < n2 && p2[j] <= p1[i] {
            j++
        }
        res += n2 - j
    }
    //fmt.Println("res=%v", res)
    return res
}
func countRangeSum(nums []int, lower int, upper int) int {
    pos := make([][]int, 1)
    pos[0] = []int{0,-1}
    sum := 0
    for i, num := range nums {
        sum += num
        pos = append(pos, []int{sum, i})
    }

    sort.Slice(pos, func(i,j int)bool{
        return pos[i][0] < pos[j][0]
    })
    //fmt.Println("pos=%v", pos)

    n := len(pos)
    res, k := 0, 0
    for _, p := range pos {
        // find j, k such that m[]
        for k < n && pos[k][0] < p[0]+lower {
            k++
        }
        if k >= n {
            break
        }
        
        j := k
        for j < n && pos[j][0] <= p[0]+upper {
            if pos[j][1] > p[1] {
                res++
            }
            j++
        }
    }
    return res
}
