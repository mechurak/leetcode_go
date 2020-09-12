package q0017_LetterCombinations

import (
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{"Example 1", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations(tt.digits); !check(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func check(got []string, want []string) bool {
	if len(got) != len(want) {
		return false
	}

	wantMap := map[string]bool{}
	for _, letter := range want {
		wantMap[letter] = true
	}

	for _, letter := range got {
		_, ok := wantMap[letter]
		if !ok {
			return false
		}
	}
	return true
}
