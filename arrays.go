package main

import "fmt"

func main() {

	// var arrayName [size]elementType
	// fix size and not variable

	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Grapes", "Orange"}
	fmt.Println("Fruits array:", fruits)

	fmt.Println("Third element:", fruits[2])

	// originalArray := [3]int{1, 2, 3}
	// copiedArray := originalArray

	// copiedArray[0] = 100

	// fmt.Println("Original Array:", originalArray)
	// fmt.Println("Copied Array:", copiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index,", i, ":", numbers[i])
	}

	for index, value := range numbers { // for i, v := range numbers (preferred) (_ discard index)
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	//underscore is blank identifier, used to store unused values
	a, _ := someFunction()
	fmt.Println(a)
	//fmt.Println(b)

	// b := 2
	// _ = b

	// len() function
	fmt.Println("The length of numbers array is:", len(numbers))

	// Comparing Arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}

	fmt.Println("Are array1 and array2 equal?", array1 == array2)

	// Multidimensional Arrays
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:", matrix)

	// Pointer to an Array
	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int // declare pointer to an array
	copiedArray = &originalArray

	copiedArray[0] = 100

	fmt.Println("Original Array:", originalArray)
	//fmt.Println("Copied Array:", *copiedArray)
}

func someFunction() (int, int) {
	return 1, 2
}
