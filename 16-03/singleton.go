package main

import (
	"fmt"
	"sync"
)

type ChocolateBoiler struct {
	empty  bool
	boiled bool
}

var instance *ChocolateBoiler
var once sync.Once

// Get Instance
func GetInstance() *ChocolateBoiler {
	once.Do(func() {
		fmt.Println("Creating a single instance of Chocolate Boiler")
		instance = &ChocolateBoiler{empty: true, boiled: false}
	})
	return instance
}

// Fills the boiler with chocolate and milk
func (b *ChocolateBoiler) Fill() {
	if b.empty {
		fmt.Println("The boiler is empty. Filling it up now")
		b.empty = false
		b.boiled = false
	} else {
		fmt.Println("The boiler is already filled")
	}
}

// Boil the content
func (b *ChocolateBoiler) Boil() {
	if !b.empty && !b.boiled {
		fmt.Println("Boiling now")
		b.boiled = true
	} else if b.empty {
		fmt.Println("The tank is empty")
	} else {
		fmt.Println("The tank is full and boiled")
	}
}

// Drain empties the tank if it is full and boiled
func (b *ChocolateBoiler) Drain() {
	if !b.empty && b.boiled {
		fmt.Println("Draining the tank")
		b.empty = true
	} else if b.empty {
		fmt.Println("Cannot drain. The tank is empty")
	} else {
		fmt.Println("Cannot drain, not yet boiled")
	}
}

func main() {
	boiler1 := GetInstance()
	boiler2 := GetInstance()

	//check if boiler1 and boiler2 are the same instances
	if boiler1 == boiler2 {
		fmt.Println("Both are the same instance!")
	}

	boiler1.Fill()
	boiler1.Boil()
	boiler1.Drain()

	boiler2.Fill()
}
