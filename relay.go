package main

import (
	"fmt"
	"time"
)

type ball uint

func main() {
	table := make(chan *ball)

	go player("ping", table)
	go player("pong", table)

	table <- new(ball)
	time.Sleep(1 * time.Second)
	<-table
}

func player(name string, table chan *ball) {
	for {
		ball := <-table
		(*ball)++
		fmt.Println(name, *ball)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
