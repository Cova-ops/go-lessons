package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	var data, _ = os.ReadFile(filePath)

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, value := range slice {
		sum += value
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice(intSlice))
	fmt.Println(isEmpty(intSlice))

	var float32Slice = []float32{1.1, 2.2, 3.3}
	fmt.Println(sumSlice(float32Slice))
	fmt.Println(isEmpty(float32Slice))

	var contacts []contactInfo = loadJSON[contactInfo]("./cmd/tutorial_10/contactInfo.json")
	fmt.Printf("%+v\n", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./cmd/tutorial_10/purchaseInfo.json")
	fmt.Printf("%+v\n", purchases)
}
