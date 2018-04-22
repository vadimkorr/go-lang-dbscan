# Go cheatsheet

## **Basics**

### Packages
* program made of pkgs
* start running in pkg `main`
* pkg name repeats the last element of path
* 
```go
package main
func main() {}
```

### Imports
* factored, paranthesized
```go
import ("fmt"; "math")
```
* 
```go
import "fmt"
import "math"
```
*
```go
import (
    "fmt"
    "math"
)
```
### Exported names
* Start with a capital letter

### Functions
* 
```go
func add(x int, y int) int {
    return x + y
}
// or
func add(x, y int) int {
    return x + y
}
// add(5, 10)
```

### Multiple results
*
```go
func swap(x, y string) (string, string) {
    return y, x
}
// a, b := swap("1", "2")
```

### Named return values
* return values named
* names should be used
* "naked" return return named return values
* harms readability

### Variables
* `var` declares a list of variables
* `var c, python, java bool`

### Variables with initializers
* `var i, j  int = 1, 2`
* initializer is present - type can be omitted

### Short variable declaration
* `:=` used in place of `var` with implicit type
* only inside a function

### Basic types
* `bool`
* `string`
* `int` `int8`(`byte`) `int16` `int32`(`rune`) `int64`
* `uint` `uint8` `uint16` `uint32` `uint64` `uintptr`
* `float32` `float64`
* `complex64` `complex64`

### Zero values
* variables declared w/out an exlicit init value are given their zero value
* `0` - numeric
* `false` - boolean
* `""` - strings

### Type conversions
* `T(v)` - converts value `v` to the type `T`
* for different types explicit conversion is required

### Type inference
* the type is inferred from value on the right hand side

### Constants
* `const`
* no `:=` syntax
* `const Pi = 3.14` 

## **Flow control**

### For
* 
```go
for i:=0; i<10; i++ {

}
```
* `{ }` are required
* init and post are optional
```go
sum := 1
for ; sum < 1000; {
    sum += sum
}
```
* `while` loop
```go
sum := 1
for sum < 1000 {
    sum += sum
}
```
* infinite loop
```go
for {
    
}
```
* If
```go
if x < 0 {

}
// or 
if (x < 0) {

}
```
### If with a short statement
*
```go
if v := math.Pow(x, n); v < lim {
    return v
}
```
### If and else
*
```go
if x < 10 {

} else {

}
```
* variables declared inside short statement are also available inside any of `else` blocks

### Switch
* runs only selected case
* case need not to be constant
*
```go
switch os := runtime.GOOS: os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default: 
        fmt.Println("%s.", os)
}
```

### Switch with no condition
* clean way to write long if-then-else chains
*
```go
t := time.Now()
switch {
    case t.Hour()<12:
        fmt.Println("Good morning")
    case t.Hour()<17:
        fmt.Println("Good afternoon")
    default:
        fmt.Println("Good evening")
}
```
### Defer
* defers the exec of a func until the surrounding func returns
* `defer fmt.Println("world")`

### Stacking defers
* Deferred function calls are pushed onto a stack

## **More types**

### Pointers

* hold the memory address of value
* `*T` is a pointer to a `T` value
```go
var p *int
```
* `&` operator generates a pointer
```go
i := 42
p = &i
```
* `*` operator denotes the pointer's underlying value (dereferencing/indirecting)
```go
fmt.Println(*p)
*p = 21
```

### Structs
* `struct` is a collection of fields
* 
```go
type Vertex struct {
    X int
    Y int
}
fmt.Println(Vertex{1, 2})
```
* Struct fields are accessing using a dot
```go
v := Vertex{1, 2}
v.X = 4
```
### Pointers to a struct
*
```go
(*p).X // or just p.X
```
*
```go
v := Vertex{1, 2}
p := &v
p.X = 1e9
```

### Struct Literals
* a newly allocated struct value by listing the values of its fields
* 
```go
v1 = Vertex{1, 2} // has type Vertex
v2 = Vertex{X: 1} // Y:0 is implicit
v3 = Vertex{} // X:0 and Y:0
p = &Vertex{1, 2} // has type *Vertex
```
### Arrays
* The type `[n]T` is an array of `n` values of type `T`
* `var a [10]int` declares a variable `a` as an array of ten integers

### Slices
* an array has a fixed size 
* slice - is a dynamically-sized, flexible view into the elements of an array
* The type `[]T` is a slice with elements of type `T`
* Slice is formed by specifying two indices, a low and high bound, separated by a colon: `a[low : high]` - includes the first, excludes the last
*
```go
primes := [6]int{2, 3, 5, 7, 11, 13}
var s []int = primes[1:4]
```

### Slices are like references to arrays

* A slice does not store any data, it just describes a section of an underlying array
* Changing the elements of a slice modifies the corresponding elements of its underlying array

### Slice literals
* A slice literal is like an array w/out the length
* array literal `[3]bool{true, true, false}`
* this creates the same array as above, then builds a slice that references it:
`[]bool{true, true, false}`

### Slice defaults
* when slicing, you may omit the high or low bounds to use their defaults instead
* 
```go
var a [10]int
// these slice expressions are eq
a[0:10]
a[:10]
a[0:]
a[:]
```
