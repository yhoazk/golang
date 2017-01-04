# golang
Tutorials and demo codes for go




Go exported functions are differentiated as the first letter in the function
name is uppercase.

Go brackets for a function need to be in the same line of the function
deglaration, for example:
This code will give a compilation error:
```
package main
import "fmt"

func main()
{
    fmt.Println("I wont compile ")
}
```
>> syntax error: unexpected semicolon or newline before {
But this code will run w/o problems:
```
package main
import "fmt"

func main(){
    fmt.Println("I wont compile ")
}
```

Go variables are explicitily declared before use, much like C.
the keyword `var` declares one or multiple variables at once.

If a variable is not initialized, its zeroed.

The symbol `:=` is a short hand for declare and initialize.
e.g.:
This two codes are equivalent:
var f string = "short"


