package maximum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maximumDepthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maximumDepthOfBinaryTree(root.Left)
	rightDepth := maximumDepthOfBinaryTree(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
