/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil && subRoot == nil {
		return true
	}

	q := []*TreeNode{root}

	var isSameTree func(p *TreeNode, q *TreeNode) bool

	isSameTree = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if p != nil && q != nil && p.Val == q.Val {
			return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}

		return false
	}

	for len(q) > 0 {
		item := q[0]
		q = q[1:]
		if item.Val == subRoot.Val {
			isSame := isSameTree(item, subRoot)
			if isSame {
				return true
			}
		}
		if item.Left != nil {
			q = append(q, item.Left)
		}

		if item.Right != nil {
			q = append(q, item.Right)
		}
	}

	return false
}
