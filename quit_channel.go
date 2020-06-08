package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan bool)
	c := boring("Joe", quit)

	for {
		select {
		case msg := <-c:
			fmt.Println(msg)
		case <-quit:
			fmt.Println("quit")
			// quit <- sth. // receive something when quit
			return
		}
	}
}

func boring(msg string, signal chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 10; i >= 0; i-- {
			c <- fmt.Sprintf("%s, %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		signal <- true
	}()
	return c
}
