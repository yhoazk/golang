# golang
Tutorials and demo codes for go




Go exported functions are differentiated by **the first letter in the function
name is uppercase**.

Go brackets for a function need to be in **the same line of the function
deglaration**, for example:
This code will give a compilation error:
```go
package main
import "fmt"

func main()
{
    fmt.Println("I wont compile ")
}
```
> syntax error: unexpected semicolon or newline before {
But this code will run w/o problems:
```go
package main
import "fmt"

func main(){
    fmt.Println("I will compile ")
}
```

Go variables are explicitily declared before use, much like C.
the keyword `var` declares one or multiple variables at once.

If a variable is not initialized, it's zeroed.

The symbol `:=` is a short hand for declare and initialize.
e.g.:
This two codes are equivalent:
```go
var f string = "short"
```
and
```go
f := "short"
```

`go` has the next built in types:
* `bool`:
* `string`:
* `int int8 int16 int16 int32 int64`:
* `uint uint8 uint16 uint16 uint32 uint64 uintptr`:
* `byte` == `uint8`:
* `float32 float64`:
* `complex64 complex128`

```go
	z      complex128 = cmplx.Sqrt(-5 + 12i)
  g := 0.867 + 0.5i // complex128
```
`int`, `uint`, `uintptr`:
These types are 32-bit in 32 bit sytems and 64 in 64 bit systems. When an int is needed use an int
unless you have a specific reason to use an Uint
