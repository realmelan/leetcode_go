/*
1584. Min Cost to Connect All Points
Solved
Medium
Topics
Companies
Hint
You are given an array points representing integer coordinates of some points on a 2D-plane, where points[i] = [xi, yi].

The cost of connecting two points [xi, yi] and [xj, yj] is the manhattan distance between them: |xi - xj| + |yi - yj|, where |val| denotes the absolute value of val.

Return the minimum cost to make all points connected. All points are connected if there is exactly one simple path between any two points.

 

Example 1:


Input: points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
Output: 20
Explanation: 

We can connect the points as shown above to get the minimum cost of 20.
Notice that there is a unique path between every pair of points.
Example 2:

Input: points = [[3,12],[-2,5],[-4,1]]
Output: 18
 

Constraints:

1 <= points.length <= 1000
-106 <= xi, yi <= 106
All pairs (xi, yi) are distinct.
*/


func minCostConnectPoints(points [][]int) int {
    // Prim's using a priority queue
    // better than Kruskal as connections are dense(every two nodes can connect)
    n := len(points)
    added := make(map[int]bool)
    added[0] = true
    var q PQ
    cur, res := 0, 0
    for len(added) < n {
        for i := 0; i < n; i++ {
            if added[i] {
                continue
            }
            heap.Push(&q, []int{dist(points, cur, i), i})
        }

        for q.Len() > 0 {
            next := heap.Pop(&q).([]int)
            if added[next[1]] {
                continue
            }
            res += next[0]
            cur = next[1]
            added[cur]=true
            break
        }
        
        //fmt.Println("next=%v, res=%v", next, res)
        
    }
    return res
}

func dist(pts [][]int, i, j int) int {
    xi, xj := pts[i][0], pts[j][0]
    yi, yj := pts[i][1], pts[j][1]
    if xi > xj {
        xi, xj = xj, xi
    }
    if yi > yj {
        yi, yj = yj, yi
    }
    return xj - xi + yj - yi
}

type PQ [][]int
func (p PQ) Len() int {
    return len(p)
}
func (p PQ) Less(i, j int) bool {
    return p[i][0] < p[j][0]
}
func (p PQ) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}
func (p *PQ) Push(o any) {
    *p = append(*p, o.([]int))
}
func (p *PQ) Pop() any {
    res := (*p)[len(*p)-1]
    *p = (*p)[:len(*p)-1]
    return res
}
