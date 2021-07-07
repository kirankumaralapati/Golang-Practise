package main

import("fmt"
"time")

func senddata(c chan <-int){
	c <- 5
}

func main(){
	someval:= make(chan int)
	go senddata(someval)

	//someval2:= <-senddata
	time.Sleep(2* time.Second)
	fmt.Println(someval)
}