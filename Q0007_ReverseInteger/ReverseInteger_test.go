package q0007

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"Example 1", 123, 321},
		{"Example 2", -123, -321},
		{"Example 3", 120, 21},
		{"TC 1", 21474_47412, 21474_47412},
		{"TC 2", 21474_47413, 0},
		{"TC 3", -21474_47412, -21474_47412},
		{"Runtime Error 1", 0, 0},
		{"Wrong Answer 1", 15342_36469, 0},
		{"Wrong Answer 2", -21474_83412, -2143847412},
		{"Wring Answer 3", 15638_47412, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
