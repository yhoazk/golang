package main

import "fmt"
/* Maps are associative data types, sometimes called hashes or dictionaries */
func main() {
	/* To create an empty map, use the make built in  */
	/* make(map[<key-type>]<val-type>) */
	m := make(map[string] int) /*<- watch NO comma*/
	/* set key/value pairs  */
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
	/* get a value for a key with name[key] as in python */
	v1 := m["k1"]
	fmt.Println("v1:", v1)
	/* The builtin len returns the number of key/value pairs when called on a map.  */
	fmt.Println("len:", len(m))
/* The builtin delete removes key/value pairs from a map.*/ 
	delete(m, "k2")

	fmt.Println("len after del:", len(m))
	/* The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didnt need the value itself, so we ignored it with the blank identifier _. */

	_, prs := m["k2"]

	fmt.Println("prs:", prs)

	n := map[string]int {
		"foo":1,
		"bar":13,
	}

	fmt.Println("map:", n)

}
