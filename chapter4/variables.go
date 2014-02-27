package main

import "fmt"

func main() {
	var x string = "The world says Hello!" // Pretty basic, x is a string based variable holding "The world says Hello!"
	var y string
	y = "Contents of y being defined later" // You don't need to define the contents of y when declaring the variable, it can be done later, like this.
	fmt.Println(x)                          // Printing a predefined variable
	fmt.Println(y)                          // Printing out a later defined variable
	var z string                            //Changing variables example with var string z
	z = "Initial String"
	fmt.Println(z)
	z = "Changed String"
	fmt.Println(z)
}
