package main

import "fmt"

type cleaner struct {
	cleaningMethod string
	tools          []string
}

type coffeeMachine struct {
	cleaner cleaner
}

type cleanable interface {
	clean()
	setTools(tools []string)
}

func (c *cleaner) setTools(tools []string) {
	c.tools = tools
}

func (c *cleaner) clean() {
	fmt.Printf("%s cleaning with the following tools: \n", c.cleaningMethod)
	for _, tool := range c.tools {
		fmt.Println(tool)
	}
	fmt.Println()
}

func newCoffeeMachine(cleaner *cleaner) *coffeeMachine {
	return &coffeeMachine{
		cleaner: *cleaner,
	}
}

func main() {
	// Init the wet cleaner
	wetCleaner := &cleaner{
		cleaningMethod: "wet",
	}
	// Set the tools that are needed for the wet cleaning
	wetCleaner.setTools([]string{"water", "soap"})

	// Init the coffee machine with the wet cleaner
	coffeeMachine := newCoffeeMachine(wetCleaner)
	// Clean the coffee machine
	coffeeMachine.cleaner.clean()

	// Add a new tool to the wet cleaner
	coffeeMachine.cleaner.setTools([]string{"water", "soap", "sponge"})

	// Clean the coffee machine again with the wet cleaner and the new tools
	coffeeMachine.cleaner.clean()

	// Init the dry cleaner
	dryCleaner := cleaner{
		cleaningMethod: "dry",
	}

	// change the implementation of the cleaner
	coffeeMachine.cleaner = dryCleaner

	// Set the tools that are needed for the dry cleaning
	coffeeMachine.cleaner.setTools([]string{"cloth", "brush"})

	// Clean the coffee machine with the dry cleaner
	coffeeMachine.cleaner.clean()

}
