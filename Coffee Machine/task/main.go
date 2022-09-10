package main

import (
	"fmt"
	"strconv"
)

type CoffeeMachine struct {
	waterML        int
	milkML         int
	coffeeBeansGR  int
	disposableCups int
	money          int
}

var coffeeResources = make(map[int]CoffeeMachine)

func (machine *CoffeeMachine) GetStatus() string {
	a := fmt.Sprintf("The coffee machine has:\n%d ml of water\n", machine.waterML)
	b := fmt.Sprintf("%d ml of milk\n%d g of coffee beans\n", machine.milkML, machine.coffeeBeansGR)
	c := fmt.Sprintf("%d disposable cups\n$%d of money", machine.disposableCups, machine.money)
	return fmt.Sprintf("%s%s%s", a, b, c)
}

func (machine *CoffeeMachine) BuyExpresso() {
	machine.waterML -= 250
	machine.coffeeBeansGR -= 16
	machine.money += 4
	machine.disposableCups -= 1
}

func (machine *CoffeeMachine) BuyLatte() {
	machine.waterML -= 350
	machine.milkML -= 75
	machine.coffeeBeansGR -= 20
	machine.money += 7
	machine.disposableCups -= 1
}

func (machine *CoffeeMachine) BuyCappuccino() {
	machine.waterML -= 200
	machine.milkML -= 100
	machine.coffeeBeansGR -= 12
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

func (machine *CoffeeMachine) HasEnoughResources(coffeeType int) bool {
	resource, exists := coffeeResources[coffeeType]

	if exists {
		if machine.waterML < resource.waterML {
			fmt.Println("Sorry, not enough water!")
			return false
		}
		if machine.milkML < resource.milkML {
			fmt.Println("Sorry, not enough milk!")
			return false
		}
		if machine.coffeeBeansGR < resource.coffeeBeansGR {
			fmt.Println("Sorry, not enough coffee beans!")
			return false
		}
		if machine.disposableCups == 0 {
			fmt.Println("Sorry, not enough disposable cups!")
			return false
		}
		return true
	} else {
		fmt.Println("Type of Coffee Choosed DOES NOT EXIST")
		return false
	}

}

const (
	Buy       string = "buy"
	Fill             = "fill"
	Take             = "take"
	Remaining        = "remaining"
	Exit             = "exit"
)

const (
	Expresso   int = 1
	Late           = 2
	Cappuccino     = 3
)

func initCoffeeResources() {
	coffeeResources[Expresso] = CoffeeMachine{
		waterML:        250,
		milkML:         0,
		coffeeBeansGR:  16,
		money:          4,
		disposableCups: 1,
	}

	coffeeResources[Late] = CoffeeMachine{
		waterML:        350,
		milkML:         75,
		coffeeBeansGR:  20,
		money:          7,
		disposableCups: 1,
	}

	coffeeResources[Cappuccino] = CoffeeMachine{
		waterML:        200,
		milkML:         100,
		coffeeBeansGR:  12,
		money:          6,
		disposableCups: 1,
	}
}

func main() {
	machine := CoffeeMachine{
		waterML:        400,
		milkML:         540,
		coffeeBeansGR:  120,
		disposableCups: 9,
		money:          550,
	}

	initCoffeeResources()

	for {
		var action = ""
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		fmt.Scan(&action)

		if action == Buy {
			buyCoffeeAction(&machine)
		} else if action == Fill {
			fillAction(&machine)
		} else if action == Take {
			moneyTaken := machine.TakeMoney()
			fmt.Printf("I gave you $%d\n", moneyTaken)
		} else if action == Remaining {
			fmt.Println(machine.GetStatus())
		} else if action == Exit {
			break
		} else {
			fmt.Println("Action does not exist")
		}
	}
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
	fmt.Println("What do you want to buy? 1 - expresso, 2 - latte, 3 - cappucino, back - to main menu")
	var optionChosed = ""
	fmt.Scan(&optionChosed)

	if optionChosed != "back" {

		coffeeType, _ := strconv.Atoi(optionChosed)

		if machine.HasEnoughResources(coffeeType) {
			fmt.Println("I have enough resources, making you a coffee!")
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
	}

}
