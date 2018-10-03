package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "io"

type Name struct {
	fname string
	lname string
}

func main() {
	people := make([]Name, 0)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter a file name: ")
	name, _ := reader.ReadString('\n')
	name = strings.Trim(name, "\n")

	f, _ := os.Open(name)
	r := bufio.NewReader(f)

	for {
		f, _ := r.ReadString(' ')
		f = strings.Trim(f, " ")
		f = truncate(f)

		l, e := r.ReadString('\n')
		if e == io.EOF {
			break
		}
		l = strings.Trim(l, "\n")
		l = truncate(l)

		people = append(people, Name{f, l})
	}

	for _, n := range people {
		fmt.Printf("first name: %s, last name: %s\n", n.fname, n.lname)
	}
}

func truncate(s string) string {
	i := min(len(s), 20)
	return s[:i]
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
