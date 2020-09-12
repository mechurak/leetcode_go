package q0018_4Sum

import (
	"testing"
)

func Test_threeSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{"Example 1", []int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-1,  0, 0, 1}, {-2, -1, 1, 2}, {-2,  0, 0, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.nums, tt.target); !check(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func check(got [][]int, want [][]int) bool {
	if len(got) != len(want) {
		return false
	}

	wantMap := map[[4]int]bool{}
	for _, quadruplet := range want {
		wantMap[[4]int{quadruplet[0], quadruplet[1], quadruplet[2], quadruplet[3]}] = true
	}

	for _, quadruplet := range got {
		_, ok := wantMap[[4]int{quadruplet[0], quadruplet[1], quadruplet[2], quadruplet[3]}]
		if !ok {
			return false
		}
	}
	return true
}
