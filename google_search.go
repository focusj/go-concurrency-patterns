package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result string

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %s", kind, query))
	}
}

func Google(query string) (result []Result) {
	c := make(chan Result)
	go func() { c <- First(query, Web) }()
	go func() { c <- First(query, Image) }()
	go func() { c <- First(query, Video) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case r := <-c:
			result = append(result, r)
		case <-timeout:
			fmt.Println("timed out.")
			return
		}
	}

	return result
}

func First(query string, replicas ...Search) Result {
	r := make(chan Result)
	for _, replica := range replicas {
		go func() {
			r <- replica(query)
		}()
	}
	return <-r
}

func main() {
	rs := Google("tom & jerry")
	fmt.Println(rs)
}
