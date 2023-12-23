package main

import (
	"fmt"
	"time"
)

func timeLoop(slice []int, n int) time.Duration {
	var start = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(start)
}

func main() {
	var intArr [3]int32
	intArr[1] = 123

	fmt.Println(intArr[0])   // 0
	fmt.Println(intArr[1:3]) // [123 0]

	fmt.Println(&intArr[0]) // 0xc0000a4000
	fmt.Println(&intArr[1]) // 0xc0000a4004
	fmt.Println(&intArr[2]) // 0xc0000a4008

	// var intArr2 [3]int32 = [3]int32{1, 2, 3}
	// intArr2 := [3]int32{1, 2, 3}
	// intArr2 := [...]int32{1, 2, 3} // --> The length is inferred from the elements

	// Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data.
	var intSlice []int32 = []int32{4, 5, 6}

	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice)) // The length is 3 with capacity 3
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice)) // The length is 4 with capacity 6
	// The capacity has increased to 6 because the underlying array has been copied to a new location with more capacity

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice) // [4 5 6 7 8 9]

	var intSlice3 []int32 = make([]int32, 3, 5) // make([]T, length, capacity)
	fmt.Println(intSlice3)                      // [0 0 0]

	var myMap map[string]uint32 = make(map[string]uint32)
	fmt.Println(myMap) // map[]

	var myMap2 map[string]uint32 = map[string]uint32{"Adam": 23, "Eve": 19}
	fmt.Println(myMap2)         // map[Adam:23 Eve:19]
	fmt.Println(myMap2["Adam"]) // 23
	fmt.Println(myMap2["John"]) // 0 --> The value is 0 because the key does not exist, 0 is the zero value for uint32

	var age, ok = myMap2["Adam"] // ok is a boolean that indicates if the key exists
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("The key does not exist")
	}

	// delete(myMap2, "Adam") // Delete the key Adam

	for name, age := range myMap2 { // range can be use to iterate over maps, slices, arrays, strings and channels
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i, v := range intArr { // i is the index and v is the value
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	// while loop
	var i int = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}

	for i := 10; i < 15; i++ {
		fmt.Println(i)
	}

	var n int = 100000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n)) // Preallocation is much faster
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))   // Preallocation is much faster
}
