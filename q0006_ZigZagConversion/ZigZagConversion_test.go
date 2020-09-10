package q0006

import "testing"

func Test_convert(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		numRows int
		want    string
	}{
		{"Example 1", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"Example 2", "PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"Wrong Answer 1", "PAYPALISHIRING", 5, "PHASIYIRPLIGAN"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.s, tt.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
