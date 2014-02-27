package main

import "fmt"

func main() {
	fmt.Println(len("Hello World")) // the len function counts the length of the str "Hello World"
	fmt.Println("Hello World"[0])   // the [0] selects the first byte of the str, usually one byte = one letter
	fmt.Println("Hello" + "World")  // Oh that's cool, you can concat strings just like you can in python.
}
