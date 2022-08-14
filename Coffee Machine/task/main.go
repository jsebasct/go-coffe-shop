package main

import "fmt"

func main() {
	fmt.Println("Write how many cups of coffee you will need:")

	cups := 0
	fmt.Scan(&cups)
	fmt.Printf("For %d cups of coffee you will need:\n", cups)

	var (
		waterUnit         = 200
		milkUnit          = 50
		coffeeBeansByUnit = 15
	)

	fmt.Printf("%d ml of water\n", waterUnit*cups)
	fmt.Printf("%d ml of milk\n", milkUnit*cups)
	fmt.Printf("%d g of coffee beans\n", coffeeBeansByUnit*cups)

}
