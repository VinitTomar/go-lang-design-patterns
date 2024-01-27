package structural_patterns

import "fmt"

type iPizza interface {
	getPrice() int
}

type veggiePizza struct {}

func (sp *veggiePizza) getPrice() int {
	return 17
}

type tomatoTopingPizza struct {
	pizza iPizza
}

func (ttp *tomatoTopingPizza) getPrice() int {
	return ttp.pizza.getPrice() + 3
}

type cheeseToppingPizza struct {
	pizza iPizza
}

func (ctp *cheeseToppingPizza) getPrice() int {
	return ctp.pizza.getPrice() + 5
}

func DecoratorPattern() {
	myPizza := &veggiePizza{}

	// Add tomato topping
	myPizzaWithTomatoToppings := &tomatoTopingPizza{myPizza}

	// Add cheese topping
	myPizzaWithTomatoAndCheeseToppings := &cheeseToppingPizza{myPizzaWithTomatoToppings}

	fmt.Printf("Total price of pizza with tomato and cheese toppings is %v\n", myPizzaWithTomatoAndCheeseToppings.getPrice())
}