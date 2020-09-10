package q0013_RomanToInteger

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"Example 1", "III", 3},
		{"Example 2", "IV", 4},
		{"Example 3", "IX", 9},
		{"Example 4", "LVIII", 58},
		{"Example 5", "MCMXCIV", 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
