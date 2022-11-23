package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	result := make(chan interface{})
	unloading := func(channel <-chan interface{}) {
		for val := range channel {
			result <- val
		}
		wg.Done()
	}

	wg.Add(len(channels))
	for _, v := range channels {
		go unloading(v)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(1*time.Second),
		sig(1*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(2*time.Second),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}
