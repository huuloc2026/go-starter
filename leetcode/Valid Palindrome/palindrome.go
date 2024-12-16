package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "race a car"
	isPalindrome(s)
}
func isPalindrome(s string) bool {
	if len(s) != 1 {
		result := filterAndLowercase(s)
		fmt.Println(result)
		left := 0
		right := len(result) - 1
		for left < right {
			if result[left] == result[right] {
				left++
				right--
			}
		}
		return false
	}
	return true
}

func filterAndLowercase(s string) string {
	var result strings.Builder
	for _, char := range s {
		if unicode.IsLetter(char) {
			result.WriteRune(unicode.ToLower(char)) // Chuyển thành chữ thường
		}
	}
	return result.String()
}
