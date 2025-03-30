/*
1489. Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree
Solved
Hard
Topics
Companies
Hint
Given a weighted undirected connected graph with n vertices numbered from 0 to n - 1, and an array edges where edges[i] = [ai, bi, weighti] represents a bidirectional and weighted edge between nodes ai and bi. A minimum spanning tree (MST) is a subset of the graph's edges that connects all vertices without cycles and with the minimum possible total edge weight.

Find all the critical and pseudo-critical edges in the given graph's minimum spanning tree (MST). An MST edge whose deletion from the graph would cause the MST weight to increase is called a critical edge. On the other hand, a pseudo-critical edge is that which can appear in some MSTs but not all.

Note that you can return the indices of the edges in any order.

 

Example 1:



Input: n = 5, edges = [[0,1,1],[1,2,1],[2,3,2],[0,3,2],[0,4,3],[3,4,3],[1,4,6]]
Output: [[0,1],[2,3,4,5]]
Explanation: The figure above describes the graph.
The following figure shows all the possible MSTs:

Notice that the two edges 0 and 1 appear in all MSTs, therefore they are critical edges, so we return them in the first list of the output.
The edges 2, 3, 4, and 5 are only part of some MSTs, therefore they are considered pseudo-critical edges. We add them to the second list of the output.
Example 2:



Input: n = 4, edges = [[0,1,1],[1,2,1],[2,3,1],[0,3,1]]
Output: [[],[0,1,2,3]]
Explanation: We can observe that since all 4 edges have equal weight, choosing any 3 edges from the given 4 will yield an MST. Therefore all 4 edges are pseudo-critical.
 

Constraints:

2 <= n <= 100
1 <= edges.length <= min(200, n * (n - 1) / 2)
edges[i].length == 3
0 <= ai < bi < n
1 <= weighti <= 1000
All pairs (ai, bi) are distinct.
*/

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	// step 1: build a MST and get the weight
	// step 2: for each edge, remove from the list and build a new MST
	//.   1) if new weight is larger or can't build an MST, then edge is critical
	//    2) if new weight is the same, then edge isn't
	// part 2, test pseudo critical edge:
	//.   1) for all edges in step 1 MST, add them to mst edges
	//.   2) for edges not in step 1, starts with the edge, use Krushkal to build a MST
	//       if new weight is larger, then this edge should be excluded from MST
	//       else add edge to mst edges
	en := len(edges)
	es := make([][]int, 0)
	for i, e := range edges {
		es = append(es, []int{e[0], e[1], e[2], i})
	}
	sort.Slice(es, func(i, j int) bool {
		return es[i][2] < es[j][2]
	})

	// part 1: find critical edges
	// step 1: build MST
	used := make([]bool, en)
	target := mst(n, 0, -1, es, used)
    //fmt.Println("mst=%v", target)
	// step 2: skip e,
    cr := make([]int, 0)
	for i, e := range es {
		used := make([]bool, en)
        used[i] = true
		weight := mst(n, 0, -1, es, used)
		if weight != target {
			cr = append(cr, e[3])
		}
	}

    // part 2:
    pc := make(map[int]bool)
    for i, e := range es {
        used := make([]bool, en)
        used[i] = true
        weight := mst(n, e[2], i, es, used)
        if weight == target {
            pc[e[3]] = true
        }
    }

    var res2 []int
    for _, k := range cr {
        delete(pc, k)
    }
    for k := range pc {
        res2 = append(res2, k)
    }
    return [][]int{cr, res2}
}

func mst(n, w, e int, es [][]int, used []bool) int {
	// use krushkal
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
    res := w
    if e >= 0 {
        // connect
        p[es[e][0]] = es[e][1]
    }

	for i, e := range es {
		if used[i] {
			continue
		}
		p1 := find(p, e[0])
		p2 := find(p, e[1])
		if p1 == p2 {
			continue
		}
		used[i] = true
		res += e[2]
		p[p1] = p2
	}
	return res
}

func uni(p []int, i, j int) {
	p[find(p, i)] = find(p, j)
}
func find(p []int, x int) int {
	if p[x] != x {
		p[x] = find(p, p[x])
	}
	return p[x]
}
