package main

import "fmt"

const f2mValue float64 = 0.3048

func main() {
	feet := 638.0
	fmt.Println(f2mValue * feet) // takes const and * by feet to give meters.

}
