# Graph Algorithms

Shortest path
* Dijkstra
  * use of Priority queue 
* Floyd-Warshall
  * multi-source shortest path 
  * relaxation of distance for V times
* Bellman-Ford
  * single source shortest path
  * relaxation of distance for V times

Minimum Spanning Tree
* Prim's
  * priority queue to find shortest path from established tree to unconnected nodes 
* Kruskal
  * Priotity queue to find shortest edge
  * Then use union-find/disjoint set to find out whether the two vertices are already connected or not

Topological Sort
* Kahn's
* DFS
```
  int n; // number of vertices
vector<vector<int>> adj; // adjacency list of graph
vector<bool> visited;
vector<int> ans;

void dfs(int v) {
    visited[v] = true;
    for (int u : adj[v]) {
        if (!visited[u]) {
            dfs(u);
        }
    }
    ans.push_back(v);
}

void topological_sort() {
    visited.assign(n, false);
    ans.clear();
    for (int i = 0; i < n; ++i) {
        if (!visited[i]) {
            dfs(i);
        }
    }
    reverse(ans.begin(), ans.end());
}
```

# post-order DFS graph traversal

```
vector<vector<int>> adj; // graph represented as an adjacency list
int n; // number of vertices

vector<int> color;

vector<int> time_in, time_out;
int dfs_timer = 0;

void dfs(int v) {
    time_in[v] = dfs_timer++;
    color[v] = 1;
    for (int u : adj[v])
        if (color[u] == 0)
            dfs(u);
    color[v] = 2;
    time_out[v] = dfs_timer++;
}
```

* Tarjan
* Find bridge in graphs
* Eulerian Cycles or Path
  * Hierholzer's 
```
procedure FindEulerPath(V)
  1. iterate through all the edges outgoing from vertex V;
       remove this edge from the graph,
       and call FindEulerPath from the second end of this edge;
  2. add vertex V to the answer.
```

# Reference
* https://cp-algorithms.com/graph/depth-first-search.html
