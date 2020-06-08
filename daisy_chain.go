package main

import "fmt"

func f(left, right chan int) {
	left <- 1 + <-right
}

// daisy_chain is a circle network topology like this:
//           c  -  d
//         /         \
//        b           e
//        |           |
//        a           f
//         \         /
//           g  -  h
//
func main() {
	leftmost := make(chan int)
	left := leftmost
	right := leftmost

	for i := 0; i < 1000; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
