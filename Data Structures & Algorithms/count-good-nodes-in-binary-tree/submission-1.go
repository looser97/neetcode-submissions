/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    count := 0

	var dfs func(node *TreeNode, maxi int)

	dfs = func(node *TreeNode, maxi int) {
		if node == nil {
			return
		}
		if node.Val >= maxi {
			count++
		}

		if node.Left != nil {
			dfs(node.Left, max(maxi, node.Val))
		}

		if node.Right != nil {
			dfs(node.Right, max(maxi, node.Val))
		}
	}

	dfs(root, root.Val)
	return count
}
