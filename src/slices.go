package main

import (
	"fmt"
)

func main(){
	slicesDeclerations()
	lengthAndCapacity()
	copyMethod()
}

func slicesDeclerations(){
	//Slices are similar to arrays, but are more powerful and flexible.
	/*Like arrays, slices are also used to store multiple values of the same type in a single variable.
			However, unlike arrays, the length of a slice can grow and shrink as you see fit.
	*/
			// In Go, there are several ways to create a slice:
				// Using the []datatype{values} format
					//slice_name := []datatype{values}
					slice_name := []int{3, 2, 1}
					empty_slice := []int{}
				// Create a slice from an array
					// var myarray = [length]datatype{values} // An array
					// myslice := myarray[start:end] // A slice made from the array
					arr1 := [6]int{10, 11, 12, 13, 14,15}
					myslice := arr1[2:4]
				// Using the make() function
					//slice_name := make([]type, length, capacity)
					make_function_slice := make([]int, 6, 10)

				fmt.Println(slice_name, empty_slice, myslice, make_function_slice )
				make_function_slice[1] =10;
				fmt.Println(make_function_slice)

}

func lengthAndCapacity(){
	slice_name := []int{3, 2, 1}
	//len() function - returns the length of the slice (the number of elements in the slice)
		length := len(slice_name)
	//cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	    capacity  := cap(slice_name)
		fmt.Println("length =>", length, "Capicity =>",capacity)
}


func copyMethod(){
	/*
	Memory Efficiency
 	When using slices, Go loads all the underlying elements into the memory.
	If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.
	The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program. 
	*/
	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	println("before =>", numbersCopy, neededNumbers)
	copy(numbersCopy, neededNumbers)
	println("After => ", numbersCopy, neededNumbers)
	
}