/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	level := 0

	q := []*TreeNode{root}

	for len(q) > 0 {
		size := len(q)

		for size > 0 {
			element := q[0]
			q = q[1:]

			if element.Left != nil {
				q = append(q, element.Left)
			}

			if element.Right != nil {
				q = append(q, element.Right)
			}
			size--
		}

		level++

	}

	return level
}
