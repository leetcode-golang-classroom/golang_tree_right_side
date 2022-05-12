package sol

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		qLen := len(queue)
		var rightSide *TreeNode
		// loop over current queue find rightest
		for idx := 0; idx < qLen; idx++ {
			// shift left ele
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				// 更新
				rightSide = node
				queue = append(queue, node.Left, node.Right)
			}
		}
		if rightSide != nil {
			result = append(result, rightSide.Val)
		}
	}
	return result
}
