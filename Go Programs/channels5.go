package main
import ("fmt"
"time")

func a(c chan int){
	for i:=0;i<=10;i++ {
		c <- i
	}
	close(c)
}

func main(){
	c:= make(chan int, 10)
	go a(c)

	for {
		v,ok:= <-c

		if ok == false {
			time.Sleep(2*time.Second)
			fmt.Println("\nCondition Terminated")
			break
		}
		time.Sleep(1*time.Second)
		fmt.Println("Channel Recived", v,ok)
	}
}
