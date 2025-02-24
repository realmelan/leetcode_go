/*
103. Binary Tree Zigzag Level Order Traversal
Solved
Medium
Topics
Companies
Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).

 

Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: [[3],[20,9],[15,7]]
Example 2:

Input: root = [1]
Output: [[1]]
Example 3:

Input: root = []
Output: []
 

Constraints:

The number of nodes in the tree is in the range [0, 2000].
-100 <= Node.val <= 100
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
    var queue []*TreeNode
    if root != nil {
        queue = append(queue, root)
    }
    i, k := 0, 0
    for i < len(queue) {
        start := i
        j := len(queue)
        line := make([]int, j-start)
        for i < j {
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }

            if k%2==0 {
                line[i-start] = queue[i].Val
            } else {
                line[j-i-1] = queue[i].Val
            }

            i++
        }
        k++
        res = append(res, line)
    }
    return res
}
