package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// var intNum int16 = 32767 + 1 --> this will cause an error
	var intNum int = 32767
	fmt.Println(intNum)

	// var floatNum float32 = 1234567.9 --> this will return 123456789.000
	// var floatNum float64 = 1234567.9 --> this will return 12345678.9000

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	// var result float32 = floatNum32 + intNum32 --> this will cause an error because we can't add float32 and int32
	var result int32 = int32(floatNum32) + intNum32
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2) // this will return 1
	fmt.Println(intNum1 % intNum2) // this will return 1

	var myString string = "Hello \nWorld!" + " " + "I'm here!"
	fmt.Println(myString)

	var myString2 string = `Hello
World!`
	fmt.Println(myString2)

	fmt.Println(len("Hello World!"))                    // this will return 12. Its returned value is the number of bytes in the string
	fmt.Println(utf8.RuneCountInString("Hello World!")) // this will return 12. Its returned value is the number of characters in the string

	var myRune rune = 'a'
	fmt.Println(myRune) // this will return 97

	var myBoolean bool = true
	fmt.Println(myBoolean) // this will return true

	var intNum3 int
	fmt.Println(intNum3) // this will return 0

	const myConst string = "const value"
	// myConst = "new const value" --> this will cause an error because we can't change the value of a constant
	fmt.Println(myConst)
}
