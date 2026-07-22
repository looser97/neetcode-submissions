/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
    q := []*TreeNode{root}

	for len(q) > 0 {
		element := q[0]
		q = q[1:]
		element.Left, element.Right = element.Right, element.Left
		if element.Left != nil {
			q = append(q, element.Left)
		}
		if element.Right != nil {
			q = append(q, element.Right)
		}
	}
	return root
}
