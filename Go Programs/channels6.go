package main

import("fmt"
"time")

func a(c chan int){
	for i:=0;i<=10;i++ {
		c<-i
	}
	close(c)
}

func main(){
	ch:= make(chan int)
	go a(ch)

	for v := range ch{
		time.Sleep(1*time.Second)
		fmt.Println("value Received is", v)
	}
	time.Sleep(1*time.Second)
	fmt.Println("Program Terminated")
}