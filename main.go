package main

import "fmt"

func main() {
	//strings
	// var nameOne string = "mario" //This is one of the ways to declare a variable

	// var nameTwo = "luigi" //In this type of declaration, Go is automatically inferring the type of variable, for us.

	// var nameThree string //In this one we set up the type of variable to use add a value in it in the future.

	// fmt.Println(nameOne, nameTwo, nameThree)
	// nameTwo = "peach"
	// nameThree = "bowser"

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour := "yoshi" //This is a shorthand. Another way to initialize or declare a variable for the first time, without the "var" keyword.
	// //Big Warning: You can't use the var shorthand ":=" outside of a function.

	// fmt.Println(nameFour)

	//ints

	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)
}
