package Q0002_AddTwoNumbers

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"tc1",
			args{
				&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, // 342
				&ListNode{5, &ListNode{6, &ListNode{4, nil}}}, // 465
			},
			&ListNode{7, &ListNode{0, &ListNode{8, nil}}}, // 807
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
