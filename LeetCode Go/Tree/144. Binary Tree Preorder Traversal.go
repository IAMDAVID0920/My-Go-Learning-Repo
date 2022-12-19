package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	// var res [] int
	// if root != nil{
	//     res = append(res, root.Val)
	//     res = append(res, preorderTraversal(root.Left)...)
	//     res = append(res, preorderTraversal(root.Right)...)
	// }
	// return res

	res := []int{}

	var preorder func(*TreeNode)

	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}

	preorder(root)
	return res
}
