package main

import "fmt"

type details interface{
	details()
}

type person struct{
	name string
	age int
}

func (p person) details(){
	fmt.Printf("%s is %d years old\n",p.name,p.age)
}

type address struct{
	city string
	state string
}

func (a *address) details(){
	fmt.Printf("City is %s State is %s\n",a.city,a.state)
}

func main(){
	var d1 details
	p1:= person{"Kirankumar", 24}
	d1 = p1
	d1.details()
	p2:= person{"Kiran Alapati", 24}
	d1=p2
	d1.details()
	
	var d2 details
	a1:= address{"Cape Girardeau","Missouri"}
	d1= &a1
	a2:= address{"Frisco", "Texas"}
	d2=&a2
	d1.details()
	d2.details()
}