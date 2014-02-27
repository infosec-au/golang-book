package main

import "fmt"

func main() {
	fmt.Println(true && true)  //True and True = True
	fmt.Println(true && false) //True and False = False
	fmt.Println(true || true)  // True or True? = True
	fmt.Println(true || false) // True or False? = True
	fmt.Println(!true)         // Not True = False
}
