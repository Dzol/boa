package main

import "fmt"
import "sync"

func main() {

	var wg sync.WaitGroup

	var ping func ()
	var pong func ()

	c := make(chan uint)
	d := make(chan uint)

	ping = func () {
		var n uint
		defer wg.Done()
		for {
			n = <- c
			fmt.Println("ping: ", n)
			if n < 10 {
				n++
				d <- n
			} else {
				close(d)
				break
			}
		}
	}

	pong = func () {
		defer wg.Done()
		for {
			if n, ok := <- d; ok {
				fmt.Println("pong: ", n)
				n++
				c <- n
			} else {
				break
			}
		}
	}

	wg.Add(2)

	go ping()
	go pong()

	c <- 0

	wg.Wait()
}
