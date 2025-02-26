/*
473. Matchsticks to Square
Solved
Medium
Topics
Companies
Hint
You are given an integer array matchsticks where matchsticks[i] is the length of the ith matchstick. You want to use all the matchsticks to make one square. You should not break any stick, but you can link them up, and each matchstick must be used exactly one time.

Return true if you can make this square and false otherwise.

 

Example 1:


Input: matchsticks = [1,1,2,2,2]
Output: true
Explanation: You can form a square with length 2, one side of the square came two sticks with length 1.
Example 2:

Input: matchsticks = [3,3,3,3,4]
Output: false
Explanation: You cannot find a way to form a square with all the matchsticks.
 

Constraints:

1 <= matchsticks.length <= 15
1 <= matchsticks[i] <= 108
*/

func makesquare(matchsticks []int) bool {
    sum := 0
    m := matchsticks
    n := len(m)
    for _, l := range m {
        sum += l
    }
    if sum % 4 != 0 {
        return false
    }
    for _, l := range m {
        if l > sum/4 {
            return false
        }
    }

    t := sum/4
    //fmt.Println("t=%v", t)

    sideMasks := make(map[int]bool)
    halfMasks := make(map[int]bool)

    totalMasks := 1<<n
    for i := 0; i < totalMasks; i++ {
        l := 0
        for j := 0; j < 32; j++ {
            if ((i>>j)&1) == 1 {
                l += m[j]
            }
        }
        if l != t {
            continue
        }

        //fmt.Println("side mask=%v", i)
        for mask := range sideMasks {
            if (i & mask) > 0 {
                continue
            }

            halfMask := i | mask
            if halfMasks[totalMasks - 1 - halfMask] {
                return true
            }
            halfMasks[halfMask] = true
        }
        sideMasks[i] = true
    }
    return false
}
func makesquare2(matchsticks []int) bool {
    sum := 0
    m := matchsticks
    n := len(m)
    for _, l := range m {
        sum += l
    }
    if sum % 4 != 0 {
        return false
    }
    for _, l := range m {
        if l > sum/4 {
            return false
        }
    }

    // put ml[0] to sides[0]
    sort.Slice(m, func(i,j int)bool{
        return m[i] > m[j]
    })
    t := sum/4
    sides := make([]int, 4)
    used := make([]bool, n)
    sides[0] += m[0]
    used[0] = true
    k := 0
    if sides[0] == t {
        k++
    }


    return dfs(m, sides, used, t, k)
}

func dfs(ml, sides []int, used []bool, t, k int) bool {
    if k >= 3 {
        return true
    }
    visited := make(map[int]bool)
    for i, l := range ml {
        if used[i] {
            continue
        }
        if sides[k] + l > t {
            continue
        }
        if visited[l] {
            continue
        }

        visited[l] = true
        used[i] = true
        sides[k] += l
        nk := k
        if sides[k] == t {
            nk++
        }
        if dfs(ml, sides, used, t, nk) {
            return true
        }
        used[i] = false
        sides[k] -= l
    }
    return false
}
