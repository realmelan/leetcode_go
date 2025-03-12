/*
691. Stickers to Spell Word
Solved
Hard
Topics
Companies
Hint
We are given n different types of stickers. Each sticker has a lowercase English word on it.

You would like to spell out the given string target by cutting individual letters from your collection of stickers and rearranging them. You can use each sticker more than once if you want, and you have infinite quantities of each sticker.

Return the minimum number of stickers that you need to spell out target. If the task is impossible, return -1.

Note: In all test cases, all words were chosen randomly from the 1000 most common US English words, and target was chosen as a concatenation of two random words.

 

Example 1:

Input: stickers = ["with","example","science"], target = "thehat"
Output: 3
Explanation:
We can use 2 "with" stickers, and 1 "example" sticker.
After cutting and rearrange the letters of those stickers, we can form the target "thehat".
Also, this is the minimum number of stickers necessary to form the target string.
Example 2:

Input: stickers = ["notice","possible"], target = "basicbasic"
Output: -1
Explanation:
We cannot form the target "basicbasic" from cutting letters from the given stickers.
 

Constraints:

n == stickers.length
1 <= n <= 50
1 <= stickers[i].length <= 10
1 <= target.length <= 15
stickers[i] and target consist of lowercase English letters.

*/

func minStickers(stickers []string, target string) int {
    n := len(stickers)
    tchs := strToChs(target)
    ts := slices.Clone(tchs)
    chs := make([][]int, n)
    for i, s := range stickers {
        chs[i] = strToChs(s)
        for j, v := range chs[i] {
            if v > 0 {
                ts[j]--
            }
        }
    }
    for i, v := range tchs {
        if v > 0 && tchs[i] == ts[i] {
            return -1
        }
    }

    memo := make(map[string]int)
    res := dfs(tchs, chs, memo)
    //fmt.Println("memo=%v", memo)
    return res
}

func dfs(tchs []int, chs [][]int, memo map[string]int) int {
    if done(tchs) {
        return 0
    }
    s := chsToStr(tchs)
    if v, ok := memo[s]; ok {
        return v
    }

    res := math.MaxInt
    for _, ch := range chs {
        // if ch doesn't change tchs, skip
        changed := false
        for i, v := range ch {
            if tchs[i] > 0 && v > 0 {
                changed = true
            }
            tchs[i] -= v
        }
        if changed {
            res = min(res, 1+dfs(tchs, chs, memo))
        }
        for i, v := range ch {
            tchs[i] += v
        }
    }
    memo[s] = res
    return res
}

func done(tchs []int) bool {
    for _, v := range tchs {
        if v > 0 {
            return false
        }
    }
    return true
}

func strToChs(s string)[]int {
    res := make([]int, 26)
    for i := 0; i < len(s); i++ {
        res[int(s[i])-int('a')]++
    }
    return res
}

func chsToStr(chs []int) string {
    var buf strings.Builder
    for i, cnt := range chs {
        for j := 0; j < cnt; j++ {
            buf.WriteByte(byte('a')+byte(i))
        }
    }
    return buf.String()
}
