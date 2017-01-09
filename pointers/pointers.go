package main


/* Go also has pointers, and they work pretty much as C pointers */


import "fmt"

func zeroval(ival int){

	ival = 0
}

func zeroptr(iptr *int){
	*iptr = 0
}

func main(){
	i := 1
	fmt.Println("initial: ",i)

	zeroval(i)
	fmt.Println("Zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr: ", i)


	fmt.Println("ptr: ", &i)

}
