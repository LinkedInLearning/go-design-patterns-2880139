package main

import "fmt"

// The CoffeeMachine struct represents an API to a hypothetical coffee maker
type CoffeeMachine struct {
	beanAmount   float32 // amount in ounces of beans to use
	grinderLevel int     // the granularity of the bean grinder
	waterTemp    int     // temperature of the water to use
	waterAmt     float32 // amount of water to use
	milkAmount   float32 // amount of milk to use
	addFoam      bool    // whether to add foam or not
}

func (c *CoffeeMachine) startCoffee(beanAmount float32, grind int) {
	c.beanAmount = beanAmount
	c.grinderLevel = grind
	fmt.Println("Starting coffee order with beans:", beanAmount, "and grind level", c.grinderLevel)
}

func (c *CoffeeMachine) endCoffee() {
	fmt.Println("Ending coffee order")
}

func (c *CoffeeMachine) grindBeans() bool {
	fmt.Println("Grinding the beans:", c.beanAmount, "beans at", c.grinderLevel, "granularity")
	return true
}

func (c *CoffeeMachine) useMilk(amount float32) float32 {
	fmt.Println("Adding milk:", amount, "oz")
	c.milkAmount = amount
	return amount
}

func (c *CoffeeMachine) setWaterTemp(temp int) {
	fmt.Println("Setting water temp:", temp)
}

func (c *CoffeeMachine) useHotWater(amount float32) float32 {
	fmt.Println("Adding hot water:", amount)
	c.waterAmt = amount
	return amount
}

func (c *CoffeeMachine) doFoam(useFoam bool) {
	fmt.Println("Foam setting:", useFoam)
	c.addFoam = useFoam
}
