package main

import "fmt"

func main() {
	var x string = "eshays"     // string 1
	var y string = "adlay"      // string 2
	fmt.Println(x == y)         // string 1 != string 2, therefore: false
	var a string = "samestring" // string 1
	var b string = "samestring" // string 2
	fmt.Println(a == b)         // string 1 == string 2, therefore: true
}
