/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}

	balanced := true
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftRes := dfs(node.Left)
		rightRes := dfs(node.Right)
		if (math.Abs(float64(leftRes) - float64(rightRes))) > 1 {
			balanced = false
		}
		return 1 + max(leftRes, rightRes)
	}
	dfs(root)

	return balanced
}
