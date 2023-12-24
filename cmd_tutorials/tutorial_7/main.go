package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("The value p points to is %v\n", *p) // The value p points to is 0
	fmt.Printf("The value of i is %v\n", i)         // The value of i is 0

	*p = 42
	fmt.Printf("The value p points to is %v\n", *p) // The value p points to is 42

	var z = i      // z is a copy as value of i
	i = 27         // i is now 27
	fmt.Println(i) // 27
	fmt.Println(z) // 0

	var slice = []int{1, 2, 3}
	var sliceCopy = slice // sliceCopy is a copy as reference of slice
	sliceCopy[2] = 4

	fmt.Println(slice)     // [1 2 4]
	fmt.Println(sliceCopy) // [1 2 4]
}
