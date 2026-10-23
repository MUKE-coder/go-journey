package main

import "fmt"

func main() {

	//ASSIGNMENT ONE
	// Challenge: Variables & Type Conversion
	// 1. Declare a string variable "name" with any name
	name := "Muke JB"
	// 2. Declare an int variable "age" with a number
	age := 30

	// 3. Declare a float64 variable "score" with a decimal number
	score := 10.5
	// 4. Declare a bool variable "passed" set to true
	passed := true
	// 5. Convert age to float64 and add it to score, store in "total"
	total := float64(age) + score
	// 6. Print each variable with its value and type using %v and %T

	fmt.Printf("The variable is Name and its val is %s and its type is %T \n", name, name)
	fmt.Printf("The variable is Age %d and its type is %T \n", age, age)
	fmt.Printf("The variable is Score %.2f and its type is %T \n", score, score)
	fmt.Printf("The variable is Passed %t and its type is %T \n", passed, passed)
	fmt.Printf("The variable is Total %.2f and its type is %T \n", total, total)

}
