/*

1851. Minimum Interval to Include Each Query
Solved
Hard
Topics
Companies
Hint
You are given a 2D integer array intervals, where intervals[i] = [lefti, righti] describes the ith interval starting at lefti and ending at righti (inclusive). The size of an interval is defined as the number of integers it contains, or more formally righti - lefti + 1.

You are also given an integer array queries. The answer to the jth query is the size of the smallest interval i such that lefti <= queries[j] <= righti. If no such interval exists, the answer is -1.

Return an array containing the answers to the queries.

 

Example 1:

Input: intervals = [[1,4],[2,4],[3,6],[4,4]], queries = [2,3,4,5]
Output: [3,3,1,4]
Explanation: The queries are processed as follows:
- Query = 2: The interval [2,4] is the smallest interval containing 2. The answer is 4 - 2 + 1 = 3.
- Query = 3: The interval [2,4] is the smallest interval containing 3. The answer is 4 - 2 + 1 = 3.
- Query = 4: The interval [4,4] is the smallest interval containing 4. The answer is 4 - 4 + 1 = 1.
- Query = 5: The interval [3,6] is the smallest interval containing 5. The answer is 6 - 3 + 1 = 4.
Example 2:

Input: intervals = [[2,3],[2,5],[1,8],[20,25]], queries = [2,19,5,22]
Output: [2,-1,4,6]
Explanation: The queries are processed as follows:
- Query = 2: The interval [2,3] is the smallest interval containing 2. The answer is 3 - 2 + 1 = 2.
- Query = 19: None of the intervals contain 19. The answer is -1.
- Query = 5: The interval [2,5] is the smallest interval containing 5. The answer is 5 - 2 + 1 = 4.
- Query = 22: The interval [20,25] is the smallest interval containing 22. The answer is 25 - 20 + 1 = 6.
 

Constraints:

1 <= intervals.length <= 105
1 <= queries.length <= 105
intervals[i].length == 2
1 <= lefti <= righti <= 107
1 <= queries[j] <= 107
*/

func minInterval(intervals [][]int, queries []int) []int {
    mqi := make(map[int][]int)
    res := make([]int, len(queries))
    for i, q := range queries {
        mqi[q] = append(mqi[q], i)
        res[i] = -1
    }
    sort.Ints(queries)
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    ans := make(map[int]int)
    que := make(PQ, 0)
    heap.Init(&que)
    x, xn := 0, len(intervals)
    for _, q := range queries {
        for x < xn && intervals[x][0] <= q {
            // add to priority queue
            heap.Push(&que, Item{intervals[x][1]-intervals[x][0]+1, intervals[x][1]+1})
            x++
        }
        //fmt.Println("q=%q, interval=%v, que=%v", q, intervals[x], que)

        // find open smallest interval using heap
        for que.Len() > 0 && que[0].closed <= q {
            _ = heap.Pop(&que)
        }
        if que.Len() > 0 {
            ans[q] = que[0].size
        }
        //fmt.Println("ans=%v", ans)
    }
    for q, a := range ans {
        for _, idx := range mqi[q] {
            res[idx] = a
        }
    }
    return res
}

type Item struct {
    size int
    closed int
}

type PQ []Item

func (p *PQ) Len() int {
    return len(*p)
}
func (p *PQ) Less(i, j int) bool {
    return (*p)[i].size < (*p)[j].size
} 
func (p *PQ) Swap(i, j int) {
    (*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}
func (p *PQ) Push(x any) {
    item := x.(Item)
    *p = append(*p, item)
}
func (p *PQ) Pop() any {
    if len(*p) == 0 {
        panic("empty queue")
    }
    x := (*p)[len(*p)-1]
    *p = (*p)[0:len(*p)-1]
    return x
}

