package main

import "fmt"
import "sort"
import "strconv"

func main() {
	x := []int{} // don't have to `make([]int, 0, 0)`
	for i := 999; i > 100; i-- {
		for j := 999; j > i; j-- {
			if p := i * j; isPalindrome(strconv.Itoa(p)) {
				x = append(x, p)
			}
		}
	}
	fmt.Println(max(x))
}

func isPalindrome(s string) bool {
	return s == reverse(s)
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func max(s []int) int {
	sort.Ints(s)
	return s[len(s)-1]
}
