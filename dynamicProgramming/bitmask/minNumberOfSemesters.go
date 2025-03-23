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
    count := make([]int, n+1)
    count[0] = 1
    update := make([][]int, n+1)
    for _, r := range relations {
        count[r[1]]++
        update[r[0]] = append(update[r[0]], r[1])
    }
    dp := make([]int, 1<<n)
    for i := range dp {
        dp[i] = math.MaxInt
    }
    dp[0] = 0
    for s := 0; s < 1<<n; s++ {
        if dp[s] == math.MaxInt {
            continue
        }
        cc := slices.Clone(count)
        sSet := make(map[int]bool)
        for i := 0; i < n; i++ {
            if (s>>i) & 1 > 0 {
                sSet[i+1] = true
                for _, next := range update[i+1] {
                    cc[next]--
                }
            }
        }
        var nc []int
        for i, c := range cc {
            if !sSet[i] && c == 0 {
                nc = append(nc, i)
            }
        }

        if len(nc) <= k {
            ns := s
            for _, c := range nc {
                ns |= (1<<(c-1))
            }
            dp[ns] = min(dp[ns], dp[s] + 1)
            continue
        }

        // > k course, need to loop through 
        nn := len(nc)
        for ncs := (1<<k)-1; ncs < 1<<nn; ncs++ {
            cnt := 0
            ns := s
            for i := 0; i < nn; i++ {
                if ncs >> i & 1 > 0 {
                    cnt++
                    ns |= (1<<(nc[i]-1))
                }
            }
            if cnt == k {
                dp[ns] = min(dp[ns], dp[s]+1)
            }
        }
    }
    return dp[(1<<n)-1]
}
