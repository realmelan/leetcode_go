/*

1349. Maximum Students Taking Exam
Solved
Hard
Topics
Companies
Hint
Given a m * n matrix seats  that represent seats distributions in a classroom. If a seat is broken, it is denoted by '#' character otherwise it is denoted by a '.' character.

Students can see the answers of those sitting next to the left, right, upper left and upper right, but he cannot see the answers of the student sitting directly in front or behind him. Return the maximum number of students that can take the exam together without any cheating being possible.

Students must be placed in seats in good condition.

 

Example 1:


Input: seats = [["#",".","#","#",".","#"],
                [".","#","#","#","#","."],
                ["#",".","#","#",".","#"]]
Output: 4
Explanation: Teacher can place 4 students in available seats so they don't cheat on the exam. 
Example 2:

Input: seats = [[".","#"],
                ["#","#"],
                ["#","."],
                ["#","#"],
                [".","#"]]
Output: 3
Explanation: Place all students in available seats. 

Example 3:

Input: seats = [["#",".",".",".","#"],
                [".","#",".","#","."],
                [".",".","#",".","."],
                [".","#",".","#","."],
                ["#",".",".",".","#"]]
Output: 10
Explanation: Place students in available seats in column 1, 3 and 5.
 

Constraints:

seats contains only characters '.' and'#'.
m == seats.length
n == seats[i].length
1 <= m <= 8
1 <= n <= 8

*/

func maxStudents(seats [][]byte) int {
    // state compression dp, rows by rows
    // let dp[i][mask] = max students for rows[0:i] with row[i]'s bit mask = mask
    m, n := len(seats), len(seats[0])
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, 1<<n)
    }

    res := 0
    for i := 0; i < m; i++ { 
        for s := 0; s < 1<<n; s++ {
            // check whther s is a valid bit mask for row i or not
            // 1: students not assigned a broken chair
            // 2: no two students are next to each other
            if (s & (s<<1)) > 0 {
                dp[i+1][s] = -1
                continue
            }
            cnt := 0
            for j := 0; j < n; j++ {
                if (s >> j) & 1 > 0 {
                    if seats[i][j] == '#' {
                        dp[i+1][s] = -1
                        break
                    }
                    cnt++
                }
            }
            if dp[i+1][s] == -1 {
                continue
            }
            // fmt.Println("check row %v, s %v", i, s)
            // if valid, loop through row[i-1]'s bit mask,
            for mask, val := range dp[i] {
                if val < 0 { // invalid bit mask like student is assigned a broken chair
                    continue
                }
                // cheat is possible
                if ((mask >> 1) & s) > 0 || ((mask << 1) & s) > 0 {
                    continue
                }
                dp[i+1][s] = max(dp[i+1][s], val+cnt)
            }

            if i == m-1 {
                res = max(res, dp[i+1][s])
            }
        }
    }
    return res
}
