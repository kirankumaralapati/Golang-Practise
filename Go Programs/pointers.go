package main

import("fmt")

func function(val *int){
	*val=25
}
func main(){
	b:=35
	fmt.Printf("\nType of b is %T\n", b)
	fmt.Println("\nLocation of a is", b)
	c:=&b
	function(c)
	fmt.Println("\nValue after changing is", b)
}