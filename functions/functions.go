package main

import "fmt"


func test()int{
	i := 45
	i++
	return i
}

func main(){
	fmt.Println(test())
}
