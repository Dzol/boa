package main

import "fmt"

func main() {

	var ping func ()
	var pong func ()

	c := make(chan uint)

	ping = func () {
		var n uint
		n = <- c
		fmt.Println("ping: ", n)
		if n < 10 {
			n++
			c <- n
			ping()
		}
	}

	pong = func () {
		var n uint
		n = <- c
		fmt.Println("pong: ", n)
		if n < 10 {
			n++
			c <- n
			pong()
		}
	}

	go ping()
	go pong()

	c <- 0
}
