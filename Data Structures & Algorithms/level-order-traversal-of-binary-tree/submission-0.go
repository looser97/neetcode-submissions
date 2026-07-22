/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
    result := [][]int{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		levelSize := len(q)
		levelItems := []int{}

		for levelSize > 0 {
			item := q[0]
			q = q[1:]
			levelItems = append(levelItems, item.Val)
			if item.Left != nil {
				q = append(q, item.Left)
			}
			if item.Right != nil {
				q = append(q, item.Right)
			}
			levelSize--
		}
		result = append(result, levelItems)

	}
	return result
}
