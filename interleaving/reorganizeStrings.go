/*
767. Reorganize String
Solved
Medium
Topics
Companies
Hint
Given a string s, rearrange the characters of s so that any two adjacent characters are not the same.

Return any possible rearrangement of s or return "" if not possible.

 

Example 1:

Input: s = "aab"
Output: "aba"
Example 2:

Input: s = "aaab"
Output: ""
 

Constraints:

1 <= s.length <= 500
s consists of lowercase English letters.
*/
type pair struct {
        char byte
        cnt int
    }
func reorganizeString(s string) string {
    // sort characters in descending order
    // pick two characters with largest counts and interleaving until one character is done
    // then pick the longest from the pool
    m := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        m[s[i]]++
    }
    
    chars := make([]pair, 0)
    for k, v := range m {
        chars = append(chars, pair{k, v})
    }
    sort.Slice(chars, func(i,j int)bool {
        return chars[i].cnt > chars[j].cnt
    })
    if len(chars) == 1 {
        if chars[0].cnt > 1 {
            return ""
        } else {
            return s
        }
    }
    if chars[0].cnt > (len(s)+1)/2 {
        return ""
    }

    return compose(chars)
    
}

func compose(chars []pair) string {
    // find the two most frequent characters
    sort.Slice(chars, func(i,j int)bool {
        return chars[i].cnt > chars[j].cnt
    })
    if chars[0].cnt == 0 {
        return ""
    }
    if chars[1].cnt == 0 {
        return string(chars[0].char)
    }

    a, b := chars[0].char, chars[1].char
    chars[0].cnt--
    chars[1].cnt--

    s := compose(chars)
    if s == "" {
        return string(a)+string(b)
    } else if s[0] == a {
        return string(a) + string(b) + s
    } else {
        return string(b) + string(a) + s
    }
}
