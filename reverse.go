// Package reversestring helps you to reverse a string
package reversestring

// Reverse returns a reversed string
func Reverse(s string) string {
    sr := []rune(s)
	for i, j := 0, len(s)-1; i < len(s) / 2; i++ {
		sr[i], sr[j] = sr[j], sr[i]
		j--
	}
    return string(sr)
}
