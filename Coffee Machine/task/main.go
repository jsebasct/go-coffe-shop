package main

import (
	"fmt"
)

type CoffeeMachine struct {
	waterML        int
	milkML         int
	coffeeBeansGR  int
	disposableCups int
	money          int
}

func (machine *CoffeeMachine) GetStatus() string {
	a := fmt.Sprintf("The coffee machine has:\n%d ml of water\n", machine.waterML)
	b := fmt.Sprintf("%d ml of milk\n%d g of coffee beans\n", machine.milkML, machine.coffeeBeansGR)
	c := fmt.Sprintf("%d disposable cups\n$%d of money", machine.disposableCups, machine.money)
	return fmt.Sprintf("%s%s%s", a, b, c)
}

func (machine *CoffeeMachine) BuyExpresso() {
	machine.waterML = machine.waterML - 250
	machine.coffeeBeansGR = machine.coffeeBeansGR - 16
	machine.money += 4
	machine.disposableCups -= 1
}

func (machine *CoffeeMachine) BuyLatte() {
	machine.waterML = machine.waterML - 350
	machine.milkML -= 75
	machine.coffeeBeansGR = machine.coffeeBeansGR - 20
	machine.money += 7
	machine.disposableCups -= 1
}

func (machine *CoffeeMachine) BuyCappuccino() {
	machine.waterML = machine.waterML - 200
	machine.milkML -= 100
	machine.coffeeBeansGR = machine.coffeeBeansGR - 12
	machine.money += 6
	machine.disposableCups -= 1
}

func (machine *CoffeeMachine) Refill(waterML, milkML, coffeeGR, cups int) {
	machine.waterML += waterML
	machine.milkML += milkML
	machine.coffeeBeansGR += coffeeGR
	machine.disposableCups += cups
}

func (machine *CoffeeMachine) TakeMoney() int {
	var money = machine.money
	machine.money = 0
	return money
}

const (
	Buy  string = "buy"
	Fill        = "fill"
	Take        = "take"
)

const (
	Expresso   int = 1
	Late           = 2
	Cappuccino     = 3
)

func main() {
	machine := CoffeeMachine{
		waterML:        400,
		milkML:         540,
		coffeeBeansGR:  120,
		disposableCups: 9,
		money:          550,
	}
	fmt.Println(machine.GetStatus())

	var action = ""
	fmt.Println("Write action (buy, fill, take):")
	fmt.Scan(&action)

	// TODO check if valid action
	if action == Buy {
		buyCoffeeAction(&machine)
	} else if action == Fill {
		fillAction(&machine)
	} else if action == Take {
		moneyTaken := machine.TakeMoney()
		fmt.Printf("I have you $%d/n", moneyTaken)
	}
	fmt.Println(machine.GetStatus())
}

func fillAction(machine *CoffeeMachine) {
	var waterAddedMl = 0
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&waterAddedMl)

	var milkAddedML = 0
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&milkAddedML)

	var coffeeAddedGrams = 0
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&coffeeAddedGrams)

	cups := 0
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&cups)

	machine.Refill(waterAddedMl, milkAddedML, coffeeAddedGrams, cups)
}

func buyCoffeeAction(machine *CoffeeMachine) {
	fmt.Println("What do you want to buy? 1 - expresso, 2 - latte, 3 - cappucino")
	var coffeeType = 0
	fmt.Scan(&coffeeType)

	// evaluate
	if coffeeType == Expresso {
		machine.BuyExpresso()
	} else if coffeeType == Late {
		machine.BuyLatte()
	} else if coffeeType == Cappuccino {
		machine.BuyCappuccino()
	} else {
		fmt.Println("unkonw action")
	}
}
