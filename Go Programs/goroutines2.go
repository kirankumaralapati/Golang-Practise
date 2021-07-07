package main

import("fmt"
"time")

func numbers(){
	for i:=0; i<5;i++{
			time.Sleep(100*time.Millisecond)
			fmt.Printf("\nNumber is %d",i)
	}
}

func alphabets(){
	for j:='a';j<'e';j++{
		time.Sleep(200*time.Millisecond)
		fmt.Printf("\nLetter is %c",j)
	}
}

func main(){
	go numbers()
	go alphabets()
	time.Sleep(3000*time.Millisecond)
	fmt.Println("\nProgram Terminated")
}