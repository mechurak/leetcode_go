package q0019_RemoveNthNode

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		n    int
		want *ListNode
	}{
		{"Example 1",
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2,
			&ListNode{1, &ListNode{2, &ListNode{4, &ListNode{5, nil}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.head, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
