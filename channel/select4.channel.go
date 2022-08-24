package main

import (
	"fmt"
	"time"
)

func main() {
	rand := make(chan int)

	go func() {
		for {
			select {
			case rand <- 0: // no statement
			case rand <- 1:
			}
			time.Sleep(time.Second)
		}
	}()

	for data := range rand {
		fmt.Println(data)
	}
}
