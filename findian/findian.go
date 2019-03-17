package main

import "os"
import "bufio"
import "fmt"
import "strings"
import "regexp"

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter some text: ")
	s, _ := reader.ReadString('\n')
	s = strings.Trim(s, "\n")

	match, _ := regexp.MatchString("^[Ii].*[Aa].*[Nn]$", s)

	if match {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
