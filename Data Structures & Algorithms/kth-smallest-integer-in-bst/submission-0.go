/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    counter := k

	var dfs func(node *TreeNode)
	ans := -1

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		counter--
		if counter == 0 {
			ans = node.Val
			return
		}
		if node.Right != nil {
			dfs(node.Right)
		}

	}
	dfs(root)
	return ans
}
