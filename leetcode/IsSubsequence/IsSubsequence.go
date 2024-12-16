package main

import "fmt"

func main() {
	s := "axc"
	t := "ahbgdc"

	fmt.Println(isSubsequence(s, t))

}

func isSubsequence(s string, t string) bool {
	pS, pT := 0, 0
	n := len(t) - 1
	for pT <= n {
		if pS < len(s) && s[pS] == t[pT] {
			pS++
		}
		pT++
		if pS == len(s) {
			return true
		}
	}
	return false
}
