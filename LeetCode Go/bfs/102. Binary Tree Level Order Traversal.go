package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	// check if root is nil
	if root == nil {
		return nil
	}

	// otherwise start to traversal
	var res [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		curLevel := make([]int, size)

		for i := 0; i < size; i++ {
			curLevel[i] = queue[i].Val

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		res = append(res, curLevel)
		queue = queue[size:]

	}
	return res

}
