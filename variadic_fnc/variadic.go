package main



import "fmt"
/* Variadic functions take an arbitrary number of inargs */
func sum(nums ... int){
	fmt.Print(nums, " ")
	total := 0
	for _, num:= range nums{
		total += num
	}
	fmt.Println(total)
}


func main(){
	/* Variadic functions can becalled in the usual way with individual args */
	sum(1,2)
	sum(1,2,3)
/* If there is already an slice  with the arguments, apply them to a variadic function */

	nums := []int{1,2,3,4,5}
	sum(nums...)
}
