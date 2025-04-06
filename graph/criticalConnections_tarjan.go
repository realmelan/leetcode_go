/*
1192. Critical Connections in a Network
Solved
Hard
Topics
Companies
Hint
There are n servers numbered from 0 to n - 1 connected by undirected server-to-server connections forming a network where connections[i] = [ai, bi] represents a connection between servers ai and bi. Any server can reach other servers directly or indirectly through the network.

A critical connection is a connection that, if removed, will make some servers unable to reach some other server.

Return all critical connections in the network in any order.

 

Example 1:


Input: n = 4, connections = [[0,1],[1,2],[2,0],[1,3]]
Output: [[1,3]]
Explanation: [[3,1]] is also accepted.
Example 2:

Input: n = 2, connections = [[0,1]]
Output: [[0,1]]
 

Constraints:

2 <= n <= 105
n - 1 <= connections.length <= 105
0 <= ai, bi <= n - 1
ai != bi
There are no repeated connections.
*/
func criticalConnections(n int, connections [][]int) [][]int {
    // solution 1: edges between scc(strongly connected components) are critical connections.
    // solution 2: test each edge using union-find approach. (bad time complexity)
    conns := make([]map[int]bool, n)
    for i := range conns {
        conns[i] = make(map[int]bool)
    }
    for _, c := range connections {
        conns[c[0]][c[1]]=false
        conns[c[1]][c[0]]=false
    }

    index := 1
    tin := make([]int, n)
    low := make([]int, n)
    var res [][]int
    var dfs func(int)
    dfs = func(cur int) {
        tin[cur] = index
        low[cur] = index
        index++
        for v, used := range conns[cur] {
            if used {
                continue
            }
            conns[cur][v] = true
            conns[v][cur] = true
            if tin[v] <= 0 {
                dfs(v)
                low[cur] = min(low[cur], low[v])
                if low[v] > tin[cur] {
                    res = append(res, []int{cur, v})
                }
            } else {
                low[cur] = min(low[cur], low[v])
            }
        }
    }
    dfs(0)
    return res
}
