package main

import ("fmt")

func main(){
	a:= [...] int {1,2,3,4,5}

	for i:=0; i< len(a); i++ {
		fmt.Printf("\n%d th of the element is %d\n\n", i, a[i])
	}
}