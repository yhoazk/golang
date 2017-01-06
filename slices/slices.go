package main


import "fmt"

/* 	
	Slices are a key datatype in GO, givin a more
	powwerful interface to sequences than arrays
*/

func main(){
	/* unlike arrays slices are typed only by 
	the elements they contain, notthe numberof 
	elements*/
	/*create a sclice of strings size =3 */
	s := make([]string,3)
	fmt.Println("emp:",s)

	/*Set and get is the same for slices and for arrays*/
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:",s)
	fmt.Println("get:",s[2])

	fmt.Println("len:", len(s))
	/* 	Slices also implement the append method
		this method returns a slice containing
		a slice containing one or more new values
		NOTE: WE need to accept a return value from
		append as we may get a new slice value
	*/
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	/* slices can also be copyed */
	/* Here we create an empty copy of s*/
	c := make([]string, len(s))
	copy(c,s)
	fmt.Println("cpy:", c)

	/* we can get slices of slices as we do in python */
	/* This slices up to (but excluding) s[5]. */
	l := s[:5]
	fmt.Println("sl1:", l)
/* And this slices up from (and including) s[2]. */
	l = s[2:]
	fmt.Println("sl2:", l)

	l = s[2:5]
	fmt.Println("sl3:", l)
	/*  We can declare and initialize a variable for slice in a single line as well. */
	t := []string{"g", "f", "y"}
	fmt.Println("dcl:", t)

	/* 	Multidimensional data structs
		The length of the inner slice can vary, unlike with arrays
	*/
	const length = 3
	twoD := make([][]int, length)
	for i:=0; i<length; i++{
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j:=0; j < innerLen;j++{
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d: ", twoD)

	arr := [5]int{1,2,3,4,5}

	for i:=0; i<4;i++{
		fmt.Println(arr[i])
	}
}
