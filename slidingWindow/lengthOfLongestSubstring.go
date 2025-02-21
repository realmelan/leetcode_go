/*
3. Longest Substring Without Repeating Characters
Solved
Medium
Topics
Companies
Hint
Given a string s, find the length of the longest 
substring
 without duplicate characters.

 

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
 

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/
func lengthOfLongestSubstring(s string) int {
    n := len(s)
    m := make(map[byte]int)
    i, j := 0, 0
    for j < n {
        m[s[j]]++
        if len(m) < j-i+1 {
            m[s[i]]--
            if m[s[i]] == 0 {
                delete(m, s[i])
            }
            i++
        }
        j++
    }
    return j - i
}
