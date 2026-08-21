# Go by Example
by Mark McGranaghan and Eli Bendersky, https://github.com/mmcgrana/gobyexample | https://gobyexample.com/
``` Go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```
## Values
Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.
``` Go
    fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
```
## Variables
In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.
``` Go
    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)
```
`var` declares 1 or more variables. Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an `int` is `0`.

The `:=` syntax is shorthand for declaring and initializing a variable, e.g. for `var f string = "apple"` in this case. This syntax is only available inside functions.
## Constants
Go supports constants of character, string, boolean, and numeric values.
``` Go
    fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))
```
`const` declares a constant value. A `const` statement can also appear inside a function body.

Constant expressions perform arithmetic with arbitrary precision. A numeric constant has no type until it’s given one, such as by an explicit conversion.

A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here `math.Sin` expects a `float64`.
## For
`for` is Go’s only looping construct. Here are some basic types of `for` loops.
``` Go
// The most basic type, with a single condition.
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }
// A classic initial/condition/after `for` loop.
    for j := 0; j < 3; j++ {
        fmt.Println(j)
    }
// `range` over an integer.
    for i := range 3 {
        fmt.Println("range", i)
    }
// `for` without a condition will loop repeatedly.
    for {
        fmt.Println("loop")
        break
    }

    for n := range 6 {
// You can also `continue` to the next iteration of the loop.
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
```
- The most basic type, with a single condition. 
- A classic initial/condition/after `for` loop.
- Another way of accomplishing the basic “do this N times” iteration is `range` over an integer.
- `for` without a condition will loop repeatedly until you break out of the loop or return from the enclosing function. You can also `continue` to the next iteration of the loop.
## If/Else
Branching with `if` and `else` in Go is straight-forward. Here’s a basic example.
``` Go
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    if 8%2 == 0 || 7%2 == 0 {
        fmt.Println("either 8 or 7 are even")
    }
// A statement can precede conditionals.
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
```
- You can have an `if` statement without an `else`.
- Logical operators like `&&` and `||` are often useful in conditions.
- A statement can precede conditionals; any variables declared in this statement are available in the current and all subsequent branches.

Note that you don’t need parentheses around conditions in Go, but that the braces are required. There is no ternary `if` in Go, so you’ll need to use a full `if` statement even for basic conditions.
## Switch
Switch statements express conditionals across many branches.
``` Go
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }
// You can use commas to separate multiple expressions in the same case statement. 
// We use the optional default case in this example as well.
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }
// switch without an expression is an alternate way to express if/else logic. 
    t := time.Now()
    switch {
// Here we also show how the case expressions can be non-constants.
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }
// In this example, the variable t will have the type corresponding to its clause.
    whatAmI := func(i any) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
```
You can use commas to separate multiple expressions in the same case statement. We use the optional `default` case in this example as well.

switch without an expression is an alternate way to express if/else logic. Here we also show how the case expressions can be non-constants.

A type switch compares types instead of values. You can use this to discover the type of an interface value. In this example, the variable `t` will have the type corresponding to its clause.
## Arrays
In Go, an array is a numbered sequence of elements of a specific length. In typical Go code, slices are much more common; arrays are useful in some special scenarios.

Here we create an array `a` that will hold exactly 5 `int`s. The type of elements and length are both part of the array’s type. By default an array is zero-valued, which for `int`s means `0`s.
``` Go
    var a [5]int
    fmt.Println("emp:", a)
// We can set a value at an index using the array[index] = value syntax, and get a value with array[index].
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
// The builtin len returns the length of an array.
    fmt.Println("len:", len(a))
// Use this syntax to declare and initialize an array in one line.
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
// You can also have the compiler count the number of elements for you with ...
    b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
// If you specify the index with :, the elements in between will be zeroed.
    b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)
// Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
    var twoD [2][3]int
    for i := range 2 {
        for j := range 3 {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
// You can create and initialize multi-dimensional arrays at once too.
    twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }
    fmt.Println("2d: ", twoD)
```
We can set a value at an index using the `array[index] = value` syntax, and get a value with `array[index]`.

The builtin `len` returns the length of an array.

You can also have the compiler count the number of elements for you with `...`

If you specify the index with `:`, the elements in between will be zeroed.

Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
## Slices
Slices are an important data type in Go, giving a more powerful interface to sequences than arrays. Unlike arrays, slices are typed only by the elements they contain (not the number of elements). An uninitialized slice equals to `nil` and has length `0`.
``` Go
    var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0)

    s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
// We can set and get just like with arrays.
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
// len returns the length of the slice as expected.
    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)
// Here we create an empty slice c of the same length as s and copy into c from s.
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)
// This gets a slice of the elements s[2], s[3], and s[4].
    l := s[2:5]
    fmt.Println("sl1:", l)
// This slices up to (but excluding) s[5].
    l = s[:5]
    fmt.Println("sl2:", l)
// And this slices up from (and including) s[2].
    l = s[2:]
    fmt.Println("sl3:", l)
// We can declare and initialize a variable for slice in a single line as well.
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }
// Slices can be composed into multi-dimensional data structures. 
// The length of the inner slices can vary, unlike with multi-dimensional arrays.
    twoD := make([][]int, 3)
    for i := range 3 {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := range innerLen {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
```
To create a slice with non-zero length, use the builtin `make`. Here we make a slice of strings of length 3 (initially zero-valued). By default a new slice’s capacity is equal to its length; if we know the slice is going to grow ahead of time, it’s possible to pass a capacity explicitly as an additional parameter to `make`. 

We can `set` and `get` just like with arrays.

`len` returns the length of the slice as expected.

In addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin `append`, which returns a slice containing one or more new values. Note that we need to accept a return value from `append` as we may get a new slice value.

Slices can also be `copy`’d. Here we create an empty slice `c` of the same length as `s` and copy into `c` from `s`.

Slices support a “slice” operator with the syntax `slice[low:high]`. For example, this gets a slice of the elements `s[2]`, `s[3]`, and `s[4]`. This slices up to (but excluding) `s[5]`. And this slices up from (and including) `s[2]`.

We can declare and initialize a variable for slice in a single line as well. The slices package contains a number of useful utility functions for slices.

Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays. Also see [blog post](https://go.dev/blog/slices-intro).
## Maps
Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages). To create an empty map, use the builtin make: `make(map[key-type]val-type)`.
``` Go
    m := make(map[string]int)
// Set key/value pairs using typical name[key] = val syntax.
    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)
// Get a value for a key with name[key].
    v1 := m["k1"]
    fmt.Println("v1:", v1)

    v3 := m["k3"]
    fmt.Println("v3:", v3)
// The builtin len returns the number of key/value pairs when called on a map.
    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)
// To remove all key/value pairs from a map, use the clear builtin.
    clear(m)
    fmt.Println("map:", m)
// Here we didn’t need the value itself, so we ignored it with the blank identifier _.
    _, prs := m["k2"]
    fmt.Println("prs:", prs)
// You can also declare and initialize a new map in the same line with this syntax.
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

    n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
```
- To create an empty map, use the builtin make: `make(map[key-type]val-type)`.
- Set key/value pairs using typical `name[key] = val` syntax.
- Printing a map with e.g. `fmt.Println` will show all of its key/value pairs.
- Get a value for a key with `name[key]`.
- If the key doesn’t exist, the zero value of the value type is returned.
- The builtin `len` returns the number of key/value pairs when called on a map.
- The builtin `delete` removes key/value pairs from a map.
- To remove all key/value pairs from a map, use the `clear` builtin.

The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like `0` or `""`. Here we didn’t need the value itself, so we ignored it with the blank identifier `_`.

You can also declare and initialize a new map in the same line.
## Functions
Functions are central in Go. We’ll learn about functions with a few different examples. Here’s a function that takes two `int`s and returns their sum as an `int`.
``` Go
func plus(a int, b int) int {
// Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
    return a + b
}
// You may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func plusPlus(a, b, c int) int {
    return a + b + c
}
```
Go requires explicit returns, i.e. it won’t automatically return the value of the last expression. When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type. Call a function just as you’d expect, with `name(args)`.
``` Go
    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
```
### Multiple Return Values
Go has built-in support for multiple return values. This feature is used often in idiomatic Go, for example to return both result and error values from a function.
``` Go
// The (int, int) in this function signature shows that the function returns 2 ints.
func vals() (int, int) {
    return 3, 7
}

func main() {
// Here we use the 2 different return values from the call with multiple assignment.
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)
// If you only want a subset of the returned values, use the blank identifier _.
    _, c := vals()
    fmt.Println(c)
}
```
Accepting a variable number of arguments is another nice feature of Go functions; we’ll look at this next.
### Variadic Functions
Variadic functions can be called with any number of trailing arguments. For example, `fmt.Println` is a common variadic function.

Here’s a function that will take an arbitrary number of `int`s as arguments.
``` Go
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
// The type of nums is equivalent to []int. We can call len(nums), iterate over it with range, etc.
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
```
Variadic functions can be called in the usual way with individual arguments.
``` Go
    sum(1, 2)
    sum(1, 2, 3)
// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
    nums := []int{1, 2, 3, 4}
    sum(nums...)

```
Another key aspect of functions in Go is their ability to form closures, which we’ll look at next.
### Closures
Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.

This function `intSeq` returns another function, which we define anonymously in the body of `intSeq`. The returned function closes over the variable `i` to form a closure.
``` Go
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
```
We call `intSeq`, assigning the result (a function) to `nextInt`. This function value captures its own `i` value, which will be updated each time we call `nextInt`.
``` Go
    nextInt := intSeq()
// See the effect of the closure by calling nextInt a few times.
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
// To confirm that the state is unique to that particular function, create and test a new one.
    newInts := intSeq()
    fmt.Println(newInts())
```
### Recursion
This `fact` function calls itself until it reaches the base case of `fact(0)`.
``` Go
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}
```
Anonymous functions can also be recursive, but this requires explicitly declaring a variable with `var` to store the function before it’s defined.
``` Go
    var fib func(n int) int

    fib = func(n int) int {
        if n < 2 {
            return n
        }
// Since fib was previously declared in main, Go knows which function to call with fib here.
        return fib(n-1) + fib(n-2)
    }

    fmt.Println(fib(7))
```
## Range over Built-in Types
`range` iterates over elements in a variety of built-in data structures. Let’s see how to use range with some of the data structures we’ve already learned.

Here we use `range` to sum the numbers in a slice. Arrays work like this too.
``` Go
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
```
`range` on arrays and slices provides both the index and value for each entry. Above we didn’t need the index, so we ignored it with the blank identifier `_`. Sometimes we actually want the indexes though.
``` Go
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
// range on map iterates over key/value pairs.
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
// range can also iterate over just the keys of a map.
    for k := range kvs {
        fmt.Println("key:", k)
    }
// range on strings iterates over Unicode code points. 
    for i, c := range "go" {
        fmt.Println(i, c)
    }
```
The first value is the starting byte index of the rune and the second the rune itself.
## Pointers
Go supports pointers, allowing you to pass references to values and records within your program.

We’ll show how pointers work in contrast to values with 2 functions: `zeroval` and `zeroptr`. `zeroval` has an `int` parameter, so arguments will be passed to it by value. `zeroval` will get a copy of `ival` distinct from the one in the calling function.
``` Go
func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}
```
`zeroptr` in contrast has an `*int` parameter, meaning that it takes an `int` pointer. The `*iptr` code in the function body then dereferences the pointer from its memory address to the current value at that address. Assigning a value to a dereferenced pointer changes the value at the referenced address.
``` Go
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)
// The &i syntax gives the memory address of i, i.e. a pointer to i.
    zeroptr(&i)
    fmt.Println("zeroptr:", i)
// Pointers can be printed too.
    fmt.Println("pointer:", &i)
// A new pointer to a value can be created with the builtin function new.
    p := new(42)
    fmt.Println("value at *p:", *p)
    zeroptr(p)
    fmt.Println("value at *p:", *p)
```
- The &i syntax gives the memory address of i, i.e. a pointer to i.
- A new pointer to a value can be created with the builtin function new.

`zeroval` doesn’t change the `i` in main, but `zeroptr` does because it has a reference to the memory address for that variable.
## Strings and Runes
A Go string is a read-only slice of bytes. The language and the standard library treat strings specially - as containers of text encoded in UTF-8. In other languages, strings are made of “characters”. In Go, the concept of a character is called a `rune` - it’s an integer that represents a Unicode code point. This [Go blog post](https://go.dev/blog/strings) is a good introduction to the topic.

`s` is a string assigned a literal value representing the word `“hello”` in the Thai language. Go string literals are UTF-8 encoded text.
``` Go
    const s = "สวัสดี"
// Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.
    fmt.Println("Len:", len(s))
// This loop generates the hex values of all the bytes that constitute the code points in s.
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()
```
Indexing into a string produces the raw byte values at each index. 

To count how many runes are in a string, we can use the `utf8` package. Note that the run-time of `RuneCountInString` depends on the size of the string, because it has to decode each UTF-8 rune sequentially. Some Thai characters are represented by UTF-8 code points that can span multiple bytes, so the result of this count may be surprising.

A `range` loop handles strings specially and decodes each `rune` along with its offset in the `string`.
``` Go
    fmt.Println("Rune count:", utf8.RuneCountInString(s))

    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }
// We can achieve the same iteration by using the utf8.DecodeRuneInString function explicitly.
    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width
// This demonstrates passing a rune value to a function.
        examineRune(runeValue)
    }
```
Values enclosed in single quotes are `rune` literals. We can compare a `rune` value to a `rune` literal directly.
``` Go
func examineRune(r rune) {
// We can compare a rune value to a rune literal directly.
    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}
```
## Structs
Go’s `struct`s are typed collections of fields. They’re useful for grouping data together to form records. This person `struct` type has name and age fields.
``` Go
type person struct {
    name string
    age  int
}

func newPerson(name string) *person {
// newPerson constructs a new person struct with the given name.
    p := person{name: name}
    p.age = 42
    return &p
}
```
Go is a garbage collected language; you can safely return a pointer to a local variable - it will only be cleaned up by the garbage collector when there are no active references to it.
``` Go
// This syntax creates a new struct.
    fmt.Println(person{"Bob", 20})
// You can name the fields when initializing a struct.
    fmt.Println(person{name: "Alice", age: 30})
// Omitted fields will be zero-valued.
    fmt.Println(person{name: "Fred"})
// An & prefix yields a pointer to the struct.
    fmt.Println(&person{name: "Ann", age: 40})
// It’s idiomatic to encapsulate new struct creation in constructor functions
    fmt.Println(newPerson("Jon"))
// Access struct fields with a dot.
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)
// You can also use dots with struct pointers - the pointers are automatically dereferenced.
    sp := &s
    fmt.Println(sp.age)
// Structs are mutable.
    sp.age = 51
    fmt.Println(sp.age)
```
If a `struct` type is only used for a single value, we don’t have to give it a name. The value can have an anonymous struct type. This technique is commonly used for [table-driven tests](https://gobyexample.com/testing-and-benchmarking).
``` Go
    dog := struct {
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }
    fmt.Println(dog)
```
## Methods
Go supports methods defined on `struct` types.
``` Go
type rect struct {
    width, height int
}
// This area method has a receiver type of *rect.
func (r *rect) area() int {
    return r.width * r.height
}
// Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}
```
Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving `struct`.
``` Go
    r := rect{width: 10, height: 5}
// Here we call the 2 methods defined for our struct.
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
```
## Interfaces
Interfaces are named collections of method signatures. Here’s a basic `interface` for geometric shapes.
``` Go
type geometry interface {
    area() float64
    perim() float64
}
```
For our example we’ll implement this interface on `rect` and `circle` types.
``` Go
type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}
```
To implement an `interface` in Go, we just need to implement all the methods in the `interface`. Here we implement geometry on `rect`s.
``` Go
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}
// The implementation for circles.
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}
```
If a variable has an `interface` type, then we can call methods that are in the named `interface`. Here’s a generic measure function taking advantage of this to work on any geometry.
``` Go
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}
```
Sometimes it’s useful to know the runtime type of an `interface` value. One option is using a type assertion as shown here; another is a type `switch`.
``` Go
func detectCircle(g geometry) {
    if c, ok := g.(circle); ok {
        fmt.Println("circle with radius", c.radius)
    }
}
```
The `circle` and `rect` `struct` types both implement the geometry `interface` so we can use instances of these `struct`s as arguments to measure.
``` Go
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
    detectCircle(r)
    detectCircle(c)
```
To understand how Go’s interfaces work under the hood, check out this [blog post](https://research.swtch.com/interfaces).
## Enums
Enumerated types (enums) are a special case of sum types. An enum is a type that has a fixed number of possible values, each with a distinct name. Go *doesn’t have an enum type* as a distinct language feature, but enums are simple to implement using existing language idioms.

Our enum type `ServerState` has an underlying `int` type.
``` Go
type ServerState int

const (
    StateIdle ServerState = iota
    StateConnected
    StateError
    StateRetrying
)
```
The possible values for `ServerState` are defined as constants. The special keyword `iota` generates successive constant values automatically; in this case `0, 1, 2` and so on.

By implementing the `fmt.Stringer` [interface](https://pkg.go.dev/fmt#Stringer), values of `ServerState` can be printed out or converted to strings.

This can get cumbersome if there are many possible values. In such cases the `stringer` [tool](https://pkg.go.dev/golang.org/x/tools/cmd/stringer) can be used in conjunction with `go:generate` to automate the process. See [this post](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate) for a longer explanation.
``` Go
var stateName = map[ServerState]string{
    StateIdle:      "idle",
    StateConnected: "connected",
    StateError:     "error",
    StateRetrying:  "retrying",
}
	
func (ss ServerState) String() string {
    return stateName[ss]
}
// transition emulates a state transition for a server; it takes the existing state and returns a new state.
func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:
// Suppose we check some predicates here to determine the next state…
        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}
```
If we have a value of type `int`, we cannot pass it to transition - the compiler will complain about type mismatch. This provides some degree of compile-time *type safety* for enums.
``` Go
    ns := transition(StateIdle)
    fmt.Println(ns)
    ns2 := transition(ns)
    fmt.Println(ns2)
```
## Struct Embedding
Go supports embedding of structs and interfaces to express a more seamless *composition of types*. 

This is not to be confused with [`//go:embed`](https://gobyexample.com/embed-directive) which is a go directive introduced in Go version 1.16+ to embed files and folders into the application binary.
``` Go
type base struct {
    num int
}

func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}
// A container embeds a base. An embedding looks like a field without a name.
type container struct {
    base
    str string
}
```
When creating structs with literals, we have to initialize the embedding explicitly; here the embedded type serves as the field name.
``` Go
    co := container{
        base: base{
            num: 1,
        },
        str: "some name",
    }
// We can access the base’s fields directly on co, e.g. co.num.
    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
// Alternatively, we can spell out the full path using the embedded type name.
    fmt.Println("also num:", co.base.num)
// Here we invoke a method that was embedded from base directly on co.
    fmt.Println("describe:", co.describe())
```
Since `container` embeds `base`, the methods of `base` also become methods of a `container`. 

Embedding structs with methods may be used to bestow `interface` implementations onto other structs. Here we see that a `container` now implements the `describer` interface because it embeds `base`.
``` Go
    type describer interface {
        describe() string
    }

    var d describer = co
    fmt.Println("describer:", d.describe())
```
## Generics
Starting with version 1.18, Go has added support for generics, also known as type parameters.

As an example of a generic function, `SlicesIndex` takes a slice of any comparable type and an element of that type and returns the index of the first occurrence of `v` in `s`, or `-1` if not present. The comparable constraint means that we can compare values of this type with the `==` and `!=` operators. For a more thorough explanation of this type signature, see this [blog post](https://go.dev/blog/deconstructing-type-parameters). Note that this function exists in the standard library as [`slices.Index`](https://pkg.go.dev/slices#Index).
``` Go
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
    for i := range s {
        if v == s[i] {
            return i
        }
    }
    return -1
}
```
As an example of a generic type, `List` is a singly-linked list with values of any type.
``` Go
type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val  T
}
```
We can define methods on generic types just like we do on regular types, but we have to keep the type parameters in place. The type is `List[T]`, not `List`.
``` Go
func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}
```
`AllElements` returns all the `List` elements as a slice. 
``` Go
func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}
```
In the next example we’ll see a more idiomatic way of iterating over all elements of custom types.

When invoking generic functions, we can often rely on type inference. Note that we don’t have to specify the types for `S` and `E` when calling `SlicesIndex` - the compiler infers them automatically.
``` Go
    var s = []string{"foo", "bar", "zoo"}

    fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))
// … though we could also specify them explicitly.
    _ = SlicesIndex[[]string, string](s, "zoo")

    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.AllElements())
```
## Range over Iterators
Starting with version 1.23, Go has added support for [iterators](https://go.dev/blog/range-functions), which lets us range over pretty much anything!

Let’s look at the `List` type from the previous example again. In that example we had an `AllElements` method that returned a slice of all elements in the list. With Go iterators, we can do it better - as shown below.
``` Go
type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val  T
}

func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) All() iter.Seq[T] {
    return func(yield func(T) bool) {

        for e := lst.head; e != nil; e = e.next {
            if !yield(e.val) {
                return
            }
        }
    }
}
```
`All` returns an iterator, which in Go is a *function with a special signature*.

The `iterator` function takes another function as a parameter, called `yield` by convention (but the name can be arbitrary). It will call `yield` for every element we want to iterate over, and note `yield`’s return value for a potential early termination.
``` Go
func genFib() iter.Seq[int] {
    return func(yield func(int) bool) {
        a, b := 0, 1

        for {
            if !yield(a) {
                return
            }
            a, b = b, a+b
        }
    }
}
```
Iteration doesn’t require an underlying data structure, and doesn’t even have to be finite! Here’s a function returning an `iterator` over Fibonacci numbers: it keeps running as long as `yield` keeps returning true.
``` Go
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
// Since List.All returns an iterator, we can use it in a regular range loop.
    for e := range lst.All() {
        fmt.Println(e)
    }
// Collect takes any iterator and collects all its values into a slice.
    all := slices.Collect(lst.All())
    fmt.Println("all:", all)
// strings.SplitSeq iterates over parts of a byte slice without first building a result slice.
    for part := range strings.SplitSeq("go-by-example", "-") {
        fmt.Printf("part: %s\n", part)
    }

    for n := range genFib() {
// Once the loop hits break or an early return, the yield function passed to the iterator will return false.
        if n >= 10 {
            break
        }
        fmt.Println(n)
    }

```
- Since List.All returns an iterator, we can use it in a regular range loop.

Packages like slices have a number of useful functions to work with iterators. For example, `Collect` takes any `iterator` and collects all its values into a slice.

Standard library packages now expose `iterator` helpers too. For example, `strings.SplitSeq` iterates over parts of a `byte` slice without first building a result slice.

Once the loop hits `break` or an early return, the `yield` function passed to the `iterator` will return `false`.
## Errors
In Go it’s idiomatic to communicate errors via an explicit, separate return value. This contrasts with the exceptions used in languages like Java, Python and Ruby and the overloaded single result / error value sometimes used in C. Go’s approach makes it easy to see which functions return `error`s and to handle them using the same language constructs employed for other, non-`error` tasks.

Also see [error pack](https://pkg.go.dev/errors), and [blog post](https://go.dev/blog/go1.13-errors).

By convention, errors are the last return value and have type `error`, a built-in `interface`.
- `errors.New` constructs a basic `error` value with the given error message.
- A `nil` value in the error position indicates that there was no error.
- A *sentinel error* is a predeclared variable that is used to signify a specific error condition.
``` Go
func f(arg int) (int, error) {
    if arg == 42 {
// errors.New constructs a basic error value with the given error message.
        return -1, errors.New("can't work with 42")
    }
// A nil value in the error position indicates that there was no error.
    return arg + 3, nil
}
// A sentinel error is a predeclared variable that is used to signify a specific error condition.
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

func makeTea(arg int) error {
    if arg == 2 {
        return ErrOutOfTea
    } else if arg == 4 {

        return fmt.Errorf("making tea: %w", ErrPower)
    }
    return nil
}
```
We can wrap errors with higher-level errors to add context. The simplest way to do this is with the `%w` verb in `fmt.Errorf`. Wrapped errors create a logical chain (A wraps B, which wraps C, etc.) that can be queried with functions like `errors.Is` and `errors.AsType`.
``` Go
    for _, i := range []int{7, 42} {
// It’s idiomatic to use an inline error check in the if line.
        if r, e := f(i); e != nil {
            fmt.Println("f failed:", e)
        } else {
            fmt.Println("f worked:", r)
        }
    }

    for i := range 5 {
        if err := makeTea(i); err != nil {
// errors.Is checks that a given error (or any error in its chain) matches a specific error value.
            if errors.Is(err, ErrOutOfTea) {
                fmt.Println("We should buy new tea!")
            } else if errors.Is(err, ErrPower) {
                fmt.Println("Now it is dark.")
            } else {
                fmt.Printf("unknown error: %s\n", err)
            }
            continue
        }
        fmt.Println("Tea is ready!")
    }
```
`errors.Is` checks that a given error (or any `error` in its chain) matches a specific error value. This is especially useful with wrapped or nested errors, allowing you to identify specific error types or sentinel errors in a chain of errors.
### Custom Errors
It’s possible to define custom error types by implementing the `Error()` method on them. Here’s a variant on the example above that uses a custom type to explicitly represent an argument `error`.
``` Go
// A custom error type usually has the suffix “Error”.
type argError struct {
    arg     int
    message string
}
// Adding this Error method makes argError implement the error interface.
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
    if arg == 42 {
// Return our custom error.
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}
```
`errors.AsType` is a more advanced version of `errors.Is`. It checks that a given error (or any `error` in its chain) matches a specific `error` type and converts to a value of that type, also returning `true`. If there’s no match, the second return value is `false`.
``` Go
    _, err := f(42)
    if ae, ok := errors.AsType[*argError](err); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.message)
    } else {
        fmt.Println("err doesn't match argError")
    }
```
## Goroutines
A *goroutine is a lightweight thread of execution*.
``` Go
func f(from string) {
    for i := range 3 {
        fmt.Println(from, ":", i)
    }
}
```
Suppose we have a function call `f(s)`. Here’s how we’d call that in the usual way, running it synchronously.

To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.
``` Go
// Here’s how we’d call that in the usual way, running it synchronously.
    f("direct")
// This new goroutine will execute concurrently with the calling one.
    go f("goroutine")
// You can also start a goroutine for an anonymous function call.
    go func(msg string) {
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)
    fmt.Println("done")
```
Our two function calls are running asynchronously in separate goroutines now. Wait for them to finish (for a more robust approach, use a [`WaitGroup`](https://gobyexample.com/waitgroups)).

When we run this program, we see the output of the blocking call first, then the output of the two goroutines. The goroutines’ output may be interleaved, because goroutines are being run concurrently by the Go runtime.
## Channels
Channels are the *pipes that connect concurrent goroutines*. You can send values into channels from one goroutine and receive those values into another goroutine.

Create a new channel with `make(chan val-type)`. Channels are typed by the values they convey.
``` Go
    messages := make(chan string)
// Here we send "ping" to the messages channel we made above, from a new goroutine.
    go func() { messages <- "ping" }()
// Here we send "ping" to the messages channel we made above, from a new goroutine.
    msg := <-messages
    fmt.Println(msg)
```
- Send a value into a channel using the channel `<-` syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.
- The `<-channel` syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
- When we run the program the "ping" message is successfully passed from one goroutine to another via our channel.

By default sends and receives block until both the sender and receiver are ready. This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.
### Channel Buffering
By default channels are unbuffered, meaning that they will only accept sends `(chan <-)` if there is a corresponding receive `(<- chan)` ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.
``` Go
// Here we make a channel of strings buffering up to 2 values.
    messages := make(chan string, 2)
// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
    messages <- "buffered"
    messages <- "channel"
// Later we can receive these two values as usual.
    fmt.Println(<-messages)
    fmt.Println(<-messages)
```
- Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive. Later we can receive these two values as usual.
### Channel Synchronization
We can use channels to synchronize execution across goroutines. Here’s an example of using a blocking receive to wait for a goroutine to finish. When waiting for multiple goroutines to finish, you may prefer to use a `WaitGroup`.

This is the function we’ll run in a goroutine. The `done` channel will be used to notify another goroutine that this function’s work is done.
``` Go
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")
// Send a value to notify that we’re done.
    done <- true
}
```
Start a worker goroutine, giving it the channel to notify on.
``` Go
    done := make(chan bool, 1)
    go worker(done)
// Block until we receive a notification from the worker on the channel.
    <-done
```
If you removed the `<- done` line from this program, the program could exit before the worker finished its work, or in some cases even before it started.
### Channel Directions
When using channels as function parameters, you can specify if a channel is meant to only send or receive values. This specificity increases the *type-safety* of the program.

This ping function only accepts a channel for sending values. It would be a compile-time error to try to receive on this channel.
``` Go
func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}
```
The `pong` function accepts one channel for receives (`pings`) and a second for sends (`pongs`).
``` Go
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
```
### Select
Go’s `select` lets you wait on multiple channel operations. Combining goroutines and channels with `select` is a powerful feature of Go.

For our example we’ll select across two channels. Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines.
``` Go
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()
```
We’ll use `select` to await both of these values simultaneously, printing each one as it arrives.
``` Go
    for range 2 {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
```
We receive the values "one" and then "two" as expected. Note that the total execution time is only ~2 seconds since both the `1` and `2` second `Sleep`s execute concurrently.
## Timeouts
Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time. Implementing timeouts in Go is easy and elegant thanks to channels and `select`.

For our example, suppose we’re executing an external call that returns its result on a channel `c1` after 2s. Note that the channel is buffered, so the send in the goroutine is nonblocking. This is a common pattern to *prevent goroutine leaks* in case the channel is never read.
``` Go
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()
```
Here’s the `select` implementing a timeout. `res := <-c1` awaits the result and `<-time.After` awaits a value to be sent after the timeout of 1s. Since `select` proceeds with the first receive that’s ready, we’ll take the timeout case if the operation takes more than the allowed 1s.
``` Go
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }
// If we allow a longer timeout of 3s, then the receive from c2 will succeed and we’ll print the result.
    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()

    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
```
Running this program shows the first operation timing out and the second succeeding.
## Non-Blocking Channel Operations
Basic sends and receives on channels are blocking. However, we can use `select` with a `default` clause to implement non-blocking sends, receives, and even non-blocking multi-way `select`s.
``` Go
    messages := make(chan string)
    signals := make(chan bool)
// Here’s a non-blocking receive.
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }
// Here msg cannot be sent to the messages channel, because the channel has no buffer and there is no receiver.
    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }
// Here we attempt non-blocking receives on both messages and signals.
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
```
- Here’s a non-blocking receive. If a value is available on messages then `select` will take the `<-messages` case with that value. If not it will immediately take the `default` case.
- A non-blocking send works similarly. Here `msg` cannot be sent to the `messages` channel, because the channel has no buffer and there is no receiver. Therefore the `default` case is selected.
- We can use multiple cases above the `default` clause to implement a multi-way non-blocking `select`. Here we attempt non-blocking receives on both messages and signals.
### Closing Channels
Closing a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel’s receivers.

In this example we’ll use a `jobs` channel to communicate work to be done from the `main()` goroutine to a `worker` goroutine. When we have no more jobs for the `worker` we’ll close the `jobs` channel.
``` Go
    jobs := make(chan int, 5)
    done := make(chan bool)
// Here’s the worker goroutine. It repeatedly receives from jobs with j, more := <-jobs. 
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()
```
Here’s the worker goroutine. It repeatedly receives from `jobs` with `j, more := <-jobs`. In this special 2-value form of `receive`, the `more` value will be `false` if `jobs` has been closed and all values in the channel have already been received. We use this to notify on `done` when we’ve worked all our `jobs`.
``` Go
// This sends 3 jobs to the worker over the jobs channel, then closes it.
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")
// We await the worker using the synchronization approach we saw earlier.
    <-done

    _, ok := <-jobs
    fmt.Println("received more jobs:", ok)
```
*Reading from a closed channel succeeds immediately*, returning the zero value of the underlying type. The optional second return value is `true` if the value received was delivered by a successful send operation to the channel, or `false` if it was a zero value generated because the channel is closed and empty.

The idea of closed channels leads naturally to our next example.
### Range over Channels
In a previous example we saw how `for` and `range` provide iteration over basic data structures. We can also use this syntax to iterate over values received from a channel.

We’ll iterate over 2 values in the `queue` channel.
``` Go
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)
// This range iterates over each element as it’s received from queue.
    for elem := range queue {
        fmt.Println(elem)
    }
```
This `range` iterates over each element as it’s received from `queue`. Because we closed the channel above, the iteration terminates after receiving the 2 elements.

This example also showed that it’s possible to close a non-empty channel but still have the remaining values be received.
## Timers
We often want to execute Go code at some point in the future, or repeatedly at some interval. Go’s built-in `timer` and `ticker` features make both of these tasks easy. We’ll look first at timers and then at tickers.

Timers represent a single event in the future. You tell the timer how long you want to wait, and it provides a channel that will be notified at that time. This timer will wait 2 seconds.
``` Go
// This timer will wait 2 seconds.
    timer1 := time.NewTimer(2 * time.Second)
// The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer fired.
    <-timer1.C
    fmt.Println("Timer 1 fired")
// You can cancel the timer before it fires. Here’s an example of that.
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
// Give the timer2 enough time to fire, if it ever was going to, to show it is in fact stopped.
    time.Sleep(2 * time.Second)
```
- If you just wanted to wait, you could have used `time.Sleep`. One reason a timer may be useful is that you can cancel the timer before it fires. Here’s an example of that.
- Give the `timer2` enough time to fire, if it ever was going to, to show it is in fact stopped.
- The first timer will fire ~2s after we start the program, but the second should be stopped before it has a chance to fire.
### Tickers
Timers are for when you want to do something once in the future - `ticker`s are for when you want to do something repeatedly at regular intervals. Here’s an example of a `ticker` that ticks periodically until we stop it.

Tickers use a similar mechanism to timers: a channel that is sent values. Here we’ll use the `select` builtin on the channel to await the values as they arrive every 500ms.
``` Go
    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)
// Here we’ll use the select builtin on the channel to await the values as they arrive every 500ms.
    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()
```
Tickers can be stopped like `timer`s. Once a `ticker` is stopped it won’t receive any more values on its channel. We’ll stop ours after 1600ms. When we run this program the `ticker` should tick 3 times before we stop it.
``` Go
    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
```
### Worker Pools
In this example we’ll look at how to implement a worker pool using goroutines and channels.

Here’s the `worker`, of which we’ll run several concurrent instances. These workers will receive work on the `jobs` channel and send the corresponding results on `results`. We’ll `sleep` a second per job to simulate an expensive task.
``` Go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}
```
In order to use our pool of workers we need to send them work and collect their results. We make 2 channels for this.
``` Go
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
// This starts up 3 workers, initially blocked because there are no jobs yet.
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
// Here we send 5 jobs and then close that channel to indicate that’s all the work we have.
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)
// Finally we collect all the results of the work.
    for a := 1; a <= numJobs; a++ {
        <-results
    }
```
- This starts up 3 workers, initially blocked because there are no jobs yet.
- Here we send 5 jobs and then close that channel to indicate that’s all the work we have.
- Finally we collect all the results of the work. This also ensures that the worker goroutines have finished. An alternative way to wait for multiple goroutines is to use a `WaitGroup`.

Our running program shows the 5 jobs being executed by various workers. The program only takes about 2 seconds despite doing about 5 seconds of total work because there are 3 workers operating concurrently.
## WaitGroup
To wait for multiple goroutines to finish, we can use a wait group. This is the function we’ll run in every goroutine.
``` Go
func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)
// Sleep to simulate an expensive task.
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}
```
This `WaitGroup` is used to wait for all the goroutines launched here to finish. Note: if a `WaitGroup` is explicitly passed into functions, *it should be done by pointer*.
``` Go
    var wg sync.WaitGroup
// Launch several goroutines using WaitGroup.Go
    for i := 1; i <= 5; i++ {
        wg.Go(func() {
            worker(i)
        })
    }
// Block until all the goroutines started by wg are done.
    wg.Wait()
```
A goroutine is done when the function it invokes returns.

Note that this approach has no straightforward way to propagate errors from workers. For more advanced use cases, consider using the [`errgroup` package](https://pkg.go.dev/golang.org/x/sync/errgroup).

The order of workers starting up and finishing is likely to be different for each invocation.
## Rate Limiting
Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service. Go elegantly supports rate limiting with goroutines, channels, and `ticker`s.

First we’ll look at basic rate limiting. Suppose we want to limit our handling of incoming requests. We’ll serve these requests off a channel of the same name.
``` Go
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)
// This limiter channel will receive a value every 200 milliseconds.
    limiter := time.Tick(200 * time.Millisecond)
```
This `limiter` channel will receive a value every 200 milliseconds. This is the regulator in our rate limiting scheme.

By blocking on a receive from the `limiter` channel before serving each request, we limit ourselves to 1 request every 200 milliseconds.

We may want to allow short bursts of requests in our rate limiting scheme while preserving the overall rate limit. We can accomplish this by buffering our `limiter` channel. This `burstyLimiter` channel will allow bursts of up to 3 events.
``` Go
// Blocking on a receive--1 request every 200 milliseconds.
    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }
//  This burstyLimiter channel will allow bursts of up to 3 events.
    burstyLimiter := make(chan time.Time, 3)
// Fill up the channel to represent allowed bursting.
    for range 3 {
        burstyLimiter <- time.Now()
    }
// Every 200 milliseconds we’ll try to add a new value to burstyLimiter, up to its limit of 3.
    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()
// Now simulate 5 more incoming requests.
    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
```
Now simulate 5 more incoming requests. The first 3 of these will benefit from the burst capability of `burstyLimiter`. Running our program we see the first batch of requests handled once every ~200 milliseconds as desired.

For the second batch of requests we serve the first 3 immediately because of the `burstable` rate limiting, then serve the remaining 2 with ~200ms delays each.
## Atomic Counters
The primary mechanism for managing state in Go is communication over channels. We saw this for example with worker pools. There are a few other options for managing state though. Here we’ll look at using the `sync/atomic` package for atomic counters accessed by multiple goroutines.
``` Go
// We’ll use an atomic integer type to represent our (always-positive) counter.
    var ops atomic.Uint64
// A WaitGroup will help us wait for all goroutines to finish their work.
    var wg sync.WaitGroup
// A WaitGroup will help us wait for all goroutines to finish their work.
    for range 50 {
        wg.Go(func() {
            for range 1000 {
// To atomically increment the counter we use Add.
                ops.Add(1)
            }
        })
    }
// Wait until all the goroutines are done.
    wg.Wait()
    fmt.Println("ops:", ops.Load())
```
- We’ll use an `atomic` integer type to represent our (always-positive) counter.
- A `WaitGroup` will help us wait for all goroutines to finish their work.
- We’ll start 50 goroutines that each increment the counter exactly 1000 times.
- To atomically increment the counter we use `Add`.
- Here no goroutines are writing to `‘ops’`, but using `Load` it’s *safe to atomically read a value* even while other goroutines are (atomically) updating it.

We expect to get exactly 50,000 operations. Had we used a non-atomic integer and incremented it with `ops++`, we’d likely get a different number, changing between runs, because the goroutines would interfere with each other. Moreover, we’d get data race failures when running with the `-race` flag.
### Mutex
In the previous example we saw how to manage simple counter state using `atomic` operations. For more complex state we can use a mutex to safely access data across multiple goroutines.

`Container` holds a map of counters; since we want to update it concurrently from multiple goroutines, we add a Mutex to synchronize access. Note that mutexes must not be copied, so if this `struct` is passed around, it should be done by pointer.
``` Go
type Container struct {
    mu       sync.Mutex
    counters map[string]int
}

func (c *Container) inc(name string) {
// Lock the mutex before accessing counters; unlock it at the end of the function using a defer statement.
    c.mu.Lock()
    defer c.mu.Unlock()
    c.counters[name]++
}
```
Note that the zero value of a `mutex` is usable as-is, so no initialization is required here.
``` Go
    c := Container{
        counters: map[string]int{"a": 0, "b": 0},
    }

    var wg sync.WaitGroup
// This function increments a named counter in a loop.
    doIncrement := func(name string, n int) {
        for range n {
            c.inc(name)
        }
    }
// Run several goroutines concurrently.
    wg.Go(func() {
        doIncrement("a", 10000)
    })

    wg.Go(func() {
        doIncrement("a", 10000)
    })

    wg.Go(func() {
        doIncrement("b", 10000)
    })
// Wait for the goroutines to finish.
    wg.Wait()
    fmt.Println(c.counters)
```
- Lock the `mutex` before accessing counters; unlock it at the end of the function using a `defer` statement.
- This function increments a named counter in a loop.
- Run several goroutines concurrently; note that they all access the same `Container`, and two of them access the same `counter`.
- Wait for the goroutines to finish

Running the program shows that the counters updated as expected. Next we’ll look at implementing this same state management task using only goroutines and channels.
### Stateful Goroutines
In the previous example we used explicit locking with mutexes to synchronize access to shared state across multiple goroutines. Another option is to use the built-in synchronization features of goroutines and channels to achieve the same result. This channel-based approach aligns with Go’s ideas of *sharing memory by communicating* and having each piece of data owned by exactly one goroutine.

In this example our state will be owned by a single goroutine. This will guarantee that the data is never corrupted with concurrent access. In order to read or write that state, other goroutines will send messages to the owning goroutine and receive corresponding replies. These `readOp` and `writeOp` `struct`s encapsulate those requests and a way for the owning goroutine to respond.
``` Go
type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}
```
As before we’ll count how many operations we perform. The `reads` and `writes` channels will be used by other goroutines to issue read and write requests, respectively.
``` Go
    var readOps uint64
    var writeOps uint64

    reads := make(chan readOp)
    writes := make(chan writeOp)
// Here is the goroutine that owns the state, which is a map as in the previous example but now private to the stateful goroutine.
    go func() {
        var state = make(map[int]int)
        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()
```
Here is the goroutine that owns the `state`, which is a `map` as in the previous example but now private to the stateful goroutine. This goroutine repeatedly selects on the `reads` and `writes` channels, responding to requests as they arrive. A response is executed by first performing the requested operation and then sending a value on the response channel `resp` to indicate success (and the desired value in the case of `reads`).

This starts 100 goroutines to issue reads to the state-owning goroutine via the `reads` channel. Each read requires constructing a `readOp`, sending it over the `reads` channel, and then receiving the result over the provided `resp` channel.
``` Go
    for range 100 {
        go func() {
            for {
                read := readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int)}
                reads <- read
                <-read.resp
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }
// We start 10 writes as well, using a similar approach.
    for range 10 {
        go func() {
            for {
                write := writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }
// Let the goroutines work for a second.
    time.Sleep(time.Second)
// Finally, capture and report the op counts.
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
```
Running our program shows that the goroutine-based state management example completes about 80,000 total operations.

For this particular case the goroutine-based approach was a bit more involved than the `mutex`-based one. It might be useful in certain cases though, for example where you have other channels involved or when managing multiple such mutexes would be error-prone. You should use whichever approach feels most natural, especially with respect to understanding the correctness of your program.
## Sorting
Go’s `slices` package implements sorting for builtins and user-defined types. We’ll look at sorting for builtins first.

Sorting functions are generic, and work for any ordered built-in type. For a list of ordered types, see [cmp.Ordered](https://pkg.go.dev/cmp#Ordered).
``` Go
    strs := []string{"c", "a", "b"}
    slices.Sort(strs)
    fmt.Println("Strings:", strs)
// An example of sorting ints.
    ints := []int{7, 2, 4}
    slices.Sort(ints)
    fmt.Println("Ints:   ", ints)
// We can also use the slices package to check if a slice is already in sorted order.
    s := slices.IsSorted(ints)
    fmt.Println("Sorted: ", s)
```
We can also use the slices package to check if a slice is already in sorted order.
### Sorting by Functions
Sometimes we’ll want to sort a collection by something other than its natural order. For example, suppose we wanted to sort strings by their length instead of alphabetically. Here’s an example of custom sorts in Go.
``` Go
    fruits := []string{"peach", "banana", "kiwi"}
// We implement a comparison function for string lengths. cmp.Compare is helpful for this.
    lenCmp := func(a, b string) int {
        return cmp.Compare(len(a), len(b))
    }
// Now we can call slices.SortFunc with this custom comparison function to sort fruits by name length.
    slices.SortFunc(fruits, lenCmp)
    fmt.Println(fruits)

    type Person struct {
        name string
        age  int
    }
// We can use the same technique to sort a slice of values that aren’t built-in types.
    people := []Person{
        Person{name: "Jax", age: 37},
        Person{name: "TJ", age: 25},
        Person{name: "Alex", age: 72},
    }
// Sort people by age using slices.SortFunc.
    slices.SortFunc(people,
        func(a, b Person) int {
            return cmp.Compare(a.age, b.age)
        })
    fmt.Println(people)
```
- We implement a comparison function for string lengths. `cmp.Compare` is helpful for this.
- Now we can call `slices.SortFunc` with this custom comparison function to sort `fruits` by name length.
- We can use the same technique to sort a slice of values that aren’t built-in types.

Note: if the `Person` `struct` is large, you may want the slice to contain `*Person` instead and adjust the sorting function accordingly. If in doubt, benchmark!

## Panic
A `panic` typically means something went unexpectedly wrong. Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.

We’ll use panic throughout this site to check for unexpected errors. This is the only program on the site designed to panic.

A common use of panic is to abort if a function returns an error value that we don’t know how to (or want to) handle. Here’s an example of panicking if we get an unexpected error when creating a new file.
``` Go
    panic("a problem")

    path := filepath.Join(os.TempDir(), "file")
    _, err := os.Create(path)
    if err != nil {
        panic(err)
    }
// Running this program will cause it to panic, print an error message and goroutine traces, and exit with a non-zero status.
```
When first panic in main fires, the program exits without reaching the rest of the code. If you’d like to see the program try to create a temp file, comment the first panic out.

Note that unlike some languages which use exceptions for handling of many errors, in Go it is idiomatic to use error-indicating return values wherever possible.
### Defer
*Defer* is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. `defer` is often used where e.g. `ensure` and `finally` would be used in other languages.

Suppose we wanted to create a file, write to it, and then close when we’re done. Here’s how we could do that with `defer`.
``` Go
    path := filepath.Join(os.TempDir(), "defer.txt")
    f := createFile(path)
    defer closeFile(f)
    writeFile(f)
```
Immediately after getting a file object with `createFile`, we `defer` the closing of that file with `closeFile`. This will be executed at the end of the enclosing function (`main`), after `writeFile` has finished.
``` Go
func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}
// It’s important to check for errors when closing a file, even in a deferred function.
func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()

    if err != nil {
        panic(err)
    }
}
```
Running the program confirms that the file is closed after being written.
### Recover
Go makes it possible to recover from a panic, by using the recover built-in function. A recover can stop a panic from aborting the program and let it continue with execution instead.

An example of where this can be useful: a server wouldn’t want to crash if one of the client connections exhibits a critical error. Instead, the server would want to close that connection and continue serving other clients. In fact, this is what Go’s net/http does by default for HTTP servers.
``` Go
func mayPanic() {
    panic("a problem")
}
```
`recover` must be called within a deferred function. When the enclosing function panics, the defer will activate and a `recover` call within it will catch the panic.

The return value of `recover` is the `error` raised in the call to `panic`.
``` Go
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered. Error:\n", r)
        }
    }()
    mayPanic()
// This code will not run, because mayPanic panics.
    fmt.Println("After mayPanic()")
```
The execution of main stops at the point of the panic and resumes in the deferred closure.

## String Functions
The standard library’s `strings` package provides many useful string-related functions. Here are some examples to give you a sense of the package.
``` Go
    p("Contains:  ", s.Contains("test", "es"))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
```
We alias `fmt.Println` to a shorter name:

    var p = fmt.Println
as we’ll use it a lot below. Here’s a sample of the functions available in `strings`. Since these are functions from the package, not methods on the `string` object itself, we need to pass the string in question as the first argument to the function. You can find more functions in the [strings package docs](https://pkg.go.dev/strings).
### String Formatting
Go offers excellent support for string formatting in the `printf` tradition. Here are some examples of common string formatting tasks.
``` Go
type point struct {
    x, y int
}
    p := point{1, 2}
```
Go offers several printing “verbs” designed to format general Go values. For example, this prints an instance of our `point` `struct`.
``` Go
    fmt.Printf("struct1: %v\n", p)
    // struct1: {1 2}

//If the value is a struct, the %+v variant will include the struct’s field names.
    fmt.Printf("struct2: %+v\n", p)
    // struct2: {x:1 y:2}

// The %#v variant prints a Go syntax representation of the value. i.e. the source code snippet that would produce that value.
    fmt.Printf("struct3: %#v\n", p)
    // struct3: main.point{x:1, y:2}

// To print the type of a value, use %T.
    fmt.Printf("type: %T\n", p)
    // type: main.point
    fmt.Printf("bool: %t\n", true)
    // bool: true

// There are many options for formatting integers. Use %d for standard, base-10 formatting.
    fmt.Printf("int: %d\n", 123)
    // int: 123

// This prints a binary representation.
    fmt.Printf("bin: %b\n", 14)
    // bin: 1110

// This prints the character corresponding to the given integer.
    fmt.Printf("char: %c\n", 33)
    // char: !

// %x provides hex encoding.
    fmt.Printf("hex: %x\n", 456)
    // hex: 1c8

// There are also several formatting options for floats. For basic decimal formatting use %f.
    fmt.Printf("float1: %f\n", 78.9)
    // float1: 78.900000

// %e and %E format the float in (slightly different versions of) scientific notation.
    fmt.Printf("float2: %e\n", 123400000.0)
    // float2: 1.234000e+08
    fmt.Printf("float3: %E\n", 123400000.0)
    // float3: 1.234000E+08

// For basic string printing use %s.
    fmt.Printf("str1: %s\n", "\"string\"")
    // str1: "string"

// To double-quote strings as in Go source, use %q.
    fmt.Printf("str2: %q\n", "\"string\"")
    // str2: "\"string\""

// As with integers seen earlier, %x renders the string in base-16, with two output characters per byte of input.
    fmt.Printf("str3: %x\n", "hex this")
    // str3: 6865782074686973

// To print a representation of a pointer, use %p.
    fmt.Printf("pointer: %p\n", &p)
    // pointer: 0xc0000ba000

// When formatting numbers you will often want to control the width and precision of the resulting figure. To specify the width of an integer, use a number after the % in the verb. By default the result will be right-justified and padded with spaces.
    fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
    // width1: |    12|   345|

// You can also specify the width of printed floats, though usually you’ll also want to restrict the decimal precision at the same time with the width.precision syntax.
    fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
    // width2: |  1.20|  3.45|

// To left-justify, use the - flag.
    fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
    // width3: |1.20  |3.45  |

// You may also want to control width when formatting strings, especially to ensure that they align in table-like output. For basic right-justified width.
    fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
    // width4: |   foo|     b|

// To left-justify use the - flag as with numbers.
    fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")
    // width5: |foo   |b     |

// So far we’ve seen Printf, which prints the formatted string to os.Stdout. Sprintf formats and returns a string without printing it anywhere.
    s := fmt.Sprintf("sprintf: a %s", "string")
    fmt.Println(s)
    // sprintf: a string

// You can format+print to io.Writers other than os.Stdout using Fprintf.
    fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
    // io: an error
```
### Text Templates
Go offers built-in support for creating dynamic content or showing customized output to the user with the `text/template` package. A sibling package named `html/template` provides the same API but has additional security features and should be used for generating HTML.

We can create a new template and parse its body from a `string`. Templates are a mix of static text and “actions” enclosed in `{{...}}` that are used to dynamically insert content.
``` Go
    t1 := template.New("t1")
    t1, err := t1.Parse("Value is {{.}}\n")
    if err != nil {
        panic(err)
    }
```
Alternatively, we can use the `template.Must` function to panic in case `Parse` returns an error. This is especially useful for templates initialized in the global scope.
``` Go
    t1 = template.Must(t1.Parse("Value: {{.}}\n"))

    t1.Execute(os.Stdout, "some text")
    t1.Execute(os.Stdout, 5)
    t1.Execute(os.Stdout, []string{
        "Go",
        "Rust",
        "C++",
        "C#",
    })
```
By “executing” the template we generate its text with specific values for its actions. The `{{.}}` action is replaced by the value passed as a parameter to `Execute`.

Helper function we’ll use below.
``` Go
    Create := func(name, t string) *template.Template {
        return template.Must(template.New(name).Parse(t))
    }
```
If the data is a `struct` we can use the `{{.FieldName}}` action to access its fields. The fields should be exported to be accessible when a template is executing.
``` Go
    t2 := Create("t2", "Name: {{.Name}}\n")
    t2.Execute(os.Stdout, struct {
        Name string
    }{"Jane Doe"})
```
The same applies to maps; with maps there is no restriction on the case of key names.
``` Go
    t2.Execute(os.Stdout, map[string]string{
        "Name": "Mickey Mouse",
    })
```
`if`/`else` provide conditional execution for templates. A value is considered `false` if it’s the default value of a type, such as `0`, an empty string, `nil` pointer, etc. This sample demonstrates another feature of templates: using `-` in actions to trim whitespace.
``` Go
    t3 := Create("t3",
        "{{if . -}} yes {{else -}} no {{end}}\n")
    t3.Execute(os.Stdout, "not empty")
    t3.Execute(os.Stdout, "")
```
`range` blocks let us loop through slices, arrays, maps or channels. Inside the `range` block `{{.}}` is set to the current item of the iteration.
``` Go
    t4 := Create("t4",
        "Range: {{range .}}{{.}} {{end}}\n")
    t4.Execute(os.Stdout,
        []string{
            "Go",
            "Rust",
            "C++",
            "C#",
        })
```

## Regular Expressions
Go offers built-in support for regular expressions. Here are some examples of common `regexp`-related tasks in Go.

This tests whether a pattern matches a string.
``` Go
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)
```
Above we used a string pattern directly, but for other `regexp` tasks you’ll need to `Compile` an optimized `Regexp` `struct`.
``` Go
    r, _ := regexp.Compile("p([a-z]+)ch")
    fmt.Println(r.MatchString("peach"))
// This finds the match for the regexp.
    fmt.Println(r.FindString("peach punch"))
// This also finds the first match but returns the start and end indexes for the match instead of the matching text.
    fmt.Println("idx:", r.FindStringIndex("peach punch"))
// This will return information for both p([a-z]+)ch and ([a-z]+).
    fmt.Println(r.FindStringSubmatch("peach punch"))
```
The `Submatch` variants include information about both the whole-pattern matches and the submatches within those matches. Similarly this will return information about the indexes of matches and submatches.
``` Go
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))
// Find all matches for a regexp.
    fmt.Println(r.FindAllString("peach punch pinch", -1))
```
The `All` variants of these functions apply to all matches in the input, not just the first. For example to find all matches for a `regexp`. These `All` variants are available for the other functions we saw above as well.
``` Go
    fmt.Println("all:", r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))
// Providing a non-negative integer as the second argument to these functions will limit the number of matches.
    fmt.Println(r.FindAllString("peach punch pinch", 2))
    fmt.Println(r.Match([]byte("peach")))
```
Our examples above had `string` arguments and used names like `MatchString`. We can also provide `[]byte` arguments and drop `String` from the function name.

When creating global variables with regular expressions you can use the `MustCompile` variation of `Compile`. `MustCompile` panics instead of returning an error, which makes it safer to use for global variables.
``` Go
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println("regexp:", r)
// The regexp package can also be used to replace subsets of strings with other values.
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
// The Func variant allows you to transform matched text with a given function.
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
```
For a complete reference on Go regular expressions check the [regexp package docs](https://pkg.go.dev/regexp).

## JSON
Go offers built-in support for JSON encoding and decoding, including to and from built-in and custom data types. We’ll use these two structs to demonstrate encoding and decoding of custom types below.
``` Go
type response1 struct {
    Page   int
    Fruits []string
}

type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}
```
Only exported fields will be encoded/decoded in JSON. *Fields must start with capital letters to be exported*. First we’ll look at encoding basic data types to JSON strings. Here are some examples for atomic values.
``` Go
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))
    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))
    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))
    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))
// And here are some for slices and maps, which encode to JSON arrays and objects as you’d expect.
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))
    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))
```
The JSON package can automatically encode your custom data types. It will only include exported fields in the encoded output and will by default use those names as the JSON keys.
``` Go
    res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))
```
You can use tags on struct field declarations to customize the encoded JSON key names. Check the definition of response2 above to see an example of such tags.
``` Go
    res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))
```
Now let’s look at decoding JSON data into Go values. Here’s an example for a generic data structure.
``` Go
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
    var dat map[string]any
```
We need to provide a variable where the JSON package can put the decoded data. This `map[string]any` will hold a `map` of strings to arbitrary data types.
``` Go
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)
``` 
Here’s the actual decoding, and a check for associated errors. For the sake of brevity we ignore the errors in these examples; in real code, you should always check for errors and act upon them.

In order to use the values in the decoded map, we’ll need to convert them to their appropriate type. For example here we convert the value in `num` to the expected `float64` type.
``` Go
    num := dat["num"].(float64)
    fmt.Println(num)
// Accessing nested data requires a series of conversions.
    strs := dat["strs"].([]any)
    str1 := strs[0].(string)
    fmt.Println(str1)
```
We can also decode JSON into custom data types. This has the advantages of adding additional type-safety to our programs and eliminating the need for type assertions when accessing the decoded data.
``` Go
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{}
    _ = json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    _ = enc.Encode(d)
```
In the examples above we always used bytes and strings as intermediates between the data and JSON representation on standard out. We can also stream JSON encodings directly to `os.Writers` like `os.Stdout` or even HTTP response bodies.

Streaming reads from `os.Readers` like `os.Stdin` or HTTP request bodies is done with `json.Decoder`.
``` Go
    dec := json.NewDecoder(strings.NewReader(str))
    res1 := response2{}
    _ = dec.Decode(&res1)
    fmt.Println(res1)
```
Check out the [JSON and Go blog post](https://go.dev/blog/json) and [JSON package docs](https://pkg.go.dev/encoding/json) for more.

## XML
Go offers built-in support for XML and XML-like formats with the `encoding/xml` package.
``` Go
type Plant struct {
    XMLName xml.Name `xml:"plant"`
    Id      int      `xml:"id,attr"`
    Name    string   `xml:"name"`
    Origin  []string `xml:"origin"`
}
```
`Plant` will be mapped to XML. Similarly to the JSON examples, field tags contain directives for the encoder and decoder. Here we use some special features of the XML package: the `XMLName` field name dictates the name of the XML element representing this `struct`; `id,attr` means that the `Id` field is an XML attribute rather than a nested element.
``` Go
func (p Plant) String() string {
    return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
        p.Id, p.Name, p.Origin)
}
```
Emit XML representing our plant; using `MarshalIndent` to produce a more human-readable output.
``` Go
    coffee := &Plant{Id: 27, Name: "Coffee"}
    coffee.Origin = []string{"Ethiopia", "Brazil"}

    out, _ := xml.MarshalIndent(coffee, " ", "  ")
    fmt.Println(string(out))
// To add a generic XML header to the output, append it explicitly.
    fmt.Println(xml.Header + string(out))
```
Use `Unmarshal` to parse a stream of bytes with XML into a data structure. If the XML is malformed or cannot be mapped onto `Plant`, a descriptive `error` will be returned.
``` Go
    var p Plant
    if err := xml.Unmarshal(out, &p); err != nil {
        panic(err)
    }
    fmt.Println(p)
    tomato := &Plant{Id: 81, Name: "Tomato"}
    tomato.Origin = []string{"Mexico", "California"}
```
The `parent>child>plant` field tag tells the encoder to nest all plants under `<parent><child>...`
``` Go
    type Nesting struct {
        XMLName xml.Name `xml:"nesting"`
        Plants  []*Plant `xml:"parent>child>plant"`
    }
    nesting := &Nesting{}
    nesting.Plants = []*Plant{coffee, tomato}
    out, _ = xml.MarshalIndent(nesting, " ", "  ")
    fmt.Println(string(out))
```

## Time
Go offers extensive support for times and durations; here are some examples. We’ll start by getting the current time.
``` Go
    p := fmt.Println

    now := time.Now()
    p(now)
```
You can build a `time` `struct` by providing the year, month, day, etc. Times are always associated with a `Location`, i.e. time zone. You can extract the various components of the `time` value as expected.
``` Go
    then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())
// The Monday-Sunday Weekday is also available.
    p(then.Weekday())
```
These methods compare two `time`s, testing if the first occurs before, after, or at the same time as the second, respectively.
``` Go
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))
```
The `Sub` methods returns a `Duration` representing the interval between two `time`s.
``` Go
    diff := now.Sub(then)
    p(diff)
// We can compute the length of the duration in various units.
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())
```
You can use `Add` to advance a `time` by a given duration, or with a `-` to move backwards by a duration.
``` Go
    p(then.Add(diff))
    p(then.Add(-diff))
```
Next we’ll look at the related idea of time relative to the Unix epoch.
### Epoch
A common requirement in programs is getting the number of seconds, milliseconds, or nanoseconds since the Unix epoch. Here’s how to do it in Go. Use `time.Now` with `Unix`, `UnixMilli` or `UnixNano` to get elapsed time since the Unix epoch in seconds, milliseconds or nanoseconds, respectively.
``` Go
    now := time.Now()
    fmt.Println(now)
    fmt.Println(now.Unix())
    fmt.Println(now.UnixMilli())
    fmt.Println(now.UnixNano())
// You can also convert integer seconds or nanoseconds since the epoch into the corresponding time.
    fmt.Println(time.Unix(now.Unix(), 0))
    fmt.Println(time.Unix(0, now.UnixNano()))
```
Next we’ll look at another time-related task: time parsing and formatting.
### Time Formatting / Parsing
Go supports time formatting and parsing via pattern-based layouts. Here’s a basic example of formatting a time according to RFC3339, using the corresponding layout constant. Time parsing uses the same layout values as `Format`.
``` Go
    p := fmt.Println

    t := time.Now()
    p(t.Format(time.RFC3339))
// Time parsing uses the same layout values as Format.
    t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
    p(t1)
```
`Format` and `Parse` use example-based layouts. Usually you’ll use a constant from `time` for these layouts, but you can also supply custom layouts. Layouts must use the reference time `Mon Jan 2 15:04:05 MST 2006` to show the pattern with which to format/parse a given time/string. The example time must be exactly as shown: the year 2006, 15 for the hour, Monday for the day of the week, etc.
``` Go
    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, _ := time.Parse(form, "8 41 PM")
    p(t2)
```
For purely numeric representations you can also use standard string formatting with the extracted components of the `time` value. `Parse` will return an `error` on malformed input explaining the parsing problem.
``` Go
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())
// Parse will return an error on malformed input explaining the parsing problem.
    _, err := time.Parse("Mon Jan _2 15:04:05 2006", "8:41PM")
    p(err)
```

## Random Numbers
Go’s [`math/rand/v2`](https://pkg.go.dev/math/rand/v2) package provides pseudorandom number generation. For example, `rand.IntN` returns a random `int n, 0 <= n < 100`.
``` Go
    fmt.Print(rand.IntN(100), ",")
    fmt.Print(rand.IntN(100))
    fmt.Println()
    fmt.Println(rand.Float64())
```
`rand.Float64` returns a `float64 f, 0.0 <= f < 1.0`. This can be used to generate random floats in other ranges, for example `5.0 <= f' < 10.0`.
``` Go
    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)
    fmt.Println()
```
If you want a known seed, create a new `rand.Source` and pass it into the `New` constructor. `NewPCG` creates a new [PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator) source that requires a `seed` of two `uint64` numbers.
``` Go
    s2 := rand.NewPCG(42, 1024)
    r2 := rand.New(s2)
    fmt.Print(r2.IntN(100), ",")
    fmt.Print(r2.IntN(100))
    fmt.Println()

    s3 := rand.NewPCG(42, 1024)
    r3 := rand.New(s3)
    fmt.Print(r3.IntN(100), ",")
    fmt.Print(r3.IntN(100))
    fmt.Println()
```
Some of the generated numbers may be different when you run the sample. 

## Number Parsing
Parsing numbers from strings is a basic but common task in many programs; here’s how to do it in Go. The built-in package `strconv` provides the number parsing.

With `ParseFloat`, this `64` tells how many bits of precision to parse.
``` Go
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)
```
For `ParseInt`, the `0` means infer the base from the string. `64` requires that the result fit in `64` bits.
``` Go
    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)
// ParseInt will recognize hex-formatted numbers.
    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)
// A ParseUint is also available.
    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)
// Atoi is a convenience function for basic base-10 int parsing.
    k, _ := strconv.Atoi("135")
    fmt.Println(k)
    _, e := strconv.Atoi("wat")
    fmt.Println(e)
```
Atoi is a convenience function for basic base-10 int parsing. Parse functions return an error on bad input.

## URL Parsing
URLs provide a [uniform way to locate resources](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/). Here’s how to parse URLs in Go. We’ll parse this example URL, which includes a scheme, authentication info, host, port, path, query params, and query fragment.
``` Go
    s := "postgres://user:pass@host.com:5432/path?k=v#f"
// Parse the URL and ensure there are no errors.
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }
    fmt.Println(u.Scheme)
    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)
```
`User` contains all authentication info; call `Username` and `Password` on this for individual values. The `Host` contains both the hostname and the port, if present. Use `SplitHostPort` to extract them.
``` Go
    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)
// Here we extract the path and the fragment after the #.
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
```
To get query params in a string of `k=v` format, use `RawQuery`. You can also parse query params into a map. The parsed query param maps are from strings to slices of strings, so index into `[0]` if you only want the first value. Running our URL parsing program shows all the different pieces that we extracted.
``` bash
$ go run url-parsing.go 
postgres
user:pass
user
pass
host.com:5432
host.com
5432
/path
f
k=v
map[k:[v]]
v
```

## SHA256 Hashes
SHA256 hashes are frequently used to compute short identities for binary or text blobs. For example, TLS/SSL certificates use SHA256 to compute a certificate’s signature. Here’s how to compute SHA256 hashes in Go. Go implements several hash functions in various `crypto/*` packages. Here we start with a new hash.
``` Go
    s := "sha256 this string"
    h := sha256.New()
    h.Write([]byte(s))
```
`Write` expects bytes. If you have a string `s`, use `[]byte(s)` to *coerce* it to bytes. This gets the finalized hash result as a byte slice. The argument to `Sum` can be used to append to an existing byte slice: it usually isn’t needed.
``` Go
    bs := h.Sum(nil)
    fmt.Println(s)
    fmt.Printf("%x\n", bs)
```
Running the program computes the hash and prints it in a human-readable hex format. You can compute other hashes using a similar pattern to the one shown above. For example, to compute SHA512 hashes import `crypto/sha512` and use `sha512.New()`.

Note that if you need cryptographically secure hashes, you should carefully research hash strength!

## Base64 Encoding
Go provides built-in support for base64 encoding/decoding. This syntax imports the `encoding/base64` package with the `b64` name instead of the default `base64`. It’ll save us some space below. Here’s the string we’ll encode/decode.
``` Go
    data := "abc123!?$*&()'-=@~"
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)
```
Go supports both standard and URL-compatible `base64`. Here’s how to encode using the standard encoder. The encoder requires a `[]byte` so we convert our `string` to that type.

Decoding may return an `error`, which you can check if you don’t already know the input to be well-formed.
``` Go
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()
// This encodes/decodes using a URL-compatible base64 format.
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
```
The string encodes to slightly different values with the standard and URL `base64` encoders (trailing `+` vs `-`) but they both decode to the original `string` as desired.

## Reading Files
Reading and writing files are basic tasks needed for many Go programs. First we’ll look at some examples of reading files. Reading files requires checking most calls for errors. This helper will streamline our error checks below.
``` Go
func check(e error) {
    if e != nil {
        panic(e)
    }
}
```
Perhaps the most basic file reading task is slurping a file’s entire contents into memory.
``` Go
    path := filepath.Join(os.TempDir(), "dat")
    dat, err := os.ReadFile(path)
    check(err)
    fmt.Print(string(dat))
```
You’ll often want more control over how and what parts of a file are read. For these tasks, start by `Open`ing a file to obtain an `os.File` value.
``` Go
    f, err := os.Open(path)
    check(err)
// Read some bytes from the beginning of the file.
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
```
Read some bytes from the beginning of the file. Allow up to `5` to be read but also note how many actually were read. You can also `Seek` to a known location in the file and `Read` from there.
``` Go
    o2, err := f.Seek(6, io.SeekStart)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: ", n2, o2)
    fmt.Printf("%v\n", string(b2[:n2]))
```
Other methods of seeking are relative to the current cursor position,
``` Go
    _, err = f.Seek(2, io.SeekCurrent)
    check(err)
// and relative to the end of the file.
    _, err = f.Seek(-4, io.SeekEnd)
    check(err)
```
The `io` package provides some functions that may be helpful for file reading. For example, reads like the ones above can be more robustly implemented with `ReadAtLeast`.
``` Go
    o3, err := f.Seek(6, io.SeekStart)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
```
There is no built-in rewind, but Seek(0, io.SeekStart) accomplishes this.
``` Go
    _, err = f.Seek(0, io.SeekStart)
    check(err)
```
The `bufio` package implements a buffered reader that may be useful both for its efficiency with many small reads and because of the additional reading methods it provides.
``` Go
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    f.Close()
```
`Close` the file when you’re done (usually this would be scheduled immediately after `Open`ing with `defer`).
### Writing Files
``` Go
func check(e error) {
    if e != nil {
        panic(e)
    }
}
```
Writing files in Go follows similar patterns to the ones we saw earlier for reading.
``` Go
// To start, here’s how to dump a string (or just bytes) into a file.
    d1 := []byte("hello\ngo\n")
    path1 := filepath.Join(os.TempDir(), "dat1")
    err := os.WriteFile(path1, d1, 0644)
    check(err)
// For more granular writes, open a file for writing.
    path2 := filepath.Join(os.TempDir(), "dat2")
    f, err := os.Create(path2)
    check(err)
// It’s idiomatic to defer a Close immediately after opening a file.
    defer f.Close()
// You can Write byte slices as you’d expect.
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)
// A WriteString is also available.
    n3, err := f.WriteString("writes\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n3)
// Issue a Sync to flush writes to stable storage.
    f.Sync()
// bufio provides buffered writers in addition to the buffered readers we saw earlier.
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n4)
// Use Flush to ensure all buffered operations have been applied to the underlying writer.
    w.Flush()
```
Next we’ll look at applying some of the file I/O ideas we’ve just seen to the stdin and stdout streams.

## Line Filters
A line filter is a common type of program that reads input on `stdin`, processes it, and then prints some derived result to `stdout`. `grep` and `sed` are common line filters.

Here’s an example line filter in Go that writes a capitalized version of all input text. You can use this pattern to write your own Go line filters. Wrapping the unbuffered `os.Stdin` with a buffered scanner gives us a convenient `Scan` method that advances the scanner to the next token; which is the next line in the default scanner.
``` Go
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        ucl := strings.ToUpper(scanner.Text())
// Write out the uppercased line.
        fmt.Println(ucl)
    }
```
`Text` returns the current token, here the next line, from the input. Check for errors during `Scan`. End of file is expected and not reported by `Scan` as an `error`.
``` Go
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
```

## File Paths
The `filepath` package provides functions to parse and construct file paths in a way that is portable between operating systems; `dir/file` on Linux vs. `dir\file` on Windows, for example.

`Join` should be used to construct paths in a portable way. It takes any number of arguments and constructs a hierarchical `path` from them.
``` Go
    p := filepath.Join("dir1", "dir2", "filename")
    fmt.Println("p:", p)
    fmt.Println(filepath.Join("dir1//", "filename"))
    fmt.Println(filepath.Join("dir1/../dir1", "filename"))
```
You should always use `Join` instead of concatenating `/`s or `\`s manually. In addition to providing portability, `Join` will also normalize paths by removing superfluous separators and directory changes.

`Dir` and `Base` can be used to split a path to the directory and the file. Alternatively, `Split` will return both in the same call.
``` Go
    fmt.Println("Dir(p):", filepath.Dir(p))
    fmt.Println("Base(p):", filepath.Base(p))
// We can check whether a path is absolute.
    fmt.Println(filepath.IsAbs("dir/file"))
    fmt.Println(filepath.IsAbs("/dir/file"))
```
Some file names have extensions following a dot. We can split the extension out of such names with `Ext`.
``` Go
    ext := filepath.Ext(filename)
    fmt.Println(ext)
// To find the file’s name with the extension removed, use strings.TrimSuffix.
    fmt.Println(strings.TrimSuffix(filename, ext))
```
To find the file’s name with the extension removed, use `strings.TrimSuffix`.

`Rel` finds a relative path between a base and a target. It returns an `error` if the target cannot be made relative to base.
``` Go
    rel, err := filepath.Rel("a/b", "a/b/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
    rel, err = filepath.Rel("a/b", "a/c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
```
### Directories
``` Go
func check(e error) {
    if e != nil {
        panic(e)
    }
}
```
Go has several useful functions for working with directories in the file system. Create a new sub-directory in the current working directory. When creating temporary directories, it’s good practice to `defer` their removal. `os.RemoveAll` will delete a whole directory tree (similarly to `rm -rf`).
``` Go
    err := os.Mkdir("subdir", 0755)
    check(err)
    defer os.RemoveAll("subdir")
// Helper function to create a new empty file.
    createEmptyFile := func(name string) {
        d := []byte("")
        check(os.WriteFile(name, d, 0644))
    }
    createEmptyFile("subdir/file1")
```
We can create a hierarchy of directories, including parents with `MkdirAll`. This is similar to the command-line `mkdir -p`.
``` Go
    err = os.MkdirAll("subdir/parent/child", 0755)
    check(err)
    createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")
```
`ReadDir` lists directory contents, returning a slice of `os.DirEntry` objects.
``` Go
    c, err := os.ReadDir("subdir/parent")
    check(err)
    fmt.Println("Listing subdir/parent")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }
```
`Chdir` lets us change the current working directory, similarly to `cd`. Now we’ll see the contents of `subdir/parent/child` when listing the current directory.
``` Go
    err = os.Chdir("subdir/parent/child")
    check(err)
    c, err = os.ReadDir(".")
    check(err)
    fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }
// cd back to where we started.
    err = os.Chdir("../../..")
    check(err)
```
We can also `visit` a directory recursively, including all its sub-directories. `WalkDir` accepts a callback function to handle every file or directory visited.
``` Go
    err = os.Chdir("../../..")
    check(err)
    fmt.Println("Visiting subdir")
    err = filepath.WalkDir("subdir", visit)
    check(err)
```
`visit` is called for every file or directory found recursively by `filepath.WalkDir`.
``` Go
func visit(path string, d fs.DirEntry, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", path, d.IsDir())
    return nil
}
```
### Temporary Files and Directories
Throughout program execution, we often want to create data that isn’t needed after the program exits. Temporary files and directories are useful for this purpose since they don’t pollute the file system over time.
``` Go
func check(e error) {
    if e != nil {
        panic(e)
    }
}
```
The easiest way to create a temporary file is by calling `os.CreateTemp`. It creates a file and opens it for reading and writing. We provide `""` as the first argument, so `os.CreateTemp` will create the file in the default location for our OS.
``` Go
    f, err := os.CreateTemp("", "sample")
    check(err)
// Display the name of the temporary file
    fmt.Println("Temp file name:", f.Name())
// Clean up the file after we’re done
    defer os.Remove(f.Name())
// We can write some data to the file
    _, err = f.Write([]byte{1, 2, 3, 4})
    check(err)
```
- Display the name of the temporary file. On Unix-based OSes the directory will likely be `/tmp`. The file name starts with the prefix given as the second argument to `os.CreateTemp` and the rest is chosen automatically to ensure that concurrent calls will always create different file names.
- Clean up the file after we’re done. The OS is likely to clean up temporary files by itself after some time, but it’s good practice to do this explicitly.
- We can write some data to the file.

If we intend to write many temporary files, we may prefer to create a temporary directory. `os.MkdirTemp`’s arguments are the same as `CreateTemp`’s, but it returns a directory name rather than an open file.
``` Go
    dname, err := os.MkdirTemp("", "sampledir")
    check(err)
    fmt.Println("Temp dir name:", dname)
    defer os.RemoveAll(dname)
```
Now we can synthesize temporary file names by prefixing them with our temporary directory.
``` Go
    fname := filepath.Join(dname, "file1")
    err = os.WriteFile(fname, []byte{1, 2}, 0666)
    check(err)
```
### Embed Directive
`//go:embed` is a [compiler directive](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives) that allows programs to include arbitrary files and folders in the Go binary at build time. Read more about the embed directive [here](https://pkg.go.dev/embed).

Import the `embed` package; if you don’t use any exported identifiers from this package, you can do a blank import with `_ "embed".`

`embed` directives accept paths relative to the directory containing the Go source file. This directive embeds the contents of the file into the `string` variable immediately following it.
``` Go
import "embed"
//go:embed folder/single_file.txt
var fileString string
//     print(fileString)

// Or embed the contents of the file into a []byte.
//go:embed folder/single_file.txt
var fileByte []byte
//     print(string(fileByte))
```
We can also `embed` multiple files or even folders with wildcards. This uses a variable of the [`embed.FS`](https://pkg.go.dev/embed#FS) type, which implements a simple virtual file system.
``` Go
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS
```
Retrieve some files from the embedded folder.
``` Go
    content1, _ := folder.ReadFile("folder/file1.hash")
    print(string(content1))
    content2, _ := folder.ReadFile("folder/file2.hash")
    print(string(content2))
```
Use these commands to run the example.
``` bash
$ mkdir -p folder
$ echo "hello go" > folder/single_file.txt
$ echo "123" > folder/file1.hash
$ echo "456" > folder/file2.hash
$ go run embed-directive.go
```

## Testing and Benchmarking
