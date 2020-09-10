package q0012_IntegerToRoman

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{"Example 1", 3, "III"},
		{"Example 2", 4, "IV"},
		{"Example 3", 9, "IX"},
		{"Example 4", 58, "LVIII"},
		{"Example 5", 1994, "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
