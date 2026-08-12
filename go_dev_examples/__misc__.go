// https://go.dev/tour/list
package misc

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

// The var statement declares a list of variables; as in function argument lists, the type is last.
var c, python, java bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Numeric constants are high-precision values. An untyped constant takes the type needed by its context.
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

type Vertex struct {
	X int
	Y int
}

type MyFloat float64

type IAbs interface {
	abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

// Error interface
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func godevtutorial() {
	fmt.Println("Hello, 世界")

	fmt.Println("The time is", time.Now())
	// A name is exported if it begins with a capital letter.
	// For example, Pizza is an exported name, as is Pi, which is exported from the math package.
	// Any "unexported" names are not accessible from outside the package.
	fmt.Println("My favorite number is", rand.Intn(10), " or ", math.Pi)

	fmt.Println(add(42, 13))
	// If an initializer is present, the type can be omitted:
	var i, j int = 1, 2
	fmt.Println(i, j, c, python, java)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// The expression T(v) converts the value v to the type T.
	var f float64 = math.Sqrt(float64(i*i + j*j))
	var k uint = uint(f)
	fmt.Println(i, j, k)

	// Constants cannot be declared using the := syntax.
	const Truth = true
	fmt.Println("Go rules?", Truth)

	// An int can store at maximum a 64-bit integer, and sometimes less.
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// The init and post statements are optional.
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// At that point you can drop the semicolons: C's while is spelled for in Go.
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// If you omit the loop condition it loops forever, so an infinite loop is compactly expressed: for {}

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// Go's switch is like the one in C except that Go only runs the selected case,
	// not all the cases that follow. In effect, the break statement that is needed at the end
	// of each case in those languages is provided automatically in Go.
	// Another important difference is that Go's switch cases need not be
	// constants, and the values involved need not be integers.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s\n", os)
	}

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today is Saturday.")
	case today + 1:
		fmt.Println("Tomorrow is Saturday.")
	case today + 2:
		fmt.Println("Saturday is in two days.")
	default:
		fmt.Println("Saturday is too far away.")
	}
	// Switch without a condition is the same as switch true.
	// This construct can be a clean way to write long if-then-else chains.

	// A defer statement defers the execution of a function until the surrounding function returns.
	defer fmt.Println("world!")
	// The deferred call's arguments are evaluated immediately. Deferred function calls are pushed onto a stack.

	fmt.Println(Vertex{1, 2})

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// Changing the elements of a slice modifies the corresponding elements of its underlying array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	a := names[0:2]
	b := names[1:3]
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	sm := make([]int, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(sm), cap(sm), sm)

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var sn []int // nil-slice
	sn = append(sn, 0)
	sn = append(sn, 1)
	sn = append(sn, 2)
	sn = append(sn, 3, 4, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(sn), cap(sn), sn)

	for i, v := range sn {
		fmt.Printf("sn[%d] = %d\n", i, v)
	}

	for _, value := range sn {
		fmt.Printf("%d\n", value)
	}

	var m = map[string]Vertex{
		"Bell Labs": {40, -74},
		"Google":    {37, -122},
	}
	fmt.Println(m)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(math.Pow))

	fa := MyFloat(-math.Sqrt2)
	fmt.Println(fa.abs())

	// interface implementing abs
	var ia IAbs
	fia := MyFloat(-math.Sqrt2)
	via := Vertex{3, 4}

	ia = fia
	fmt.Println(ia.abs())
	ia = &via
	fmt.Println(ia.abs())

	var ii I = T{"Hello,"}
	defer ii.M()
	fmt.Printf("(%v, %T)\n", ii, ii)

	// Empty interfaces
	var ei interface{}
	fmt.Printf("(%v, %T)\n", ei, ei)

	ei = 42
	fmt.Printf("(%v, %T)\n", ei, ei)

	// Type assertions
	var ais interface{} = "hello"

	bis := ais.(string)
	fmt.Println(bis)

	_, bisp := ais.(string)
	fmt.Println(bisp)

	cis, isok := ais.(float64)
	fmt.Println(cis, isok)

	do(42)
	do(false)

}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// Stringer is a type that can describe itself as a string
func (p Vertex) String() string {
	return fmt.Sprintf("%v (%v years)", p.X, p.Y)
}

// https://go.dev/blog/declaration-syntax
func add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
// A function can return any number of results.
func swap(x, y string) (string, string) {
	return y, x
}

// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}

// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
func test() {
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(k, c, python, java)
}

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func strsqrt(x float64) string {
	if x < 0 {
		return strsqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// Like for, the if statement can start with a short statement to execute before the condition.
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// Exercise: https://go.dev/tour/flowcontrol/8
func sqrt(x float64) (z float64) {
	z = 1.
	for y := x; ; {
		z -= (z*z - x) / (2 * z)
		if z == y || y-z <= 0.0000001 {
			return
		}
		y = z
	}
}

// Exercise: https://go.dev/tour/moretypes/23
func WordCount(s string) (m map[string]int) {
	a := strings.Fields(s)
	m = make(map[string]int)
	for _, v := range a {
		m[v] += 1
	}
	return
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// adder function returns a closure.
// Each closure is bound to its own sum variable.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// Exercise: https://go.dev/tour/moretypes/26
func fibonacci() func() int {
	var a = 0
	var b = 1
	return func() (c int) {
		c = a + b
		a = b
		b = c
		return
	}
}

// Method example for Vertex struct
func (v Vertex) abs() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

// Methods can be overloaded
func (f MyFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
