package main

import "fmt"

func main() {
	var x [5]int      // x is an array, which has 5 ints
	x[4] = 100        // Mind blow, engineers count from 0 son!
	fmt.Println(x)    // Print entire array
	fmt.Println(x[4]) // Print 5th value of array
}
