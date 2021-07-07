package main

import ("fmt")

type Address struct{
	city string
	state string
}
type person struct{
	FirstName string
	LastName string
	age int
	Phonenumber int64
	address Address
}

func main(){
	p:=person{
		FirstName: "Kirankumar",
		LastName: "Alapati",
		age: 24,
		//phonenumber: 573-275-6562,
		address: Address{
			city: "Dallas",
			state: "Texas",
		},
	}
	fmt.Println("\nFirstname is ", p.FirstName)
	fmt.Println("\nLastname is", p.LastName)
	fmt.Println("\nAge is", p.age)
	//fmt.Println("\nPhoneNumber is", p.PhoneNumber)
	fmt.Println("\nCity is", p.address.city)
	fmt.Println("\nState is", p.address.state)
}