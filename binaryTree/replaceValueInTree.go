/*
2641. Cousins in Binary Tree II
Solved
Medium
Topics
Companies
Hint
Given the root of a binary tree, replace the value of each node in the tree with the sum of all its cousins' values.

Two nodes of a binary tree are cousins if they have the same depth with different parents.

Return the root of the modified tree.

Note that the depth of a node is the number of edges in the path from the root node to it.

 

Example 1:


Input: root = [5,4,9,1,10,null,7]
Output: [0,0,0,7,7,null,11]
Explanation: The diagram above shows the initial binary tree and the binary tree after changing the value of each node.
- Node with value 5 does not have any cousins so its sum is 0.
- Node with value 4 does not have any cousins so its sum is 0.
- Node with value 9 does not have any cousins so its sum is 0.
- Node with value 1 has a cousin with value 7 so its sum is 7.
- Node with value 10 has a cousin with value 7 so its sum is 7.
- Node with value 7 has cousins with values 1 and 10 so its sum is 11.
Example 2:


Input: root = [3,1,2]
Output: [0,0,0]
Explanation: The diagram above shows the initial binary tree and the binary tree after changing the value of each node.
- Node with value 3 does not have any cousins so its sum is 0.
- Node with value 1 does not have any cousins so its sum is 0.
- Node with value 2 does not have any cousins so its sum is 0.
 

Constraints:

The number of nodes in the tree is in the range [1, 105].
1 <= Node.val <= 104
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func replaceValueInTree(root *TreeNode) *TreeNode {
    // use BFS
    q := make([]*TreeNode, 0)
    q = append(q, root)
    i := 0
    for i < len(q) {
        j := len(q)
        sum := 0
        var p []*TreeNode
        for i < j {
            p = append(p, q[i])
            if q[i].Left != nil {
                q = append(q, q[i].Left)
                sum += q[i].Left.Val
            }
            if q[i].Right != nil {
                q = append(q, q[i].Right)
                sum += q[i].Right.Val
            }
            i++
        }

        for _, node := range p {
            val := sum
            if node.Left != nil {
                val -= node.Left.Val
            }
            if node.Right != nil {
                val -= node.Right.Val
            }

            if node.Left != nil {
                node.Left.Val = val
            }
            if node.Right != nil {
                node.Right.Val = val
            }
        }
    }
    root.Val = 0
    return root
 }
func replaceValueInTree2(root *TreeNode) *TreeNode {
    // first pass: build level sum using a list
    // second pass: build new tree
    if root == nil {
        return nil
    }
    m := make(map[int]int)
    sum(root, 0, m)
    return build(root, root.Val, 0, m)
}

func sum(root *TreeNode, level int, levelSum map[int]int) {
    if root == nil {
        return
    }
    levelSum[level] += root.Val
    sum(root.Left, level+1, levelSum)
    sum(root.Right, level+1, levelSum)
}

func build(root *TreeNode, siblingSum int, level int, levelSum map[int]int) *TreeNode {
    res := &TreeNode{
        Val: levelSum[level] - siblingSum,
    }

    siblingSum = 0
    if root.Left != nil {
        siblingSum += root.Left.Val
    }
    if root.Right != nil {
        siblingSum += root.Right.Val
    }
    if root.Left != nil {
        res.Left = build(root.Left, siblingSum, level+1, levelSum)
    }
    if root.Right != nil {
        res.Right = build(root.Right, siblingSum, level+1, levelSum)
    }
    return res
}
