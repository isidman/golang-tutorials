package main

import "fmt"

func main() {

	age := 35
	name := "Shaun"

	//Print
	// fmt.Print("hello, ")
	// fmt.Print("world! \n")

	//Println
	fmt.Println("Hello baby!")
	fmt.Println("Goodbye baby.")

	fmt.Println("My age is", age, "and my name is", name)

	// Printf(Formatted String)
	fmt.Printf("my age is %v and my name is %v \n", age, name) // %v for variables and order of variables matters a lot. %_ = format specifier
	fmt.Printf("my age is %q and my name is %q \n", age, name) //Ideally the %q is used only on strings. It adds double quotes around the variables
	fmt.Printf("age is of type %T \n", age)                    // The %T shows us the type of the variable we are calling
	fmt.Printf("you scored %f  points! \n", 255.55)            // The %f shows a float as it is.
	fmt.Printf("you scored %0.1f  points! \n", 255.55)         // We can limit the number of decimal points showing in the variable we've called. It rounds it automatically.

	//Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name) //It saves what we want to print as a formatted string in a variable and "fmt.Sprint()" saves the printf() in general.
	fmt.Println("the saved string is:", str)                              //It prints what we want along with the variable we have saved a while ago.

}
