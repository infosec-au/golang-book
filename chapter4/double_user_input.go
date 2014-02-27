package main

import "fmt"

func main() {
	fmt.Print("Enter a number bro?: ") // Let's loosen the tension around these tutorials ;)
	var input float64                  // The expected input allowing for floats upto 64b
	fmt.Scanf("%f", &input)            // Scanf to obtain user input and assign to the input variable

	output := input * 2 // Multiply input by 2

	fmt.Println(output)
}
