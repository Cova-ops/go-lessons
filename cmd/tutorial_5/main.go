package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"   // UTF-8
	var indexed = myString[0] // byte

	fmt.Printf("%v, %T, %c\n", indexed, indexed, indexed) // 114, uint8, r
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("The length of 'myString' is %v\n", len(myString)) // 8 bytes

	var myString2 = []rune(myString)                                // rune is an alias for int32
	fmt.Printf("The length of 'myString2' is %v\n", len(myString2)) // 6 runes

	var myRune = 'á'
	fmt.Printf("myRune = %v\n", myRune)

	// not efficient
	var strSlice = []string{"c", "h", "a", "r", "a", "c", "t", "e", "r"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("%v\n", catStr) // character
	// catStr[0] = "C" // cannot assign to catStr[0]

	// this is the correct way to do it
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Printf("%v\n", catStr2) // character

}
