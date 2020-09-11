package q0011_ContainerWithMostWater

import "testing"

func Test_maxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"Example 1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
