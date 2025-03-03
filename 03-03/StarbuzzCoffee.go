package main

import "fmt"

// Component Interface

type Beverage interface {
	Cost() float64
	Description() string
}

// Concrete component

type HouseBlend struct{}

func (h *HouseBlend) Cost() float64 {
	return 2.00
}

func (h *HouseBlend) Description() string {
	return "HouseBlend coffee"
}

// Decorators wraps beverage and adds functionallity dinamically

type Milk struct {
	beverage Beverage
}

func (m *Milk) Cost() float64 {
	return m.beverage.Cost() + 0.5
}

func (m *Milk) Description() string {
	return m.beverage.Description() + ", Milk"
}

type Mocha struct {
	beverage Beverage
}

func (m *Mocha) Cost() float64 {
	return m.beverage.Cost() + 0.75
}

func (m *Mocha) Description() string {
	return m.beverage.Description() + ", Mocha"
}

type Whip struct {
	beverage Beverage
}

func (w *Whip) Cost() float64 {
	return w.beverage.Cost() + 0.25
}

func (w *Whip) Description() string {
	return w.beverage.Description() + ", Whip"
}

func main() {
	beverage := &HouseBlend{}
	fmt.Println(beverage.Description(), "Cost: ", beverage.Cost())

	withMilk := &Milk{beverage}
	fmt.Println(withMilk.Description(), "Cost: ", withMilk.Cost())

	withMilkMocha := &Mocha{withMilk}
	fmt.Println(withMilkMocha.Description(), "Cost: ", withMilkMocha.Cost())

	withWhip := &Whip{withMilkMocha}
	fmt.Println(withWhip.Description(), "Cost: ", withWhip.Cost())
}
