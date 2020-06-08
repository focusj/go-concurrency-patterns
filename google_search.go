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
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %s", kind, query))
	}
}

func Google(query string) (result []Result) {
	web := Web(query)
	image := Image(query)
	video := Video(query)

	result = append(result, web, image, video)

	return result
}

func main() {
	rs := Google("tom & jerry")
	fmt.Println(rs)
}
