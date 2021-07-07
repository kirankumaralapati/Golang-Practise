package main

import("fmt")

func slice(num[5] int){
	num[0]= 100
	fmt.Println("Num",num)
}

func main(){
	num:=[...]int{0,100,200,300,400}
	fmt.Println("Num2",num)
	slice(num)
	fmt.Println("Num3",num)
}