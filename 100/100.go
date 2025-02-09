package _00

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	switch {
	case p == nil && q != nil:
		return false
	case p != nil && q == nil:
		return false
	case p == nil && q == nil:
		return true
	case p.Val != q.Val:
		return false
	}
	ok := isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
	return ok && p.Val == q.Val
}
