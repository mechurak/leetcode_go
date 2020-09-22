package q0020_ValidParentheses

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"Example 1", "()", true},
		{"Example 2", "()[]{}", true},
		{"Example 3", "(]", false},
		{"Example 4", "([)]", false},
		{"Example 5", "{[]}", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
