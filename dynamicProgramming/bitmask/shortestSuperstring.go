/*

943. Find the Shortest Superstring
Solved
Hard
Topics
Companies
Given an array of strings words, return the smallest string that contains each string in words as a substring. If there are multiple valid strings of the smallest length, return any of them.

You may assume that no string in words is a substring of another string in words.

 

Example 1:

Input: words = ["alex","loves","leetcode"]
Output: "alexlovesleetcode"
Explanation: All permutations of "alex","loves","leetcode" would also be accepted.
Example 2:

Input: words = ["catg","ctaagt","gcta","ttca","atgcatc"]
Output: "gctaagttcatgcatc"
 

Constraints:

1 <= words.length <= 12
1 <= words[i].length <= 20
words[i] consists of lowercase English letters.
All the strings of words are unique.

*/

func shortestSuperstring(words []string) string {
    // state compression dp
    // let dp[s][w] = smallest string for words x in state s and ends with word w
    // based on dp[s][w], append another word y to get a candidate of the new state dp[s|y][y]
    n := len(words)
    dp := make([][]string, 1<<n)
    for i := range dp {
        dp[i] = make([]string, n)
    }
    for i, w := range words {
        dp[1<<i][i] = w
    }
    for s := 1; s < 1<<n; s++ {
        for _, sw := range dp[s] {
            if len(sw) == 0 {
                continue
            }
            for j, w := range words {
                if (s >> j) & 1 > 0 {
                    continue
                }
                ns := s | 1<<j
                nsw := combine(sw, w)
                if len(dp[ns][j]) == 0 || len(dp[ns][j]) > len(nsw) {
                    dp[ns][j] = nsw
                }
            }
        }
    }

    var res string
    for _, w := range dp[(1<<n)-1] {
        if len(res) == 0 || len(res) > len(w) {
            res = w
        }
    }
    return res
}

func combine(s, w string) string {
    ns, nw := len(s), len(w)
    
    for i := 0; i < ns; i++ {
        k, j := i, 0
        for k < ns && j < nw {
            if s[k] != w[j] {
                break
            }
            k++
            j++
        }
        if k == ns {
            return s + w[j:]
        }
    }
    return s + w
}
