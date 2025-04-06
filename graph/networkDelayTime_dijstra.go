/*

743. Network Delay Time
Solved
Medium
Topics
Companies
Hint
You are given a network of n nodes, labeled from 1 to n. You are also given times, a list of travel times as directed edges times[i] = (ui, vi, wi), where ui is the source node, vi is the target node, and wi is the time it takes for a signal to travel from source to target.

We will send a signal from a given node k. Return the minimum time it takes for all the n nodes to receive the signal. If it is impossible for all the n nodes to receive the signal, return -1.

 

Example 1:


Input: times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
Output: 2
Example 2:

Input: times = [[1,2,1]], n = 2, k = 1
Output: 1
Example 3:

Input: times = [[1,2,1]], n = 2, k = 2
Output: -1
 

Constraints:

1 <= k <= n <= 100
1 <= times.length <= 6000
times[i].length == 3
1 <= ui, vi <= n
ui != vi
0 <= wi <= 100
All the pairs (ui, vi) are unique. (i.e., no multiple edges.)

*/

func networkDelayTime(times [][]int, n int, k int) int {
    m := make([][]int, n+1)
    for i := range m {
        m[i] = make([]int, n+1)
        for j := range m[i] {
            m[i][j] = 1e8
        }
    }
    for _, tm := range times {
        m[tm[0]][tm[1]] = tm[2]
    }
    dist := make([]int, n+1)
    for i := range dist {
        dist[i] = 1e8
    }
    
    dist[k] = 0
    var q PQ
    heap.Push(&q, []int{0, k})
    for q.Len() > 0 {
        cur := heap.Pop(&q).([]int)
        cost, u := cur[0], cur[1]
        if cost > dist[u] {
            continue
        }
        for v, w := range m[u] {
            if cost + w < dist[v] {
                dist[v] = cost+w
                heap.Push(&q, []int{cost+w, v})
            }
        }
    }
    //fmt.Println("dist=%v", dist)
    res := -1
    for i := 1; i <= n; i++ {
        if dist[i] == 1e8 {
            return -1
        }
        res = max(res, dist[i])
    }
    return res
}

type PQ [][]int
func (q PQ) Len() int{
    return len(q)
}
func (q PQ) Less(i, j int) bool {
    return q[i][0] < q[j][0]
}
func (q PQ) Swap(i, j int) {
    q[i], q[j] = q[j], q[i]
}
func (q *PQ) Push(o any) {
    *q = append(*q, o.([]int))
}
func (q *PQ) Pop() any {
    res := (*q)[len(*q)-1]
    *q = (*q)[:len(*q)-1]
    return res
}


func networkDelayTime_bellman_ford(times [][]int, n int, k int) int {
    tt := make([]int, n)
    for i := range tt {
        tt[i] = math.MaxInt
    }
    tt[k-1] = 0
    for i := 0; i < n-1; i++ {
        for _, t := range times {
            u, v, w := t[0]-1, t[1]-1, t[2]
            if tt[u] == math.MaxInt {
                continue
            }
            tt[v] = min(tt[v], tt[u]+w)
        }
    }

    res := -1
    for _, t := range tt {
        if t == math.MaxInt {
            return -1
        }
        res = max(res, t)
    }
    return res
}
