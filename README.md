# go-dev-tutorial
## installation
https://go.dev/doc/install

	rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz
Do not untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.

Add /usr/local/go/bin to the PATH environment variable by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):
        
	export PATH=$PATH:/usr/local/go/bin

## basics
https://go.dev/doc/tutorial/getting-started

To enable dependency tracking for your code by creating a go.mod file, run the go mod init command, giving it the name of the module your code will be in. The name is the module's module path:

	go mod init example/hello
Run your code to see the greeting:

	go run .
Use the following command to get a list of the others:

	go help

## packages
Visit https://pkg.go.dev and search for a package.

Go will add the quote module as a requirement, as well as a go.sum file for use in authenticating the module:

	go mod tidy
Publish the example.com/greetings module from its repository (with a module path that reflected its published location), where Go tools could find it to download it. For now, because you haven't published the module yet, you need to adapt the example.com/hello module so it can find the example.com/greetings code on your local file system.

To do that, use the go mod edit command to edit the example.com/hello module to redirect Go tools from its module path (where the module isn't) to the local directory (where it is):

	go mod edit -replace example.com/greetings=../greetings

## tests
At the command line in the greetings directory, run the go test command to execute the test.

The go test command executes test functions (whose names begin with Test) in test files (whose names end with _test.go). You can add the -v flag to get verbose output that lists all of the tests and their results.

	go test -v

## compilation
The go build command compiles the packages, along with their dependencies, but it doesn't install the results.

The go install command compiles and installs the packages.

You can discover the install path by running the go list command, as in the following example:

	go list -f '{{.Target}}'
As an alternative, if you already have a directory like $HOME/bin in your shell path and you'd like to install your Go programs there, you can change the install target by setting the GOBIN variable using the go env command:

	go env -w GOBIN=/path/to/your/bin
or
	go env -w GOBIN=C:\path\to\your\bin
Once you've updated the shell path, run the go install command to compile and install the package:
	go install

## links

	https://gobyexample.com/
	https://go.dev/blog/declaration-syntax
	https://go.dev/blog/slices-intro
	https://go.dev/blog/defer-panic-and-recover
	https://go.dev/doc/code
	https://go.dev/doc/articles/wiki/
	https://go.dev/doc/effective_go

## Basic types

	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // alias for uint8

	rune // alias for int32
	// represents a Unicode code point

	float32 float64

	complex64 complex128
The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

Variables declared without an explicit initial value are given their zero value. The zero value is:

	0 for numeric types,
	false for the boolean type, and
	"" (the empty string) for strings.
## Type inference
When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type. But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:

	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
## Loops
The basic for loop has three components separated by semicolons:
* the init statement: executed before the first iteration
* the condition expression: evaluated before every iteration
* the post statement: executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to false. Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.
### Excercise
As a way to play with functions and loops, let's implement a square root function: given a number x, we want to find the number z for which z² is most nearly x. Computers typically compute the square root of x using a loop. Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:

	z -= (z*z - x) / (2*z)
Repeating this adjustment makes the guess better and better until we reach an answer that is as close to the actual square root as can be.
``` Go
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
}

func main() {
	fmt.Println(Sqrt(2))
}
```
A decent starting guess for z is 1, no matter what the input. To begin with, repeat the calculation 10 times and print each z along the way. See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves. Hint: To declare and initialize a floating point value, give it floating point syntax or use a conversion:
``` Go
z := 1.0
z := float64(1)
```
Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small amount). See if that's more or fewer than 10 iterations. Try other initial guesses for z, like x, or x/2. How close are your function's results to the math.Sqrt in the standard library?

The z² − x above is how far away z² is from where it needs to be (x), and the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing. This general approach is called Newton's method. It works well for many functions but especially well for square root.

## Defer
https://go.dev/blog/defer-panic-and-recover

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

## Pointers
The type *T is a pointer to a T value. Its zero value is nil.
``` Go
var p *int
```
The & operator generates a pointer to its operand.
``` Go
i := 42
p = &i
```
The * operator denotes the pointer's underlying value.
``` Go
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```
This is known as "dereferencing" or "indirecting". Unlike C, Go has no pointer arithmetic.

## Structs
A struct is a collection of fields.
``` Go
type Vertex struct {
	X int
	Y int
}
```
Struct fields can be accessed through a struct pointer. To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

The special prefix & returns a pointer to the struct value.
``` Go
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)
```

## Arrays/Slices
The type [n]T is an array of n values of type T.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
``` Go
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	primes := [6]int{2, 3, 5, 7, 11, 13}
```
An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type []T is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:
``` Go
a[low : high]
```
This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:
``` Go
a[1:4]
```
Slices are like references to arrays. A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array. Other slices that share the same underlying array will see those changes.
### Literals
This is an array literal:
``` Go
[3]bool{true, true, false}
```
And this creates the same array as above, then builds a slice that references it:
``` Go
[]bool{true, true, false}
```
---
When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the underlying slice or array for the high bound. These slice expressions are equivalent:
``` Go
a[0:10]
a[:10]
a[0:]
a[:]
```

A slice has both a length and a capacity. The length of a slice is the number of elements it contains. The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

The zero value of a slice is nil. A nil slice has a length and capacity of 0 and has no underlying array.
``` Go
	var s []int
	if s == nil {
		fmt.Println("nil!")
	}
```
---
Slices can be created with the built-in make function; this is how you create dynamically-sized arrays. The make function allocates a zeroed array and returns a slice that refers to that array:
``` Go
a := make([]int, 5)  // len(a)=5
```
To specify a capacity, pass a third argument to make:
``` Go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

Slices can contain any type, including other slices.
``` Go
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
```
It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append: https://go.dev/pkg/builtin/#append.
```
func append(s []T, vs ...T) []T
```
The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

Also see, https://go.dev/blog/slices-intro
### Range-based for
The range form of the for loop iterates over a slice or map. When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

You can skip the index or value by assigning to _. If you only want the index, you can omit the second variable.
### Exercise
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y. (You need to use a loop to allocate each []uint8 inside the [][]uint8. Use uint8(intValue) to convert between types.)

## Maps
A map maps keys to values. The zero value of a map is nil. A nil map has no keys, nor can keys be added. The make function returns a map of the given type, initialized and ready for use.

``` Go
	type Vertex struct {
		Lat, Long float64
	}

	var m map[string]Vertex

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
```
Map literals are like struct literals, but the keys are required.

If the top-level type is just a type name, you can omit it from the elements of the literal.
``` Go
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
```
### Mutating Maps
``` Go
// Insert or update an element in map m:
m[key] = elem
// Retrieve an element:
elem = m[key]
// Delete an element:
delete(m, key)
// Test that a key is present with a two-value assignment:
elem, ok = m[key]
```
If key is in m, ok is true. If not, ok is false.

If key is not in the map, then elem is the zero value for the map's element type. If elem or ok have not yet been declared you could use a short declaration form:
``` Go
elem, ok := m[key]
```
### Exercise
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

Also see, https://pkg.go.dev/strings#Fields
