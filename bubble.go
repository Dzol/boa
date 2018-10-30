package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	s := make([]int, 0, 0)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter up to 10 integers: ")

	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\n")
	x := strings.Split(input, " ")

	for _, v := range x {
		integer, _ := strconv.Atoi(v)
		s = append(s, integer)
	}

	bubbleSort(s)
	fmt.Println(s)
}

func bubbleSort(s []int) {
	swapped := false
	for i, j := 0, 1; i < len(s) - 1 && j < len(s); i, j = i+1, j+1 {
		if s[i] > s[j] {
			swap(&s[i], &s[j])
			swapped = true
		}
	}
	if swapped {
		bubbleSort(s)
	}
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
