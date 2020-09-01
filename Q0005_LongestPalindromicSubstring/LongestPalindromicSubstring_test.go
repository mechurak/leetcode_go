package q0005

import "testing"

func Test_longestPalindrome(t *testing.T) {
    tests := []struct {
        name string
        s string
        wantMap map[string]bool
    }{
        {"Example 1", "babad", map[string]bool{"bab":true, "aba":true}},
        {"Example 2", "cbbd", map[string]bool{"bb":true}},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := longestPalindrome(tt.s); !tt.wantMap[got] {
                t.Errorf("longestPalindrome() = %v, want %v", got, tt.wantMap)
            }
        })
    }
}
