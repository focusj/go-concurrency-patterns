package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

// 这种模式的核心要点：
//     1. 每个消息中引入一个chan，控制当前的chan是否要进行下一轮消息发送；
//     2. 通过fanIn将两个chan进行合并，配合上边的chan确定消息的消费顺序；
//
// 消息拓扑：
// chan1 ------                         ------> msg1/waitChan
//             |  ------> fanIn ------>|
// chan2 ------                         ------> msg2/waitChan
func main() {
	c := fanIn(boring("Lily"), boring("Lucy"))

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)

		msg1.wait <- true
		msg2.wait <- true
	}

}

// message controller
func fanIn(chan1, chan2 <-chan Message) <-chan Message {
	c := make(chan Message)

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

func boring(msg string) <-chan Message {
	c := make(chan Message)
	wait := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s: %d", msg, i), wait}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Microsecond)

			<-wait // wait a msg to continue next msg delivery
		}
	}()

	return c
}
