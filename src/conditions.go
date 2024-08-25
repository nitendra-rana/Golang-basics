package main

import (
	"fmt"
)

func main (){
	x := 10
	y := 10
	fmt.Println("if condition")
	if x> 5 {
		fmt.Println("greater than 10")
	}
	if y>20 {
		fmt.Println("greater than 10")
	}else{
		fmt.Println("less than 20")
	}

	switchMethod(3)
	 one, two := returnMethod(2,4)
	fmt.Println(one, two)
}

func switchMethod(day int ) {
  fmt.Println("The  switch case...")
	switch day  {
	case 1:
	  fmt.Println("Monday")
	case 2:
	  fmt.Println("Tuesday")
	case 3:
	  fmt.Println("Wednesday")
	case 4:
	  fmt.Println("Thursday")
	case 5:
	  fmt.Println("Friday")
	case 6:
	  fmt.Println("Saturday")
	case 7:
	  fmt.Println("Sunday")
	}
  }

  func returnMethod(a int, b int) (c int, d int) {
	fmt.Println("The return method:")
	c = a+b
	d = a-b
	return 
  }
