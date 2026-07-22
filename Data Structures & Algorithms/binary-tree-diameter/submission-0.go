/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
		return 0
	}
	res := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftRes := dfs(node.Left)
		rightRes := dfs(node.Right)
		res = max(res, leftRes+rightRes)

		return 1 + max(leftRes, rightRes)
	}
	dfs(root)

	return res
}
