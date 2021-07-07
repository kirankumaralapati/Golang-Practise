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
	var i geomentry
	r:= rectangle{lenght: 3, width: 4}
	c:= circle{radius: 5}
	i = r
	fmt.Println(r)
	fmt.Println(c)
	
	fmt.Println(i.area())
	fmt.Println(i.perimeter())
}