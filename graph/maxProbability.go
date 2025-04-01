/*
1514. Path with Maximum Probability
Solved
Medium
Topics
Companies
Hint
You are given an undirected weighted graph of n nodes (0-indexed), represented by an edge list where edges[i] = [a, b] is an undirected edge connecting the nodes a and b with a probability of success of traversing that edge succProb[i].

Given two nodes start and end, find the path with the maximum probability of success to go from start to end and return its success probability.

If there is no path from start to end, return 0. Your answer will be accepted if it differs from the correct answer by at most 1e-5.

 

Example 1:



Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.2], start = 0, end = 2
Output: 0.25000
Explanation: There are two paths from start to end, one having a probability of success = 0.2 and the other has 0.5 * 0.5 = 0.25.
Example 2:



Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.3], start = 0, end = 2
Output: 0.30000
Example 3:



Input: n = 3, edges = [[0,1]], succProb = [0.5], start = 0, end = 2
Output: 0.00000
Explanation: There is no path between 0 and 2.
 

Constraints:

2 <= n <= 10^4
0 <= start, end < n
start != end
0 <= a, b < n
a != b
0 <= succProb.length == edges.length <= 2*10^4
0 <= succProb[i] <= 1
There is at most one edge between every two nodes.
*/
func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
    cons := make([]map[int]float64, n)
    for i := range cons {
        cons[i] = make(map[int]float64)
    }
    for i, e := range edges {
        cons[e[0]][e[1]] = succProb[i]
        cons[e[1]][e[0]] = succProb[i]
    }

    seen := make([]float64, n)
    seen[start_node] = 1

    var q PQ
    heap.Push(&q, Item{start_node, float64(1)})
    for q.Len() > 0 {
        item := heap.Pop(&q).(Item)
        if item.node == end_node {
            //fmt.Println("item.prob=%v", item.prob)
            return item.prob
        }
        
        for nd, prob := range cons[item.node] {
            if item.prob * prob <= seen[nd] {
                continue
            }
            seen[nd] = item.prob * prob
            heap.Push(&q, Item{nd, seen[nd]})
        }
    }
    return 0
    
}

type Item struct {
    node int
    prob float64
}

type PQ []Item
func (q PQ) Len() int{
    return len(q)
}
func (q PQ) Less(i, j int) bool {
    return q[i].prob > q[j].prob
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
