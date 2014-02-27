Problems
=========

1. How are integers stored on a computer?
```
Integers are stored as binary on a computer.
```
2. We know that (in base 10) the largest 1 digit number is 9 and the largest 2 digit number is 99. Given that in binary the largest 2 digit number is 11 (3), the largest 3 digit number is 111 (7) and the largest 4 digit number is 1111 (15) what's the largest 8 digit number? (hint: 101-1 = 9 and 102-1 = 99) 
```
11111111 (binary) --> decimal = 255 (? had trouble understanding this Q)
```
3. Although overpowered for the task you can use Go as a calculator. Write a program that computes 32132 × 42452 and prints it to the terminal. (Use the * operator for multiplication)
```
package main

import "fmt"

func main(){
	fmt.Println(32132*42452)
}
```
4. What is a string? How do you find its length?
```
Strings are sequences of characters with a definite length to represent text. Go strings are made up of bytes (which is awesome).
Length of strings can be found by using the len function. e.g. fmt.Println(len("ThisStringIsLong"))
```
5. What's the value of the expression (true && false) || (false && true) || !(false && false)?
```
Simplified : false || false || true
therefore: true
```