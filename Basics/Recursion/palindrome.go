package main

import (
	"fmt"
	"strings"
)

func palindromeHelper(s string, strt, end int) bool {
	if strt > end {
		return true
	}
	if !isValid(s[strt]) {
		return palindromeHelper(s, strt+1, end)
	}
	if !isValid(s[end]) {
		return palindromeHelper(s, strt, end-1)
	}
	if strings.EqualFold(string(s[strt]), string(s[end])) {
		return palindromeHelper(s, strt+1, end-1)
	}

	return false

}

func palindrome(s string) bool {
	return palindromeHelper(s, 0, len(s)-1)
}

func isValid(s byte) bool {
	if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9') {
		return true
	}

	return false
}

func main() {
	fmt.Println(palindrome("A man, a plan, a canal: Panama")) // true: amanaplanacanalpanama
}
