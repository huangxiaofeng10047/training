package main

import (
	"fmt"
	"sync"
)

func main() {
	str := []byte("foobar")
	ch := make(chan byte, len(str))
	next := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(2)

	for i := 0; i < len(str); i++ {
		ch <- str[i]
	}

	close(ch)

	go func() {
		defer wg.Done()
		for {
			stop, ok := <-next
			if !ok {
				return
			}

			v, ok := <-ch
			if !ok {
				close(next)
				return
			}

			fmt.Println("goroutine01:", string(v))
			next <- stop
		}
	}()

	go func() {
		defer wg.Done()
		for {
			stop, ok := <-next
			if !ok {
				return
			}

			v, ok := <-ch
			if !ok {
				close(next)
				return
			}

			fmt.Println("goroutine02:", string(v))
			next <- stop
		}
	}()

	next <- struct{}{}

	wg.Wait()
}
