# Tour of Go

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

## Function values
Functions are values too. They can be passed around just like other values. Function values may be used as function arguments and return values.
``` Go
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
```
### Closures
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
### Exercise
Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
``` Go
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

## Methods
Go does not have classes. However, you can define methods on types. A method is a function with a special receiver argument. The receiver appears in its own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v.
``` Go
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```
You can declare a method on non-struct types, too.

In this example we see a numeric type MyFloat with an Abs method.
``` Go
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
```
You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

You can declare methods with pointer receivers. This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)

For example, the Scale method here is defined on *Vertex.
``` Go
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
```
Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
### Methods and pointer indirection
Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer while methods with pointer receivers take either a value or a pointer as the receiver when they are called:
``` Go
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK

v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```
For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.

The equivalent thing happens in the reverse direction. Functions that take a value argument must take a value of that specific type while methods with value receivers take either a value or a pointer as the receiver when they are called:
``` Go
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!

fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```
In this case, the method call p.Abs() is interpreted as (*p).Abs().

There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

## Interfaces
An interface type is defined as a set of method signatures. A value of interface type can hold any value that implements those methods.

Interfaces are implemented implicitly.

A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.
### Interface values
Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

	(value, type)
An interface value holds a value of a specific underlying concrete type. Calling a method on an interface value executes the method of the same name on its underlying type.
#### Interface values with nil underlying values
If the concrete value inside the interface itself is nil, the method will be called with a nil receiver. In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)
``` Go
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

var i I
var t *T
i = t
describe(i)
i.M()
```
Note that an interface value that holds a nil concrete value is itself non-nil.
#### Nil interface values
A nil interface value holds neither value nor concrete type. Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
``` Go
	var i I
	describe(i)
	i.M()
	// panic: runtime error: invalid memory address or nil pointer dereference
```
###  Empty interfaces
The interface type that specifies zero methods is known as the empty interface:

	interface{}
An empty interface may hold values of any type. (Every type implements at least zero methods.) Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.
### Type assertions
A type assertion provides access to an interface value's underlying concrete value.

	t := i.(T)
This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t. If i does not hold a T, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

	t, ok := i.(T)
If i holds a T, then t will be the underlying value and ok will be true. If not, ok will be false and t will be the zero value of type T, and no panic occurs. Note the similarity between this syntax and that of reading from a map.

A type switch is a construct that permits several type assertions in series.
``` Go
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```
A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type. This switch statement tests whether the interface value i holds a value of type T or S. In each of the T and S cases, the variable v will be of type T or S respectively and hold the value held by i. In the default case (where there is no match), the variable v is of the same interface type and value as i.
### Stringers
One of the most ubiquitous interfaces is Stringer defined by the fmt package.
``` Go
type Stringer interface {
    String() string
}
```
A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
#### Exercise
Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad. For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
``` Go
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```
Solution:
``` Go
func (p IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", 
		p[0], p[1], p[2], p[3])
}
```

## Errors
Go programs express error state with error values. The error type is a built-in interface similar to fmt.Stringer:
``` Go
type error interface {
    Error() string
}
```
(As with fmt.Stringer, the fmt package looks for the error interface when printing values.) Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
``` Go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
```
A nil error denotes success; a non-nil error denotes failure.
### Exercise
Copy your Sqrt function from the earlier exercise and modify it to return an error value.

Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers. Create a new type

	type ErrNegativeSqrt float64
and make it an error by giving it a

	func (e ErrNegativeSqrt) Error() string
method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".

Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why? Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.
#### Solution
``` Go
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (z float64, e ErrNegativeSqrt) {
	if x < 0 {
		e = ErrNegativeSqrt(x)
		return 
	}
	z = 1.
	for y := x; ; {
		z -= (z*z - x) / (2 * z)
		if z == y || y-z <= 0.0000001 {
			return
		}
		y = z
	}
}
```

## Readers
The io package specifies the io.Reader interface, which represents the read end of a stream of data.

The Go standard library contains many [implementations](https://cs.opensource.google/search?q=Read%5C(%5Cw%2B%5Cs%5C%5B%5C%5Dbyte%5C)&ss=go%2Fgo) of this interface, including files, network connections, compressors, ciphers, and others. The io.Reader interface has a Read method:

	func (T) Read(b []byte) (n int, err error)
Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends. The example code creates a strings.Reader and consumes its output 8 bytes at a time.
``` Go
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
```
### Exercise: https://go.dev/tour/methods/22
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
#### Solution
https://stackoverflow.com/questions/27839140/tour-of-go-exercise-22-reader-what-does-the-question-mean
``` Go
type MyReader struct{}

func (m MyReader) Read(b []byte) (n int, e error) {
	if len(b) <= 0 {
		return
	}
	b[0] = 'A'
	return 1, nil
}
```
### Exercise
https://go.dev/tour/methods/23

A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters. The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
#### TODO
``` Go
var t map[byte]byte = table("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	"NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm")

func table(a string, b string) (m map[byte]byte) {
	m = make(map[byte]byte)
	for i, v := range a {
		m[byte(v)] = b[i]
	}
	return
}

func (r rot13Reader) Read(b []byte) (n int, e error) {
	if len(b) <= 0 {
		return
	}
	b[0] = t[b[0]]
	return 1, nil
}
```

## Images
Package image defines the Image interface:
``` Go
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
// the Rectangle return value of the Bounds method is actually an image.Rectangle, as the declaration is inside package image.
```
The color.Color and color.Model types are also interfaces, but we'll ignore that by using the predefined implementations color.RGBA and color.RGBAModel. These interfaces and types are specified by the image/color package.
### Exercise
Return an implementation of image.Image. Define your own Image type, implement the [necessary methods](https://pkg.go.dev/image#Image), and call pic.ShowImage.

Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

ColorModel should return color.RGBAModel.

At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.

https://stackoverflow.com/questions/39979956/golang-exerciseimages-missing-at-method

## Type parameters
Go functions can be written to work on multiple types using type parameters. The type parameters of a function appear between brackets, before the function's arguments.

	func Index[T comparable](s []T, x T) int
This declaration means that s is a slice of any type T that fulfills the built-in constraint comparable. x is also a value of the same type.

comparable is a useful constraint that makes it possible to use the == and != operators on values of the type.
``` Go
// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}
```
In addition to generic functions, Go also supports generic types. A type can be parameterized with a type parameter, which could be useful for implementing generic data structures.
``` Go
// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}
```

## Goroutines
A goroutine is a lightweight thread managed by the Go runtime.

	go f(x, y, z)
starts a new goroutine running f(x, y, z)

The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.
``` Go
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

go say("world")
```
Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives.
### Channels
Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
``` Go
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```
Like maps and slices, channels must be created before use:

	ch := make(chan int)
By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
``` Go
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

s := []int{7, 2, 8, -9, 4, 0}
c := make(chan int)

go sum(s[:len(s)/2], c)
go sum(s[len(s)/2:], c)

x, y := <-c, <-c // receive from c
```
The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.

Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

	ch := make(chan int, 100)
Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
### Range and Close
A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

	v, ok := <-ch
ok is false if there are no more values to receive and the channel is closed. The loop for i := range c receives values from the channel repeatedly until it is closed.
``` Go
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

c := make(chan int, 10)
go fibonacci(cap(c), c)
```
Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
### Select
The select statement lets a goroutine wait on multiple communication operations. A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
``` Go
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

c := make(chan int)
quit := make(chan int)
go func() {
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	quit <- 0
}()
fibonacci(c, quit)
```
#### Default Selection
The default case in a select is run if no other case is ready. Use a default case to try a send or receive without blocking:
``` Go
	start := time.Now()
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}
	for {
		select {
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())
		case <-boom:
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return
		default:
			fmt.Printf("[%6s]     .\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}
```
### Exercise: https://go.dev/tour/concurrency/8
This example uses the tree package, which defines the type:
``` Go
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
```
Use Go's concurrency and channels to write a simple solution: a function to check whether two binary trees store the same sequence.
#### Solution: https://medium.com/@basakabhijoy/solving-the-equivalent-binary-trees-exercise-in-go-92254cacfb76
``` Go
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Walking(t *tree.Tree, ch chan int) {
 	Walk(t, ch)
 	defer close(ch)  // close channel after Walk() finishes
}
// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) (b bool) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walking(t1, ch1)
	go Walking(t2, ch2)
	
	for {
		v1, e1 := <- ch1   
		v2, e2 := <- ch2
		if e1 != e2 || v1 != v2 {
			return false
		}
		if !e1 {
			break;
		}
	}
	return true
}
// fmt.Println(Same(tree.New(1), tree.New(1)))
```
### Mutex
We've seen how channels are great for communication among goroutines. But what if we don't need communication? What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts? This concept is called mutual exclusion, and the conventional name for the data structure that provides it is mutex. Go's standard library provides mutual exclusion with sync.Mutex and its two methods:
* Lock
* Unlock
We can define a block of code to be executed in mutual exclusion by surrounding it with a call to Lock and Unlock as shown on the Inc method. We can also use defer to ensure the mutex will be unlocked as in the Value method.
``` Go
// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
```
### Exercise: https://go.dev/tour/concurrency/10
Use Go's concurrency features to parallelize a web crawler. Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.
``` Go
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}
```
Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!
``` Go
// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
```
#### Solution
Also see, https://stackoverflow.com/questions/18207772/how-to-wait-for-all-goroutines-to-finish-without-using-time-sleep
``` Go
func Crawl(url string, depth int, fetcher Fetcher, ch chan string) {
	defer close(ch)
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		ch <- err.Error()
		return
	}
	ch <- fmt.Sprintf("found: %s %q\n", url, body)
	
	li := make([]chan string, len(urls))
	for i, u := range urls {
		li[i] := make(chan string)
		go Crawl(u, depth-1, fetcher, li[i])
	}
	
	for c := range li {
		for m := range c {
			ch <- m
		}
	}

	return
}

func main() {
	c := make(chan string)

	Crawl("https://golang.org/", 4, fetcher, c)
	for s := range c {
		fmt.Println(s)
	}
}
```
