/*
1483. Kth Ancestor of a Tree Node
Solved
Hard
Topics
Companies
Hint
You are given a tree with n nodes numbered from 0 to n - 1 in the form of a parent array parent where parent[i] is the parent of ith node. The root of the tree is node 0. Find the kth ancestor of a given node.

The kth ancestor of a tree node is the kth node in the path from that node to the root node.

Implement the TreeAncestor class:

TreeAncestor(int n, int[] parent) Initializes the object with the number of nodes in the tree and the parent array.
int getKthAncestor(int node, int k) return the kth ancestor of the given node node. If there is no such ancestor, return -1.
 

Example 1:


Input
["TreeAncestor", "getKthAncestor", "getKthAncestor", "getKthAncestor"]
[[7, [-1, 0, 0, 1, 1, 2, 2]], [3, 1], [5, 2], [6, 3]]
Output
[null, 1, 0, -1]

Explanation
TreeAncestor treeAncestor = new TreeAncestor(7, [-1, 0, 0, 1, 1, 2, 2]);
treeAncestor.getKthAncestor(3, 1); // returns 1 which is the parent of 3
treeAncestor.getKthAncestor(5, 2); // returns 0 which is the grandparent of 5
treeAncestor.getKthAncestor(6, 3); // returns -1 because there is no such ancestor
 

Constraints:

1 <= k <= n <= 5 * 104
parent.length == n
parent[0] == -1
0 <= parent[i] < n for all 0 < i < n
0 <= node < n
There will be at most 5 * 104 queries.
*/
type TreeAncestor struct {
    powerAn [][]int // layer 1: 2^k (k=0,1,2,3...), layer 2: which node
}


func Constructor(n int, parent []int) TreeAncestor {
    // for each node, create a 1, 2, 4, 8, 16,... ancestor map
    var powerAn [][]int
    p := parent
    for {
        powerAn = append(powerAn, p)
        // now how to construct next level parent
        np := make([]int, n)
        for i, v := range p {
            if v == -1 {
                np[i] = -1
            } else {
                np[i] = p[v]
            }
        }
        p = np

        // when to stop?
        cnt := 0
        for _, v := range p {
            if v == -1 {
                cnt++
            }
        }
        if cnt == n {
            powerAn = append(powerAn, p)
            break 
        }
    }
    return TreeAncestor{
        powerAn: powerAn,
    }
}


func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
    p := node
    for i := 0; i < len(this.powerAn); i++ {
        j := k & 1
        k >>=1
        if j > 0 {
            p = this.powerAn[i][p]
            if p == -1 {
                break
            }
        }
    }
    if k > 0 {
        return -1
    }
    return p
}


/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */
