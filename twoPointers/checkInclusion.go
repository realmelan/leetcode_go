/*
567. Permutation in String
Solved
Medium
Topics
Companies
Hint
Given two strings s1 and s2, return true if s2 contains a 
permutation
 of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.

 

Example 1:

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input: s1 = "ab", s2 = "eidboaoo"
Output: false
 

Constraints:

1 <= s1.length, s2.length <= 104
s1 and s2 consist of lowercase English letters.
*/
func checkInclusion(s1 string, s2 string) bool {
    m := make(map[byte]int)
    n1, n2 := len(s1), len(s2)
    for i := 0; i < n1; i++ {
        m[s1[i]]++
    }

    i := 0
    for i < n2 {
        m[s2[i]]--

        if i >= n1 {
            m[s2[i-n1]]++
        }
        if i >= n1-1 && found(m) {
            return true
        }
        i++
    }

    return false
}

func found(m map[byte]int) bool {
    for _, v := range m {
        if v != 0 {
            return false
        }
    }
    return true
}
