package main
import ("fmt")

func main(){
	fmt.Println("Loops")
	forLoop(10)
}

func forLoop(flag int){
	for i := 0; i<flag; i++ {
		fmt.Println("loop =>",i)
	}
}