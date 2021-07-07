package main

import ("fmt"
	"sync")

func numbers(wg *sync.WaitGroup){
	for i:=0;i<5;i++{
		fmt.Printf("\nNumber is %d",i)
	}
	defer wg.Done()
}

func alphabets(wg *sync.WaitGroup){
	for j:='a';j<'e';j++{
		fmt.Printf("\nLetter is %c",j)
	}
	defer wg.Done()
}

func execute(){
	wg:=new(sync.WaitGroup)
	wg.Add(2)
	go numbers(wg)
	go alphabets(wg)
	wg.Wait()
}

func main(){
	execute()
}
