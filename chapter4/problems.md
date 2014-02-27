What are two ways to create a new variable?
```
1. var **name** **type** - e.g. var x string
2. variable_name := value - e.g. snailName := "Gary"
```
What is the value of x after running: x := 5; x += 1?
```
x = 6
the += adds to the original x variable - (5 + 1) = 6
```
What is scope and how do you determine the scope of a variable in Go?
```
Scope = what your go program can access, where
Globally accessible variables need to be put outside a function, usually before the first function
Non-global variables should be put into the function they are required in
```
What is the difference between var and const?
```
var means that the variables can be changed
const means that the variables will remain unchanged and cannot be changed
```
Using the example program as a starting point, write a program that converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)
```
package main

import "fmt"

func main() {
	fahrenheit := 85
	celsius := (fahrenheit - 32) * 5/9
	fmt.Println(celsius)
}
```
Write another program that converts from feet into meters. (1 ft = 0.3048 m)
```
package main

import "fmt"

const f2mValue float64 = 0.3048 // This is a float constant

func main() {
	feet := 638.0 // Need to make sure that Go recognises this as a float
	fmt.Println(f2mValue * feet) // takes const and * by feet to give meters.

}

```