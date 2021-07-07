package main

import ("fmt"
"math")

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
	return math.Pi*c.radius*c.radius
}

func main(){
	r:= rectangle{
		lenght: 25.5,
		width: 35.5, 
	}
		fmt.Println("Area of Rectangle is", r.area())
	c:= circle{
		radius: 45.5,
	}
	fmt.Println("Area of Circle is", c.area())
}