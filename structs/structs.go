package main

import "fmt"
/* Elements are declared one per line   */
type person struct{
	name string
	age int
}

func main(){
	fmt.Println(person{"bob",	20})
	/* You can name the fields when initializing a struct. */
	fmt.Println(person{name:"tty",	age:20})
	/* Omitted fields will be zero-valued. */
	fmt.Println(person{name:"rd"})

	p := person{		name: "tacp",		age: 55545	}
	/* Access struct fields with a dot. */

	fmt.Println(p.name)

	/* You can also use dots with struct pointers - the pointers are automatically dereferenced. */
	pp := &p
	fmt.Println(pp.age)
	/* Structs are mutable */

	pp.age = 65465465
	fmt.Println(pp.age)



}
