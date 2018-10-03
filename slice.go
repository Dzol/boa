package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "sort"

const terminal = "X"

func main() {
	s := make([]int, 0, 3)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Please enter an integer (or `%s` to stop): ", terminal)
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")
		if input == terminal {
			break
		} else {
			i, _ := strconv.Atoi(input)
			s = append(s, i)
		}
	}
	sort.Ints(s)
	fmt.Println(s)
}
