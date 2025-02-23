/*
1871. Jump Game VII
Solved
Medium
Topics
Companies
Hint
You are given a 0-indexed binary string s and two integers minJump and maxJump. In the beginning, you are standing at index 0, which is equal to '0'. You can move from index i to index j if the following conditions are fulfilled:

i + minJump <= j <= min(i + maxJump, s.length - 1), and
s[j] == '0'.
Return true if you can reach index s.length - 1 in s, or false otherwise.

 

Example 1:

Input: s = "011010", minJump = 2, maxJump = 3
Output: true
Explanation:
In the first step, move from index 0 to index 3. 
In the second step, move from index 3 to index 5.
Example 2:

Input: s = "01101110", minJump = 2, maxJump = 3
Output: false
 

Constraints:

2 <= s.length <= 105
s[i] is either '0' or '1'.
s[0] == '0'
1 <= minJump <= maxJump < s.length

*/
func canReach(s string, minJump int, maxJump int) bool {
    // from left to right, test each position of s to see whether it can be reached or not
    // a position i can be reached from any j in [i-maxJump, i-minJump]
    // if j is already reached, then i can also be reached.
    n := len(s)
    if s[n-1] == '1' {
        return false
    }
    m := make(map[int]bool)
    m[0] = true
    right := 0
    for i := 1; i < n; i++ {
        if i > right + maxJump {
            return false
        }
        if s[i] == '1' {
            continue
        }

        for j := i - minJump; j >= max(i - maxJump, 0); j-- {
            if s[j] == '1' {
                continue
            }
            if m[j] {
                m[i] = true
                right = i
                break
            }
        }
    }
    return m[n-1]
}
