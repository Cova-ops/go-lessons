package main

import "fmt"

type owner struct {
	name string
}

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

type gasEngine2 struct {
	mpg     uint8
	gallons uint8
	owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// func milesLeft(e gasEngine) uint8 { // this is a function
func (e gasEngine) milesLeft() uint8 { // this is a method
	return e.mpg * e.gallons
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	// var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15, ownerInfo: owner{name: "Alex"}} // this is valid
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}}            // this is also valid
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name) // 25 15 Alex

	myEngine.mpg = 20
	myEngine.gallons = 10
	myEngine.ownerInfo.name = "John"
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name) // 20 10 John
	fmt.Println("Total miles left in the tank:", myEngine.milesLeft())   // Total miles left in the tank: 200

	var myEngine2 gasEngine2 = gasEngine2{25, 15, owner{"Alex"}}
	fmt.Println(myEngine2.mpg, myEngine2.gallons, myEngine2.name) // 25 15 Alex

	var myEngine3 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	fmt.Println(myEngine3.mpg, myEngine3.gallons) // 25 15

	var myEngine4 electricEngine = electricEngine{5, 10}
	fmt.Println(myEngine4.mpkwh, myEngine4.kwh) // 5 10

	// canMakeIt(myEngine, 20) // You can make it there!
	canMakeIt(myEngine4, 20) // You can make it there!
}
