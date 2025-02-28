/*
*/
func longestDupSubstring(s string) string {
    n := len(s)
    lo, up := 1, n
    res := ""
    for lo <= up {
        mid := (lo + up) / 2
        dup, found := check(s, mid)
        if found {
            res = dup
            lo = mid+1
        } else {
            up = mid-1
        }
    }
    return res
}

func check(s string, l int) (dup string, found bool) {
    m := make(map[int]int)
    pl := 1
    d, M := 29, (1<<30)+1
    h := 0
    for i := 0; i < l; i++ {
        if i != 0 {
            pl = (pl * d)%M
        }
        h = (h*d%M + int(s[i]-'a'+1))%M
    }
    //fmt.Println("l=%v, pl=%v, h=%v", l, pl, h)
    m[h] = 1
    for i := 1; i+l <= len(s); i++ {
        h = ((h + M - (int(s[i-1]-'a'+1)*pl)%M)*d + int(s[i+l-1]-'a'+1))%M
        //fmt.Println("h=%v", h)
        if m[h] > 0 && s[m[h]-1:m[h]-1+l] == s[i:i+l] {
            return s[i:i+l], true
        }
        m[h] = i+1
    }
    return "", false
}
