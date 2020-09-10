package q0004

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Example 1", args{[]int{1, 3}, []int{2}}, 2.0},
		{"Example 2", args{[]int{1, 2}, []int{3, 4}}, 2.5},
		{"Example 3", args{[]int{0, 0}, []int{0, 0}}, 0},
		{"Example 4", args{[]int{}, []int{1}}, 1},
		{"Example 5", args{[]int{2}, []int{}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
