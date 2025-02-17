/*
2272. Substring With Largest Variance
Solved
Hard
Topics
Companies
Hint
The variance of a string is defined as the largest difference between the number of occurrences of any 2 characters present in the string. Note the two characters may or may not be the same.

Given a string s consisting of lowercase English letters only, return the largest variance possible among all substrings of s.

A substring is a contiguous sequence of characters within a string.

 

Example 1:

Input: s = "aababbb"
Output: 3
Explanation:
All possible variances along with their respective substrings are listed below:
- Variance 0 for substrings "a", "aa", "ab", "abab", "aababb", "ba", "b", "bb", and "bbb".
- Variance 1 for substrings "aab", "aba", "abb", "aabab", "ababb", "aababbb", and "bab".
- Variance 2 for substrings "aaba", "ababbb", "abbb", and "babb".
- Variance 3 for substring "babbb".
Since the largest possible variance is 3, we return it.
Example 2:

Input: s = "abcde"
Output: 0
Explanation:
No letter occurs more than once in s, so the variance of every substring is 0.
 

Constraints:

1 <= s.length <= 104
s consists of lowercase English letters.
*/
func largestVariance(s string) int {
    m := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        m[s[i]]++
    }
    type pair struct {
        char byte
        cnt int
    }
    chars := make([]pair, 0)
    for k, v := range m {
        chars = append(chars, pair{k,v})
    }
    sort.Slice(chars, func(i,j int) bool{
        return chars[i].cnt > chars[j].cnt
    })

    res := 0
    for i := 0; i < len(chars); i++ {
        if chars[i].cnt <= res {
            break
        }
        for j := 0; j < len(chars); j++ {
            if j == i {
                continue
            }
            arr := make([]int, 0)
            cnta := 0
            for k := 0; k <= len(s); k++ {
                if k == len(s) || s[k] == chars[j].char {
                    arr = append(arr, cnta)
                    cnta = 0
                } else if s[k] == chars[i].char {
                    cnta++
                }
            }
            //fmt.Printf("(%v, %v) arr=%v\n", chars[i].char, chars[j].char, arr)

            sum := make([]int, len(arr))
            sum[0] = arr[0]
            for k := 1; k < len(arr); k++ {
                if sum[k-1] >= arr[k-1] {
                    sum[k] = arr[k]-1 + sum[k-1]
                } else {
                    sum[k] = arr[k] -1 + arr[k-1]
                }
                res = max(res, sum[k])
            }
        }
    }
    return res
}
