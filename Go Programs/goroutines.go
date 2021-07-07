package main

import ("fmt"
"time")

func a (s string){
	for i:=0;i<3;i++{
		time.Sleep(1*time.Second)
		fmt.Println(s)
	}
}

func main(){
	go a("Kiran")
	a("Kumar")
}