package merge_two_sorted_list

import "testing"

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{
			name:  "test1",
			list1: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
			list2: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
			want:  &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}},
		},
		{
			name:  "test2",
			list1: nil,
			list2: nil,
			want:  nil,
		},
		{
			name:  "test3",
			list1: nil,
			list2: &ListNode{Val: 0},
			want:  &ListNode{Val: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.list1, tt.list2); !isEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isEqual(list1 *ListNode, list2 *ListNode) bool {
	for list1 != nil && list2 != nil {
		if list1.Val != list2.Val {
			return false
		}
		list1 = list1.Next
		list2 = list2.Next
	}
	return list1 == nil && list2 == nil
}
