package main
import ("fmt"
"time"
"sync")

func process(i int, wg *sync.WaitGroup){
	fmt.Printf("Goroutines Started is %d\n",i)
	time.Sleep(2*time.Second)
	fmt.Printf("\nGoroutines ended is %d\n",i)
	wg.Done()
}

func main(){
	no:= 5
	var wg sync.WaitGroup
	for i:=0;i<=no;i++{
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	time.Sleep(2*time.Second)
	fmt.Println("\nGoroutines Terminated")
}