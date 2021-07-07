package main

import (
	"fmt"
	"time"
)

func execute(some string) {

	for i := 1; true; i++ {

		fmt.Println(some, i)

		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	go execute("First")

	execute("Second")

	fmt.Println("program ends successfully")

}
