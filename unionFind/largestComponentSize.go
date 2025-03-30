/*
952. Largest Component Size by Common Factor
Solved
Hard
Topics
Companies
You are given an integer array of unique positive integers nums. Consider the following graph:

There are nums.length nodes, labeled nums[0] to nums[nums.length - 1],
There is an undirected edge between nums[i] and nums[j] if nums[i] and nums[j] share a common factor greater than 1.
Return the size of the largest connected component in the graph.

 

Example 1:


Input: nums = [4,6,15,35]
Output: 4
Example 2:


Input: nums = [20,50,9,63]
Output: 2
Example 3:


Input: nums = [2,3,6,7,4,12,21,39]
Output: 8
 

Constraints:

1 <= nums.length <= 2 * 104
1 <= nums[i] <= 105
All the values of nums are unique.
*/

func largestComponentSize(nums []int) int {
    // use union find on prime factors
    n := len(nums)
    p := make([]int, n)
    for i := range n {
        p[i] = i
    }
    m := make(map[int][]int)
    for i, num := range nums {
        if num == 1 {
            continue
        }
        prs:=make(map[int]bool)
        findPrimes2(num, prs)
        for pr := range prs {
            m[pr] = append(m[pr], i)
        }
    }
    for _, l := range m {
        for i := 1; i < len(l); i++ {
            uni(p, l[0], l[i])
        }
    }
    size := make(map[int]int)
    for i := range nums {
        //fmt.Println("num=%v, pr=%v", num, pr)
        size[find(p, i)]++
    }

    res := 0
    for _, v := range size {
        res = max(res, v)
    }
    return res
}

func findPrimes2(num int, m map[int]bool) {
    for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
        if num % i == 0 {
            m[i] = true
            findPrimes2(num/i, m)
            return
        }
    }
    m[num]=true
}

func largestComponentSize_TLE(nums []int) int {
    n := len(nums)
    p := make([]int, n)
    for i := range p {
        p[i] = i
    }
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            if lcd(nums[i], nums[j]) <= 1 {
                continue
            }
            uni(p, i, j)
        }
    }
    m := make(map[int]int)
    for i := 0; i < n; i++ {
        m[find(p, i)]++
    }
    res := 0
    for _, v := range m {
        res = max(res, v)
    }
    return res
}
func uni(p []int, x, y int) {
    p[find(p, x)] = find(p, y)
}
func find(p []int, x int) int {
    if p[x] != x {
        p[x] = find(p, p[x])
    }
    return p[x]
}
func lcd(x, y int) int {
    if x > y {
        x, y = y, x
    }
    for x > 1 {
        if y % x == 0 {
            break
        }
        x, y = y%x, x
    }
    return x
}
