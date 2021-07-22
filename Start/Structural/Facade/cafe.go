package main

import "fmt"

func makeAmericano(size float32) {
	fmt.Println("\nMaking an Americano\n--------------------")

	// TODO: make an americano coffee using the coffeemachine API

	// determine beans amount to use - 5oz for every 8oz size

	fmt.Println("Americano is ready!")
}

func makeLatte(size float32, foam bool) {
	fmt.Println("\nMaking a Latte\n--------------------")

	// TODO: make a latte coffee using the coffeemachine API

	// determine beans amount to use - 5oz for every 8oz size

	// determine milk amount to use - 2oz for every 8oz size

	fmt.Println("Latte is ready!")
}
