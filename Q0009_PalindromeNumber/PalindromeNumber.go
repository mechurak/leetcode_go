package q0009

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var digits []int
	n := x
	for n > 0 {
		l := n % 10
		digits = append(digits, l)
		n = n / 10
	}
	size := len(digits)
	for i := 0; i < size/2; i++ {
		if digits[i] != digits[size-1-i] {
			return false
		}
	}
	return true
}
