package main

// 📍 For more information regarding the standard library, visit golang.org/pkg/#stdlib

import (
	"fmt"
	"sort"
)

func main() {
	//greeting := "Hi hi friend! Nice to see you here!"

	//fmt.Println(strings.Contains(greeting, "friend"))
	// 📍 In the above line we are using a method inside an importable package called "Contains" which searches in the variable "greeting" for the term "friend".
	// 📍 Since, it is in an fmt.Print method, it will actually return a "True or False" otherwise, boolean value.
	//fmt.Println(strings.ReplaceAll(greeting, "friend", "dearest"))
	// 📍 The method will look into the "greeting" variable to check for the word "friend" to replace it with the word "dearest". It doesn't alter the original string at all, it actually returns a new one.
	//fmt.Println(strings.ToUpper(greeting))
	// 📍 This method makes the string "greeting" all in uppercase. It also returns a new string and doesn't alter the original one.
	//fmt.Println(strings.Index(greeting, "Nice"))
	// 📍 This method returns the index position of the word "Nice" in the string called "greeting"

	//fmt.Println(strings.Split(greeting, " "))
	// 📍 This method splits the string's elements, so the words in the string or the the characters in the string are going to turn into an array and we can decide at what point we split the string up.
	// 📍 In this case the method will look into the "greeting" and wherever there's a space, it will split it so that the words of the "greeting" sentence will make up an array and the words will be the items in the array.

	// 📍 the original value is unchanged
	//fmt.Println("original string value =", greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25} //This is a slice of integers
	sort.Ints(ages)
	// 📍 The above method sorts a slice of integers. Specifically the one called within the method.
	// 📍 The method actually changes the original slice.
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	// 📍 The above method is going to search for the position of the index of number 30, IN THE SORTED Slice ,since after it's declaration it got sorted after a while.
	// 📍 In case we add an item that doesn't exist in the slice called "ages", the "sort.SearchInts" method will return positions that are not within the slice but are added to the actual slice after calling something that doesn't exist, but it doesn't sort them.

	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names) //Again this sorts the original slice of "names", it doesn't create a new one.
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser")) // This method searches the sorted alphabetical string slice of "names"
}
