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
