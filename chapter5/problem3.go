package main

import "fmt"

const divByThree string = "Fizz"
const divByFive string = "Buzz"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(divByThree + divByFive)
		} else if i%3 == 0 {
			fmt.Println(divByThree)
		} else if i%5 == 0 {
			fmt.Println(divByFive)
		} else {
			fmt.Println(i)
		}
	}
}
