/*
1771. Maximize Palindrome Length From Subsequences
Solved
Hard
Topics
Companies
Hint
You are given two strings, word1 and word2. You want to construct a string in the following manner:

Choose some non-empty subsequence subsequence1 from word1.
Choose some non-empty subsequence subsequence2 from word2.
Concatenate the subsequences: subsequence1 + subsequence2, to make the string.
Return the length of the longest palindrome that can be constructed in the described manner. If no palindromes can be constructed, return 0.

A subsequence of a string s is a string that can be made by deleting some (possibly none) characters from s without changing the order of the remaining characters.

A palindrome is a string that reads the same forward as well as backward.

 

Example 1:

Input: word1 = "cacb", word2 = "cbba"
Output: 5
Explanation: Choose "ab" from word1 and "cba" from word2 to make "abcba", which is a palindrome.
Example 2:

Input: word1 = "ab", word2 = "ab"
Output: 3
Explanation: Choose "ab" from word1 and "a" from word2 to make "aba", which is a palindrome.
Example 3:

Input: word1 = "aa", word2 = "bb"
Output: 0
Explanation: You cannot construct a palindrome from the described method, so return 0.
 

Constraints:

1 <= word1.length, word2.length <= 1000
word1 and word2 consist of lowercase English letters.
*/

func longestPalindrome(word1 string, word2 string) int {
    // two steps:
    //. 1, combine word1 and word2 => word => calculate longest palindorm subseq of word for each
    // i, j
    //. 2, starting of each position i of word1, find the corresponding character of word2 j, 
    //   if not found, skip this pair
    //.  otherwise, 2+dp[i+1][n1+j] is a possible candidate
    word := word1 + word2
    n, n1, n2 := len(word), len(word1), len(word2)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    for i := range dp {
        dp[i][i] = 1
    }
    for l := 2; l <= n; l++ {
        for i := 0; i+l-1<n; i++ {
            j := i+l-1
            if word[i] == word[j] {
                dp[i][j] = 2
                if i+1 <= j-1 {
                    dp[i][j] += dp[i+1][j-1]
                }
            } else {
                dp[i][j] = max(dp[i+1][j], dp[i][j-1])
            }
        }
    }

    res := 0
    for i := 0; i < n1; i++ {
        j := n2-1
        for j >= 0 && word2[j] != word1[i] {
            j--
        }
        if j < 0 {
            continue
        }

        res = max(res, 2+dp[i+1][n1+j-1])
    }
    return res
}
