package main

import "fmt"

func min3(n1 int, n2 int, n3 int) int {
	var res = 0
	if n1 <= n2 && n1 <= n3 {
		res = n1
	} else if n2 <= n1 && n2 <= n3 {
		res = n2
	} else {
		res = n3
	}
	return res
}

func main() {
	var waterAvailableMl = 0
	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(&waterAvailableMl)

	var milkAvailableML = 0
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(&milkAvailableML)

	var coffeeAvailableGrams = 0
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scan(&coffeeAvailableGrams)

	cups := 0
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&cups)

	var (
		waterUnit         = 200
		milkUnit          = 50
		coffeeBeansByUnit = 15
	)

	requiredWater := waterUnit * cups
	requiredMilk := milkUnit * cups
	requiredCoffe := coffeeBeansByUnit * cups

	var maxCupWater = waterAvailableMl / waterUnit
	var maxCupMilk = milkAvailableML / milkUnit
	var maxCoffee = coffeeAvailableGrams / coffeeBeansByUnit
	var minOfMaxs = min3(maxCupWater, maxCupMilk, maxCoffee)

	if waterAvailableMl >= requiredWater && milkAvailableML >= requiredMilk && coffeeAvailableGrams >= requiredCoffe {
		fmt.Print("Yes, I can make that amount of coffee")
		if minOfMaxs > cups {
			fmt.Printf(" (and even %d more than that)\n", minOfMaxs-cups)
		}
		fmt.Printf("\n")
	} else {
		fmt.Printf("No, I can make only %d cups of coffee\n", minOfMaxs)
	}
}
