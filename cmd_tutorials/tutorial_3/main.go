package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)

	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("%d / %d = %d\n", numerator, denominator, result)
	} else {
		fmt.Printf("%d / %d = %d with remainder %d\n", numerator, denominator, result, remainder)
	}

	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("%d / %d = %d\n", numerator, denominator, result)
	default:
		fmt.Printf("%d / %d = %d with remainder %d\n", numerator, denominator, result, remainder)
	}

	switch remainder {
	case 0:
		printMe("The division wa exact")
	case 1, 2:
		printMe("The division was close")
	default:
		printMe("The division was not close")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("Cannot divide by zero")
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, nil
}
