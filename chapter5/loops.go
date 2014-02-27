package main

import "fmt"

func main() {
	// Simple loop

	i := 1        // assigning variable i as 1
	for i <= 10 { // from 1 - 10 (while i is smaller than or equal to 10)
		fmt.Println(i)
		i += 1
	}

	// More compact loop

	for b := 1; b <= 10; b++ { // sets variable b as 1, if b is less than or equal to 10, add 1
		fmt.Println(b)
	}

}
