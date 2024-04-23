func isPalindrome(x int) bool {
	k := 0
	s := x
	if x < 0 {
		return false
	}
	for x > 0 {
		k = k*10 + x%10
		x = x / 10
	}
	if k == s {
		return true
	} else {
		return false
	}
}