package main

import "fmt"

func main() {
	fmt.Println(isValid("]"))
	fmt.Println(isValid("(([]){})"))
	fmt.Println(isValid("({)}"))
}

func isValid(s string) bool {
	st := make([]int32, len(s))
	si := 0
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			st[si] = c
			si++
		} else {
			if (si > 0) && ((c == ')' && st[si-1] == '(') ||
				(c == ']' && st[si-1] == '[') ||
				(c == '}' && st[si-1] == '{')) {
				si--
			} else {
				return false
			}
		}
	}
	return si == 0
}
