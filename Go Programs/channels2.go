package main

import ("fmt")

func main(){
	var a chan int
	if a == nil {
		fmt.Println("Channel is Nill")
		a = make(chan int)
		fmt.Printf("Type of Channel is %T", a)
	}
}