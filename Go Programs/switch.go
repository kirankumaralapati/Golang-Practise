package main
import("fmt")

func number() int{
	num:=15*4
	return num
}

func main(){
	switch num:=number();{
	case num<50:
		fmt.Printf("number %d is inbetween 1 and 50",num)
		fallthrough
	case num<75:
			fmt.Printf("number %d is inbetween 50 and 75",num)
		fallthrough
	case num<=100:
		fmt.Printf("number %d is inbetween 75 and 100",num)
	}

}