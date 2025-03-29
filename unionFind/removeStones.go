/*

947. Most Stones Removed with Same Row or Column
Solved
Medium
Topics
Companies
On a 2D plane, we place n stones at some integer coordinate points. Each coordinate point may have at most one stone.

A stone can be removed if it shares either the same row or the same column as another stone that has not been removed.

Given an array stones of length n where stones[i] = [xi, yi] represents the location of the ith stone, return the largest possible number of stones that can be removed.

 

Example 1:

Input: stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
Output: 5
Explanation: One way to remove 5 stones is as follows:
1. Remove stone [2,2] because it shares the same row as [2,1].
2. Remove stone [2,1] because it shares the same column as [0,1].
3. Remove stone [1,2] because it shares the same row as [1,0].
4. Remove stone [1,0] because it shares the same column as [0,0].
5. Remove stone [0,1] because it shares the same row as [0,0].
Stone [0,0] cannot be removed since it does not share a row/column with another stone still on the plane.
Example 2:

Input: stones = [[0,0],[0,2],[1,1],[2,0],[2,2]]
Output: 3
Explanation: One way to make 3 moves is as follows:
1. Remove stone [2,2] because it shares the same row as [2,0].
2. Remove stone [2,0] because it shares the same column as [0,0].
3. Remove stone [0,2] because it shares the same row as [0,0].
Stones [0,0] and [1,1] cannot be removed since they do not share a row/column with another stone still on the plane.
Example 3:

Input: stones = [[0,0]]
Output: 0
Explanation: [0,0] is the only stone on the plane, so you cannot remove it.
 

Constraints:

1 <= stones.length <= 1000
0 <= xi, yi <= 104
No two stones are at the same coordinate point.

*/

func removeStones(stones [][]int) int {
    // after join, a set of k can remove k-1 stones
    n := len(stones)
    p := make([]int, n)
    for i := range p {
        p[i] = i
    }

    rg := make(map[int][]int)
    cg := make(map[int][]int)
    // group stones by row and column
    for i, s := range stones {
        for _, j := range rg[s[0]] {
            uni(p, i, j)
        }
        rg[s[0]] = append(rg[s[0]], i)

        for _, j := range cg[s[1]] {
            uni(p, i, j)
        }
        cg[s[1]] = append(cg[s[1]], i)
    }

    g := make(map[int]int)
    for i := range stones {
        g[find(p, i)]++
    }
    res := 0
    for _, v := range g {
        res += v-1
    }
    return res
}

func uni(p []int, i, j int) {
    p[find(p, i)] = find(p, j)
}
func find(p []int, x int) int {
    if p[x] != x {
        p[x] = find(p, p[x])
    }
    return p[x]
}
