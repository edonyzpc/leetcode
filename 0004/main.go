package main

import (
	"fmt"
)

const LIMITS = 10

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
		fmt.Println("i=", i, " sent to", &ch)
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		fmt.Println("recived ", &in)
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	var ch = make(chan int)
	go Generate(ch)
	// Daisy-chain process
	for i := 0; i < LIMITS; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}
