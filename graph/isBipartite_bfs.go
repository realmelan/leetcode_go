/*

785. Is Graph Bipartite?
Solved
Medium
Topics
Companies
There is an undirected graph with n nodes, where each node is numbered between 0 and n - 1. You are given a 2D array graph, where graph[u] is an array of nodes that node u is adjacent to. More formally, for each v in graph[u], there is an undirected edge between node u and node v. The graph has the following properties:

There are no self-edges (graph[u] does not contain u).
There are no parallel edges (graph[u] does not contain duplicate values).
If v is in graph[u], then u is in graph[v] (the graph is undirected).
The graph may not be connected, meaning there may be two nodes u and v such that there is no path between them.
A graph is bipartite if the nodes can be partitioned into two independent sets A and B such that every edge in the graph connects a node in set A and a node in set B.

Return true if and only if it is bipartite.

 

Example 1:


Input: graph = [[1,2,3],[0,2],[0,1,3],[0,2]]
Output: false
Explanation: There is no way to partition the nodes into two independent sets such that every edge connects a node in one and a node in the other.
Example 2:


Input: graph = [[1,3],[0,2],[1,3],[0,2]]
Output: true
Explanation: We can partition the nodes into two sets: {0, 2} and {1, 3}.
 

Constraints:

graph.length == n
1 <= n <= 100
0 <= graph[u].length < n
0 <= graph[u][i] <= n - 1
graph[u] does not contain u.
All the values of graph[u] are unique.
If graph[u] contains v, then graph[v] contains u.

*/

func isBipartite(graph [][]int) bool {
    return isBipartite_BFS(graph)
}
func isBipartite_BFS(graph [][]int) bool {
    n := len(graph)
    flags := make([]int, n) // 1, 0, -1

    for i := range graph {
        if flags[i] != 0 {
            continue
        }
        if !bfs(graph, flags, i, 1) {
            return false
        }
    }
    return true
}

func bfs(graph [][]int, flags []int, cur, flg int) bool {
    q := []int{cur}
    flags[cur] = flg
    i := 0
    for i < len(q) {
        j := len(q)
        for ; i < j; i++ {
            for _, node := range graph[q[i]] {
                if flags[node] == flg {
                    return false
                } else if flags[node] == -flg {
                    continue
                } else {
                    flags[node] = -flg
                    q = append(q, node)
                }
            }
        }
        flg = -flg
    }
    return true
}

func isBipartite_DFS(graph [][]int) bool {
    n := len(graph)
    flags := make([]int, n) // 1, 0, -1
    for i := range graph {
        if flags[i] != 0 {
            continue
        }
        if !dfs(graph, flags, i, 1) {
            return false
        }
    }
    return true
}
func dfs(graph [][]int, flags []int, cur, flg int) bool {
    if flags[cur] == flg {
        return true
    } else if flags[cur] == -flg {
        return false
    }

    flags[cur] = flg
    for _, node := range graph[cur] {
        if !dfs(graph, flags, node, -flg) {
            return false
        }
    }
    return true
}
