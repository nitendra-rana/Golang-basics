package main

import (
	"fmt"
)

func main () {
	arrayDecleartion()
	fmt.Println("Only specific elemet initialization.")
	initializeSpecificElem()
	fmt.Println("Length of an array.")
	getArrayLength()
}

func arrayDecleartion(){
	// 	var array_name = [length]datatype{values} // here length is defined
		var arrayWithLength = [3]int{1,2,3}
	// 	var array_name = [...]datatype{values} // here length is inferred
		var arrayWithoutLength =[...]bool{true, false, true}
	// array with inferred types
		// array_name := [length]datatype{values} // here length is defined
		// array_name := [...]datatype{values} // here length is inferred
		array_with_length := [3]string{"a", "b", "c"}
		fmt.Println("arrays =>", arrayWithLength,arrayWithoutLength,array_with_length)
}

func initializeSpecificElem(){
	// {INDEX:VALUE}
		arr1 := [5]int{1:10,2:40}
		fmt.Println(arr1)
}

func getArrayLength(){
	var arrayWithLength = [3]int{1,2,3}
	const LENGTH = len(arrayWithLength);
	fmt.Println(LENGTH)
}