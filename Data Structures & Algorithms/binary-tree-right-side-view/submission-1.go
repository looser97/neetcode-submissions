/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
    result := []int{}

	q := []*TreeNode{root}

	for len(q) > 0 {
		levelSize := len(q)
		rightMostElement := q[len(q)-1]
		result = append(result, rightMostElement.Val)

		for levelSize > 0 {
			item := q[0]
			q = q[1:]

			if item.Left != nil {
				q = append(q, item.Left)
			}

			if item.Right != nil {
				q = append(q, item.Right)
			}
			levelSize--
		}

	}
	return result
}
