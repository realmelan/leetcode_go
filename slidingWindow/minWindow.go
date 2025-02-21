
/*
76. Minimum Window Substring
Solved
Hard
Topics
Companies
Hint
Given two strings s and t of lengths m and n respectively, return the minimum window 
substring
 of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.

 

Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.
 

Constraints:

m == s.length
n == t.length
1 <= m, n <= 105
s and t consist of uppercase and lowercase English letters.
*/
func minWindow(s string, t string) string {
    mt := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        mt[t[i]]++
    }
    i, j, n := 0, 0, len(s)
    minStart, res := 0, n+1
    for j < n {
        ch := s[j]
        j++
        if _, ok := mt[ch]; !ok {
            continue
        }

        mt[ch]--
        for i < j && condition(mt) {
            //fmt.Println("j=%v, i=%v, mt=%v", j, i, mt)
            if res > j-i {
                res = j-i
                minStart = i
            }
            if _, ok := mt[s[i]]; ok {
                mt[s[i]]++
            }
            i++
        }
    }

    if res == n+1 {
        return ""
    }
    return s[minStart:minStart+res]
}

func condition(m map[byte]int) bool {
    for _, v := range m {
        if v > 0 {
            return false
        }
    }
    return true
}
