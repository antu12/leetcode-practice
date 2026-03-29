package minimum_depth_of_binary_tree

import "testing"

func TestMinimumDepthOfBinaryTree(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			expected: 2,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: 2,
		},
	}

	for _, test := range tests {
		result := minimumDepthOfBinaryTree(test.root)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}
