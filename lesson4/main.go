package main

import "fmt"

func main() {
	// Arrays

	//	var ages [3]int = [3]int{20,25,30} // This says that I've made an array that is named "ages" that has 3 items (length of the array) of type 'integer'. So this is an array of 3 integers.
	// We add the specific items on the right side of the var array we have created, and since it's a set array we can not add something over the length of the array.
	//The array needs to always be a fixed length, it's not possible to change the length or the type of the items in an array once it has been declared.
	var ages = [3]int{20, 25, 30}
	//This is a shorthand way to create and set an array of 3 items of type integer, with 20,25 and thirty as the actual items in the array named "ages".

	names := [4]string{"yoshi", "mario", "peach", "bowser"} //Another shorthand way to create an array called names with 4 items of type string and their specified names in curly braces.
	names[1] = "luigi"                                      //In arrays we can change the item itself like this. We call the name of the array and within square brackets we set the position of the item we want to change.
	// We then add the equal sign and depending on the type of the item, we need to format it the same way as the other ones, otherwise we get an error as it is not possible to put different types of items in the same array, once it has been declared.

	fmt.Println(ages, len(ages)) //We can print the array itself as well as it's length like this.

	fmt.Println(names, len(names)) //Remember: In arrays, the length is always fixed!~

	//Slices (use arrays under the hood)

	var scores = []int{100, 50, 60}
	//When we don't place a number into the square brackets, a slice is created meaning that it doesn't have a fixed length that needs to stay unchanged once we declare it.
	scores[2] = 25 //In the same way we change the item in an array, we do it in a Slice, too!
	// __append(scores, 85)__//This function doesn't change the variable slice itself! Instead it returns a *NEW* Slice for us
	//Meaning that if we want to update the actual "scores" SIice we need to update it the same way we update items in a slice or an array.
	//So it will be:
	scores = append(scores, 85) //This is needed as the function doesn't append it automatically and reassignment must be done so as to not get an error!

	fmt.Println(scores, len(scores))

	//Slice ranges
	// A "range" is a way of getting a range of elements from an existing array or slice and store them in a *NEW* slice.
	rangeOne := names[1:3] // So when we declare a range for the first time we add the name of the array or the slice we want to make the range from.
	// Then in square brackets we can specify what we can include or exclude from the array or the slice we are creating the range from.
	// Specifically in range "rangeOne" we declare a range from the array called "names".
	// Specified in the square brackets we want to get in the range, the items from position 1 to position 3, while excluding the third item.
	// the colon always includes what is left of it in the range, as it starts from it and always excludes what is right of it, within the square brackets.

	rangeTwo := names[2:]   // When we set it like this within the square brackets, it means that go from the position two of the array, and take everything up until the end.
	rangeThree := names[:3] //In this setup, it means that the range should start from position 0 up until the position 3, while excluding it.

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "koopa") //since the rangeOne is also a slice we can append a new item within that slice.
	fmt.Println(rangeOne)
}
