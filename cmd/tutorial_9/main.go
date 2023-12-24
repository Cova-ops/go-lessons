package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_CHICKEN_PRICE = 5
const MAX_TOFU_PRICE = 3

func process(c chan int) {
	defer close(c)
	for i := 5; i < 10; i++ {
		c <- i
	}
	fmt.Println("Existing process")
}

func checkTofuPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			c <- website
			break
		}
	}
}

func checkChickenPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			c <- website
			break
		}
	}
}

func sendMesssage(chickenChannel chan string, tofuChannel chan string) {
	// This will block until a message is received from one of the channels, but it will only receive one message
	select {
	case website := <-chickenChannel:
		fmt.Println("Chicken is available at ", website)
	case website := <-tofuChannel:
		fmt.Println("Tofu is available at ", website)
	}
}

func main() {
	var c = make(chan int, 5)
	go process(c)

	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}

	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)

	var websites = []string{"walmart.com", "amazon.com", "target.com", "costco.com", "samsclub.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMesssage(chickenChannel, tofuChannel)
}
