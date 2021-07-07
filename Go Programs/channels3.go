package main

import ("fmt")

func practice(a chan int){
	a <- 2
}

func main(){
	a:= make(chan int)
	go practice(a)
	<- a
	fmt.Printf("\nType of Number is %T",a)
}