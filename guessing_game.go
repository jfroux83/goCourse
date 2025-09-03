package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// Welcome message
	println("Welcome to the Guessing Game!")
	println("I have chosen a number between 1 and 100.")
	println("Can you guess what it is?")

	var guess int
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		// Check if the guess is correct
		if guess == target {
			fmt.Println("Congratulations! You've guessed the correct number!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing a higher number.")
		} else {
			fmt.Println("Too high! Try guessing a lower number.")
		}
	}
}
