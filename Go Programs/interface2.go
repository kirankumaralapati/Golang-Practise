package main

import ("fmt"
"math")

type geomentry interface{
	area() float64
	perimeter() float64
}

type rectangle struct{
	lenght float64
	width float64
}

type circle struct{
	radius float64
}

func (r rectangle) area() float64{
	return r.lenght * r.width
}
func (c circle) area() float64{
	return math.Pi* c.radius
}
func (r rectangle) perimeter() float64{
	return 2 * r.lenght * 2 * r.width 
}
func (c circle) perimeter() float64{
	return 2 * math.Pi * c.radius
}
func main(){

	var r geomentry

	var rect rectangle
	
}