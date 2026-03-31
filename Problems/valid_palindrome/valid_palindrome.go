package valid_palindrome

import (
	"regexp"
	"strings"
)

func validPalindrome(s string) bool {
	s = regexp.MustCompile("[^a-zA-Z0-9]").ReplaceAllString(s, "")
	t := (strings.ToLower(s))
	i, j := 0, len(t)-1
	for i < j {
		if t[i] != t[j] {
			return false
		}
		i++
		j--
	}
	return true
}
