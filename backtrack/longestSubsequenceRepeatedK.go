/*
*/
func longestSubsequenceRepeatedK(s string, k int) string {
    // 1, count character appearances
    // 2, count upper bound
    // 3, use binary search to determine the longest
    //
    // [how to test candidate length] 
    // for each candidate length, compose subsequence in reverse lexicographical order
    // and test whether it repeats for k times using two pointers.
    //
    // [how to compose subsequence]
    // 1, componse the largest seq,
    // 2, use current seq, generate next smaller seq
    m := make(map[byte]int)
    for i := 0;i < len(s); i++ {
        m[s[i]]++
    }
    upper := 0
    for c := range m {
        m[c] /= k
        upper += m[c]
    }
    lower := 1
    res := ""
    for lower <= upper {
        mid := (lower + upper) / 2
        //fmt.Println("mid=%v, l=%v, r=%v", mid, lower, upper)
        cand := check(s, m, mid, k)
        if cand != "" {
            res = cand
            lower = mid+1
        } else {
            upper = mid-1
        }
        //fmt.Println("mid=%v, l=%v, r=%v, res=%v", mid, lower, upper, res)
    }
    return res
}
 
func check(s string, m map[byte]int, size, k int) string{
    // build largest seq of size
    var chars []byte
    for c, v := range m {
        for i := 0; i < v; i++ {
            chars = append(chars, c)
        }
    }
    sort.Slice(chars, func(i, j int)bool {
        return chars[i] > chars[j]
    })

    //fmt.Println("chars=%v", string(chars))
    used := make([]bool, len(chars))
    return dfs(s, nil, chars, used, size, k)
}

// return next smaller candidate
func dfs(s string, seq, chars []byte, used []bool, size, k int) string { 
    if len(seq) == size {
        if match(s, seq, k) {
            return string(seq)
        }
    }

    visited := make(map[byte]bool)
    for i := 0; i < len(chars); i++ {
        if used[i] {
            continue
        }
        if visited[chars[i]] {
            continue
        }

        visited[chars[i]] = true
        used[i] = true
        seq = append(seq, chars[i])
        res := dfs(s, seq, chars, used, size, k)
        if res != "" {
            return res
        }
        seq = seq[0:len(seq)-1]
        used[i] = false
    }
    return ""
}

func match(s string, seq []byte, k int) bool {
    ns, nq := len(s), len(seq)
    i, j := 0, 0
    for i < ns && j < nq*k {
        if s[i] == seq[j%nq] {
            j++
        }
        i++
    }
    return j >= nq*k
}
