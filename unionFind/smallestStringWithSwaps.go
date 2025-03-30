/*
1202. Smallest String With Swaps
Solved
Medium
Topics
Companies
Hint
You are given a string s, and an array of pairs of indices in the string pairs where pairs[i] = [a, b] indicates 2 indices(0-indexed) of the string.

You can swap the characters at any pair of indices in the given pairs any number of times.

Return the lexicographically smallest string that s can be changed to after using the swaps.

 

Example 1:

Input: s = "dcab", pairs = [[0,3],[1,2]]
Output: "bacd"
Explaination: 
Swap s[0] and s[3], s = "bcad"
Swap s[1] and s[2], s = "bacd"
Example 2:

Input: s = "dcab", pairs = [[0,3],[1,2],[0,2]]
Output: "abcd"
Explaination: 
Swap s[0] and s[3], s = "bcad"
Swap s[0] and s[2], s = "acbd"
Swap s[1] and s[2], s = "abcd"
Example 3:

Input: s = "cba", pairs = [[0,1],[1,2]]
Output: "abc"
Explaination: 
Swap s[0] and s[1], s = "bca"
Swap s[1] and s[2], s = "bac"
Swap s[0] and s[1], s = "abc"
 

Constraints:

1 <= s.length <= 10^5
0 <= pairs.length <= 10^5
0 <= pairs[i][0], pairs[i][1] < s.length
s only contains lower case English letters.
*/
func smallestStringWithSwaps(s string, pairs [][]int) string {
    n := len(s)
    p := make([]int, n)
    for i := range p {
        p[i] = i
    }
    for _, pa := range pairs {
        uni(p, pa[0], pa[1])
    }
    gm := make(map[int][]byte)
    cnt := make(map[int]int)
    for i := 0; i < n; i++ {
        pi := find(p, i)
        gm[pi] = append(gm[pi], s[i])
    }
    for _, g := range gm {
        sort.Slice(g, func(i, j int)bool{
            return g[i] < g[j]
        })
    }

    res := make([]byte, n)
    for i := range s {
        pi := find(p, i)
        res[i] = gm[pi][cnt[pi]]
        cnt[pi]++
    }
    return string(res)
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
