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