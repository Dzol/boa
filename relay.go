package main

import "fmt"
import "sync"

var wg sync.WaitGroup

var c = make(chan uint)
var d = make(chan uint)

func main() {

	wg.Add(2)

	go ping()
	go pong()

	c <- 0

	wg.Wait()
}

func ping() {
	defer wg.Done()
	for {
		n := <-c
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

func pong() {
	defer wg.Done()
	for {
		if n, ok := <-d; ok {
			fmt.Println("pong: ", n)
			n++
			c <- n
		} else {
			break
		}
	}
}
