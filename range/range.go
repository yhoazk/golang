package main

import "fmt"
/* range iterates over elements in a variety of data
	structures. */
func main(){
	nums := []int{2,3,4}

	/* range provides both, the index and the value for each entry */
	sum := 0
	for _, num :=range nums{
		sum += num
	}
	fmt.Println("sum:", sum)


	for i, num := range nums{
		if num == 3{
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{
		"a":"apple",
		"b":"banana",
	}

	for k,v := range kvs {
		fmt.Printf("%s -> %s\n", k,v)
	}

	for k := range kvs {

		fmt.Println("key:", k)
	}

	for i,c := range "go" {
		fmt.Println(i,c)
	}
}
