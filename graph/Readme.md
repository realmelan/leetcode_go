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
  *  

# post-order DFS graph traversal

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
