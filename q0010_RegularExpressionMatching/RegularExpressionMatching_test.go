package q0010

import "testing"

func Test_isMatch(t *testing.T) {
    tests := []struct {
        name string
        s  string
        p  string
        want bool
    }{
        {"Example 1", "aa", "a", false},
        {"Example 2", "aa", "a*", true},
        {"Example 3", "ab", ".*", true},
        {"Example 4", "aab", "c*a*b", true},
        {"Example 5", "mississippi", "mis*is*p*.", false},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := isMatch(tt.s, tt.p); got != tt.want {
                t.Errorf("isMatch() = %v, want %v", got, tt.want)
            }
        })
    }
}
