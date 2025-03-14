/*

44. Wildcard Matching
Solved
Hard
Topics
Companies
Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:

'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).
The matching should cover the entire input string (not partial).

 

Example 1:

Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input: s = "aa", p = "*"
Output: true
Explanation: '*' matches any sequence.
Example 3:

Input: s = "cb", p = "?a"
Output: false
Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.
 

Constraints:

0 <= s.length, p.length <= 2000
s contains only lowercase English letters.
p contains only lowercase English letters, '?' or '*'.

*/


func isMatch(s string, p string) bool {
    // use dfs
    sn, pn := len(s), len(p)
    memo := make([][]int, sn)
    for i := range memo {
        memo[i] = make([]int, pn)
    }

    return dfs(s, p, sn-1, pn-1, memo) > 0
}

func dfs(s, p string, si, pi int, memo [][]int) int {
    if si < 0 {
        for i := 0; i <= pi; i++ {
            if p[i] != '*' {
                return -1
            }
        }
        return 1
    } else if pi < 0 {
        return -1
    }
    if memo[si][pi] != 0 {
        return memo[si][pi]
    }

    if p[pi] == '?' {
        memo[si][pi] = dfs(s, p, si-1, pi-1, memo)
        return memo[si][pi]
    } else if p[pi] != '*' {
        if s[si] == p[pi] {
            memo[si][pi] = dfs(s, p, si-1, pi-1, memo)
        } else {
            memo[si][pi] = -1
        }
        return memo[si][pi]
    }

    for i := si; i >= -1; i-- {
        res := dfs(s, p, i, pi-1, memo)
        if res > 0 {
            memo[si][pi] = 1
            return 1
        }
    }

    memo[si][pi] = -1
    return -1
} 
