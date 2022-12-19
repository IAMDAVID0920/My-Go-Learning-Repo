package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	// 1 way
	// var res [] int
	// if root != nil{
	//     res = append(res, inorderTraversal(root.Left)...)
	//     res = append(res, root.Val)
	//     res = append(res, inorderTraversal(root.Right)...)
	// }
	// return res

	// another way
	if root == nil {
		return []int{}
	}

	return append(
		append(
			inorderTraversal(root.Left), root.Val),
		inorderTraversal(root.Right)...)
}
