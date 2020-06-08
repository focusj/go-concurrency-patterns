package main

import (
	"fmt"
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
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("exit")
}
