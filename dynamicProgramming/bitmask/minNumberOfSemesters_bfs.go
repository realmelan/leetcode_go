/*

1494. Parallel Courses II
Solved
Hard
Topics
Companies
Hint
You are given an integer n, which indicates that there are n courses labeled from 1 to n. You are also given an array relations where relations[i] = [prevCoursei, nextCoursei], representing a prerequisite relationship between course prevCoursei and course nextCoursei: course prevCoursei has to be taken before course nextCoursei. Also, you are given the integer k.

In one semester, you can take at most k courses as long as you have taken all the prerequisites in the previous semesters for the courses you are taking.

Return the minimum number of semesters needed to take all courses. The testcases will be generated such that it is possible to take every course.

 

Example 1:


Input: n = 4, relations = [[2,1],[3,1],[1,4]], k = 2
Output: 3
Explanation: The figure above represents the given graph.
In the first semester, you can take courses 2 and 3.
In the second semester, you can take course 1.
In the third semester, you can take course 4.
Example 2:


Input: n = 5, relations = [[2,1],[3,1],[4,1],[1,5]], k = 2
Output: 4
Explanation: The figure above represents the given graph.
In the first semester, you can only take courses 2 and 3 since you cannot take more than two per semester.
In the second semester, you can take course 4.
In the third semester, you can take course 1.
In the fourth semester, you can take course 5.
 

Constraints:

1 <= n <= 15
1 <= k <= n
0 <= relations.length <= n * (n-1) / 2
relations[i].length == 2
1 <= prevCoursei, nextCoursei <= n
prevCoursei != nextCoursei
All the pairs [prevCoursei, nextCoursei] are unique.
The given graph is a directed acyclic graph.

*/

func minNumberOfSemesters(n int, relations [][]int, k int) int {
    // bit mask
    // let dp[s] = # of semesters needed to takes courses in s
    // based on s, find out all courses, A. that can be taken
    // enumerate subset s0 of size k of A, and update dp[s|s0] if semester is smaller
    // one way to speed prerequsite check: use prerequisite mask
    // use BFS + bitmask
    target := (1<<n) - 1
    pmask := make([]int, n)
    for _, r := range relations {
        pmask[r[1]-1] |= 1<<(r[0]-1)
    }

    dp := make([]int, 1<<n)
    for i := range dp {
        dp[i] = math.MaxInt
    }
    dp[0] = 0
    q := make([]int, 1)
    q[0] = 0
    i, ses := 0, 0
    for len(q) > i {
        ses++
        j := len(q)
        for ; i < j; i++ {
            s := q[i]
            nc := 0
            for j, pm := range pmask {
                if (pm & s) == pm && ((s >> j) & 1) == 0 {
                    nc |= 1<<j 
                }
            }
            
            // loop through subset of nc
            for ns := nc; ns >= 1; ns = (ns-1)&nc {
                cnt := 0
                for j := 0; j < n; j++ {
                    if (ns >> j) & 1 > 0 {
                        cnt++
                    }
                }
                if cnt > k {
                    continue
                }
                if (s|ns) == target {
                    return ses
                }
                if dp[s|ns] > ses {
                    dp[s|ns] = ses
                    q = append(q, s|ns)
                }
            }
        }
    }
    return 0
}
