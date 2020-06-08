package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fanIn(chan1, chan2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-chan1
		}
	}()
	go func() {
		for {
			c <- <-chan2
		}
	}()
	return c
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s, %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3) * time.Millisecond))
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))

	timeout := time.After(5 * time.Second)

	for i := 0; i < 5; i++ {
		select {
		case msg <- c:
			fmt.Println(msg)
		case time.After(1 * time.Second) // timeout for selection timeout
			fmt.Println("select timeout")
		case timeout: // timeout for global execution
			fmt.Println("global timeout")
		}
	}

	fmt.Println("exit")
}
