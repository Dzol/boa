package main

import "fmt"
import "encoding/json"
import "os"
import "bufio"
import "strings"

func main() {
	person := make(map[string]string)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter a name: ")
	name, _ := reader.ReadString('\n')
	name = strings.Trim(name, "\n")

	fmt.Print("Please enter an address: ")
	address, _ := reader.ReadString('\n')
	address = strings.Trim(address, "\n")

	person["name"] = name
	person["address"] = address

	o, _ := json.Marshal(person)

	os.Stdout.Write(o)
	fmt.Println()
}
