package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	counter := 0

	wg.Add(10)
	for i := 0; i < 10; i++ {

		go func() {
			for j := 0; j < 10; j++ {

				counter += 1
			}
			wg.Done()
		}()
	}
	fmt.Println("The number of goroutines before wait = ", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Counter value = ", counter)

	fmt.Println("The number of goroutines after wait = ", runtime.NumGoroutine())
}
