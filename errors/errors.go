package main

/* In Go it's idiomatic to communicate errors via an explicit
	separate return value.
*/



import "fmt"
import "errors"
/*   name arguments return args */
/* By convention errors are the last return value and have rror type, */
/* error is a built in interface in Go  */
func f1(arg int) (int, error){
	if arg == 42 {
		/* errors.New creates a basic error value with given error message */
		return -1, errors.New("Cant work with 42")
	}

	return arg + 3, nil /* Nil indicates no error */
}

/* it's possible to use custom types as errors by implementing the Error() method */
type argError struct {
	arg int
	prob string
}
/* this is a variant  */
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
/* In this case we use &argError syntax to build a new struct*/
func f2(arg int) (int, error){
	if arg == 42 {
		return -1, &argError{arg, "cant work with it"}
	}
	return arg+3, nil
}

func main() {
	fmt.Println("I work")
/* This loop test out each error returning functions
note that the use of an inline error check on the if line is a common idiom in Go*/
	for _, i := range [] int{7,42} {
		if r,e := f1(i); e!= nil{
			fmt.Println("f1 failed: ",e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7,42}{
		if r,e := f2(i); e != nil{
			fmt.Println("f2 failed: ", e)
		} else{
			fmt.Println("f2 worked: ", r)
		}
	}

	/* If its neccesary to use the error data in the program, create an instance if the custom error type via 
		type assetion*/

	_,e := f2(42)

	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
