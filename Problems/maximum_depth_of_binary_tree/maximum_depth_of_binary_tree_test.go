package maximum_depth_of_binary_tree

import "testing"

func TestMaximumDepthOfBinaryTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 5,
						},
					},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		got := maximumDepthOfBinaryTree(tt.root)
		if got != tt.want {
			t.Errorf("maximumDepthOfBinaryTree() = %v, want %v", got, tt.want)
		}
	}
}
