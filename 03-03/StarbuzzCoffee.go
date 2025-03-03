package main

import "fmt"

// Component Interface

type Beverage interface {
	Cost() float64
	Description() string
	SetSize(size string)
	GetSize() string
}

// size struct which we want to share with base coffees
type BaseBeverage struct {
	description string
	size        string
}

func (bb *BaseBeverage) GetSize() string {
	return bb.size
}

func (bb *BaseBeverage) SetSize(size string) {
	bb.size = size
}

// Concrete component

type HouseBlend struct {
	BaseBeverage
}

func (h *HouseBlend) Cost() float64 {
	baseCost := 2.00
	if h.size == "Medium" {
		baseCost += 0.5
	} else if h.size == "Large" {
		baseCost += 1.00
	}
	return baseCost
}

func (h *HouseBlend) Description() string {
	return h.size + " HouseBlend coffee"
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

// forward size-related methods to the wrapped beverages
func (m *Milk) GetSize() string {
	return m.beverage.GetSize()
}

func (m *Milk) SetSize(size string) {
	m.beverage.SetSize(size)
}

// Mocha
type Mocha struct {
	beverage Beverage
}

func (m *Mocha) Cost() float64 {
	return m.beverage.Cost() + 0.75
}

func (m *Mocha) Description() string {
	return m.beverage.Description() + ", Mocha"
}

func (m *Mocha) GetSize() string {
	return m.beverage.GetSize()
}

func (m *Mocha) SetSize(size string) {
	m.beverage.SetSize(size)
}

// Whip
type Whip struct {
	beverage Beverage
}

func (w *Whip) Cost() float64 {
	return w.beverage.Cost() + 0.25
}

func (w *Whip) Description() string {
	return w.beverage.Description() + ", Whip"
}

func (m *Whip) GetSize() string {
	return m.beverage.GetSize()
}

func (m *Whip) SetSize(size string) {
	m.beverage.SetSize(size)
}

func main() {
	beverage := &HouseBlend{}
	beverage.SetSize("Large")
	fmt.Println(beverage.Description(), "Cost: ", beverage.Cost())

	withMilk := &Milk{beverage}
	fmt.Println(withMilk.Description(), "Cost: ", withMilk.Cost())

	withMilkMocha := &Mocha{withMilk}
	fmt.Println(withMilkMocha.Description(), "Cost: ", withMilkMocha.Cost())

	withWhip := &Whip{withMilkMocha}
	fmt.Println(withWhip.Description(), "Cost: ", withWhip.Cost())
}
