package main

import "fmt"

var middleName = "Cane"

func main() {
	// var age int
	// var name string = "John"
	// var name1 = "Jane"

	// count := 10
	// lastName := "Smith"
	//middleName := "Cane"
	middleName := "Mayor"

	fmt.Println(middleName)

	// Default values:
	// Numeric Types: 0
	// Boolean Type: false
	// String Type: ""
	// Pointers, Slices, maps, functions, and structs: nil

	// ---- SCOPE
	// block scope
	// global scop: gopher syntax := only used in functions, use var outside functions

	// fmt.Println(firstName)

}

func printname() {
	firstName := "Michael"
	fmt.Println(firstName)
}
