package main

import (
	"fmt"
	"time"
)

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	or = orChannel
	start := time.Now()
	<-or(
		sig(1*time.Second),
		sig(4*time.Second),
		sig(10*time.Second),
		sig(8*time.Second),
		sig(6*time.Second),
	)
	fmt.Printf("fone after %v \n", time.Since(start))
}

func orChannel(channels ...<-chan interface{}) <-chan interface{} {
	result := make(chan interface{})
	for _, channel := range channels {
		go func(ch <-chan interface{}) {
			for {
				select {
				case <-ch:
					result <- 1

				}
			}

		}(channel)
	}
	select {
	case <-result:
		return result
	}

}
