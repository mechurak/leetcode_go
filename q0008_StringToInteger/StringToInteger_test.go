package q0008

import "testing"

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{"Example 1", "42", 42},
		{"Example 2", "   -42", -42},
		{"Example 3", "4193 with words", 4193},
		{"Example 4", "words and 987", 0},
		{"Example 5", "-91283472332", -2147483648}, // INT_MIN
		{"Runtime Error 1", "-912-8332", -912},
		{"TC 1", "-+912", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
