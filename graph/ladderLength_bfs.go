/*
127. Word Ladder
Solved
Hard
Topics
Companies
A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:

Every adjacent pair of words differs by a single letter.
Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
sk == endWord
Given two words, beginWord and endWord, and a dictionary wordList, return the number of words in the shortest transformation sequence from beginWord to endWord, or 0 if no such sequence exists.

 

Example 1:

Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.
Example 2:

Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: 0
Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation sequence.
 

Constraints:

1 <= beginWord.length <= 10
endWord.length == beginWord.length
1 <= wordList.length <= 5000
wordList[i].length == beginWord.length
beginWord, endWord, and wordList[i] consist of lowercase English letters.
beginWord != endWord
All the words in wordList are unique.
*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
    return ladderLength_two_end_BFS(beginWord, endWord, wordList)
}
func ladderLength_two_end_BFS(beginWord string, endWord string, wordList []string) int {
    n := len(wordList)
    cons := make([][]bool, n+1)
    for i := range cons {
        cons[i] = make([]bool, n+1)
    }
    for i, w := range wordList {
        if isConnected(beginWord, w) {
            cons[0][i+1] = true
        }
    }
    src, dst := 0, 0
    for i := 0; i < n; i++ {
        if wordList[i] == endWord {
            dst = i+1
        }
        for j := i+1; j < n; j++ {
            if isConnected(wordList[i], wordList[j]) {
                cons[i+1][j+1] = true
                cons[j+1][i+1] = true
            }
        }
    }
    //fmt.Println(cons)
    if dst == 0 {
        return 0
    }

    srcSet := make([]bool, n+1)
    srcSet[src] = true
    srcq := []int{src}
    dstSet := make([]bool, n+1)
    dstq := []int{dst}
    dstSet[dst] = true
    
    is, id, steps := 0, 0, 1
    for is < len(srcq) || id < len(dstq) {
        steps++
        j := len(srcq)
        for ; is < j; is++ {
            for w, con := range cons[srcq[is]] {
                if !con || srcSet[w] {
                    continue
                }
                if dstSet[w] {
                    return steps
                }
                srcSet[w]=true
                srcq = append(srcq, w)
            }
        }
        steps++
        j = len(dstq)
        for ; id < j; id++ {
            for w, con := range cons[dstq[id]] {
                if !con || dstSet[w] {
                    continue
                }
                if srcSet[w] {
                    return steps
                }
                dstSet[w]=true
                dstq = append(dstq, w)
            }
        }
    }
    return 0
}
func ladderLength_BFS(beginWord string, endWord string, wordList []string) int {
    n := len(wordList)
    cons := make([][]bool, n+1)
    for i := range cons {
        cons[i] = make([]bool, n+1)
    }
    for i, w := range wordList {
        if isConnected(beginWord, w) {
            cons[0][i+1] = true
        }
    }
    src, dst := 0, 0
    for i := 0; i < n; i++ {
        if wordList[i] == endWord {
            dst = i+1
        }
        for j := i+1; j < n; j++ {
            if isConnected(wordList[i], wordList[j]) {
                cons[i+1][j+1] = true
                cons[j+1][i+1] = true
            }
        }
    }
    //fmt.Println(cons)
    if dst == 0 {
        return 0
    }

    visited := make([]bool, n+1)
    q := []int{src}
    i, steps := 0, 1
    for i < len(q) {
        steps++
        j := len(q)
        for ; i < j; i++ {
            for w, con := range cons[q[i]] {
                if !con || visited[w] {
                    continue
                }
                if w == dst {
                    return steps
                }
                visited[w]=true
                q = append(q, w)
            }
        }
    }
    return 0
}
func ladderLength_SPF(beginWord string, endWord string, wordList []string) int {
    n := len(wordList)
    cons := make([][]bool, n+1)
    for i := range cons {
        cons[i] = make([]bool, n+1)
    }
    for i, w := range wordList {
        if isConnected(beginWord, w) {
            cons[0][i+1] = true
        }
    }
    src, dst := 0, -1
    for i := 0; i < n; i++ {
        if wordList[i] == endWord {
            dst = i+1
        }
        for j := i+1; j < n; j++ {
            if isConnected(wordList[i], wordList[j]) {
                cons[i+1][j+1] = true
                cons[j+1][i+1] = true
            }
        }
    }
    //fmt.Println(cons)
    if dst == -1 {
        return 0
    }

    // now use dijstra to find shortest path from 0 to end word
    steps := make([]int, n+1)
    for i := range steps {
        steps[i] = n+2
    }
    steps[0] = 1

    var q PQ
    heap.Push(&q, Item{src, 1})
    for q.Len() > 0 {
        item := heap.Pop(&q).(Item)
        if item.w == dst {
            return item.steps
        }
        if item.steps > steps[item.w] {
            continue
        }
        for w, con := range cons[item.w] {
            if !con {
                continue
            }
            if item.steps+1 < steps[w] {
                steps[w] = item.steps+1
                heap.Push(&q, Item{w, steps[w]})
            }
        }
    }
    return 0
}

func isConnected(w1, w2 string) bool {
    diff := 0
    for i := 0; i < len(w1); i++ {
        if w1[i] != w2[i] {
            diff++
        }
    }
    return diff == 1
}

type Item struct {
    w int
    steps int
}

type PQ []Item
func (q PQ) Len() int{
    return len(q)
}
func (q PQ) Less(i, j int) bool {
    return q[i].steps < q[j].steps
}
func (q PQ) Swap(i, j int) {
    q[i], q[j] = q[j], q[i]
}
func (q *PQ) Push(o any) {
    *q = append(*q, o.(Item))
}
func (q *PQ) Pop() any {
    res := (*q)[len(*q)-1]
    *q = (*q)[:len(*q)-1]
    return res
}
