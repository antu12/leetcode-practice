package minimum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumDepthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := minimumDepthOfBinaryTree(root.Left)
	rightDepth := minimumDepthOfBinaryTree(root.Right)

	if leftDepth == 0 || rightDepth == 0 {
		return leftDepth + rightDepth + 1
	} else {
		return min(leftDepth, rightDepth) + 1
	}
}
