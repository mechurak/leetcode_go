package q0014_LongestCommonPrefix

import "testing"

func Test_isMatch(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{"Example 1", []string{"flower", "flow", "flight"}, "fl"},
		{"Example 2", []string{"dog", "racecar", "car"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.strs); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
