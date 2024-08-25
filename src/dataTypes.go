package main

import (
	"fmt"
)

func main(){
	fmt.Printf("Data Types =>")
	datatypes();
	fmt.Printf("Constants =>")
 	constants()
	 fmt.Printf("Formatting =>")
	formatting()

}
func datatypes(){
	// 	int- stores integers (whole numbers), such as 123 or -123
	// float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
	// string - stores text, such as "Hello World". String values are surrounded by double quotes
	// bool- stores values with two states: true or false
	//=> Go Variable Types
	// var variablename type = value
	var num int = 10 
	var flo float32 = 100.20
	var str string = "this is test string"
	var boolean bool = false
	test := 'A' //type is inferred
	//Multi var decelaration
	var a, b int = 1, 3
	 c, d := 7, "World!"
	fmt.Println("Values are =>" , num,flo, str, boolean,test);
	fmt.Println("multi var=> ", a, b, c, d)
	//Go variable naming rules:
		//1.  A variable name must start with a letter or an underscore character (_)
		// A variable name cannot start with a digit
		// A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
		// Variable names are case-sensitive (age, Age and AGE are three different variables)
		// There is no limit on the length of the variable name
		// A variable name cannot contain spaces
		// The variable name cannot be any Go keywords

		//Techniques to make variables more readable
		//var camelCaseVariableName
		//var PascalCaseVariableName
		//var  snake_case_variable_name 
}

func constants(){
	// const CONSTNAME type = value  (Decleration)
	const PI float32 = 3.14
	fmt.Printf("Pi =>" , PI);
}

func formatting(){
	// 	%v	Prints the value in the default format
	// %#v	Prints the value in Go-syntax format
	// %T	Prints the type of the value
	// %%	Prints the % sign
	const PI float32 = 3.14
	fmt.Printf("%v\n", PI)
	fmt.Printf("%#v\n", PI)
	fmt.Printf("%v%%\n", PI)
	fmt.Printf("%T\n", PI, "\n")
}