package q0015_3Sum

import (
	"testing"
)

func Test_threeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"Example 1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"Example 2", []int{}, [][]int{}},
		{"Example 3", []int{0}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.nums); !check(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func check(got [][]int, want [][]int) bool {
	if len(got) != len(want) {
		return false
	}

	wantMap := map[[3]int]bool{}
	for _, triplet := range want {
		wantMap[[3]int{triplet[0], triplet[1], triplet[2]}] = true
	}

	for _, triplet := range got {
		_, ok := wantMap[[3]int{triplet[0], triplet[1], triplet[2]}]
		if !ok {
			return false
		}
	}
	return true
}
