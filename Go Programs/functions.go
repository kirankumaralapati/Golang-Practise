package main

import("fmt")

func CalculateBill(Price, Number float64) (float64 float64) {
	TotalPrice:= Price * Number
	return TotalPrice
}

func main(){
	Price:= 1555.5
	Number:= 95.5
	TotalPrice:= CalculateBill(Price, Number)
	fmt.Println("Total Price is", TotalPrice)
} 