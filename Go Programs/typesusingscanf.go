package main

import("fmt")

func main(){
	var(
		a int;
		b float64;
		c string;
		d string;
		sum string 
	)
	sum = c + d
	fmt.Scanln(&a,&b,&c)
	fmt.Println("a is", a,"b is",b,"c is", c, "d is", d,"sum",sum)
}