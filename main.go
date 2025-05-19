package main //Every go file made, is going to be a part of a "package" which is just a collection of files and code. If the package is called "main", like in this example, this tells the Go compiler that our code should be compiled into an executable program at the end.
//That means that Go will choose to create a standalone executable program in case we choose to build the application and running that file will start the program, meaning that all the files we create for this program will be part of it and they will specify the main package at the start of the file.

import "fmt" // This is a package from the standard Go Library and this is how we import things in Go.

func main() { //All functions that we create will start with the "func" keyword. The name of the function plays an important role as it is the entry point of our application. So when we run our program, Go will look through our code and files to find this "main" function and will run it automatically.
	// We are only allowed to have ONE "main" function. In other files we can create other functions which can be invoked from inside the "Main" Function.
	fmt.Println("Hi! I hope you're doing great today.")
	// fmt is the package acronym and after the dot (.) we call this a method. The reason all methods start with a capital is because, inside these packages, in order to export something you have to use a capital for whatever you're exporting.
}
