package main

import "fmt"

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}

type BasePizza struct {
	name string
}

func (p *BasePizza) Bake() {
	fmt.Println("Baking")
}

func (p *BasePizza) Cut() {
	fmt.Println("Cutting")
}

func (p *BasePizza) Box() {
	fmt.Println("Boxing")
}

// Creating NewYorkPizza struct
type NewYorkPizza struct {
	BasePizza
}

func NewNewYorkPizza() *NewYorkPizza {
	return &NewYorkPizza{BasePizza{name: "New York Pizza"}}
}

func (p *NewYorkPizza) Prepare() {
	fmt.Println("Preparing New York style ")
}

func (p *NewYorkPizza) GetName() string {
	return p.name
}

// Creating ChicagoPizza struct
type ChicagoPizza struct {
	BasePizza
}

func NewChicagoPizza() *ChicagoPizza {
	return &ChicagoPizza{BasePizza{name: "ChicagoPizza"}}
}

func (p *ChicagoPizza) Prepare() {
	fmt.Println("Preparing Chicago style ")
}

func (p *ChicagoPizza) GetName() string {
	return p.name
}

// Creating PizzaFactory

type PizzaStore interface {
	CreatePizza(pizzaType string) Pizza
	OrderPizza(pizzaType string) Pizza
}

// Subclass: NewYorkPizzaStore (subclass of PizzaStore)
type NYPizzaStore struct{}

func (c *NYPizzaStore) OrderPizza(pizzaType string) Pizza {
	pizza := c.CreatePizza(pizzaType)
	if pizza == nil {
		fmt.Println("Invalid pizza type")
	}

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	fmt.Println("Pizza is ready: ", pizza.GetName())
	return pizza
}

func (c *NYPizzaStore) CreatePizza(pizzaType string) Pizza {
	if pizzaType == "cheese" {
		return NewNewYorkPizza()
	}
	return nil
}

// Subclass: ChicagoPizzaStore (subclass of PizzaStore)
type ChicagoPizzaStore struct{}

func (c *ChicagoPizzaStore) OrderPizza(pizzaType string) Pizza {
	pizza := c.CreatePizza(pizzaType)

	if pizza == nil {
		fmt.Println("Invalid pizza type")
	}

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	fmt.Println("Pizza is ready: ", pizza.GetName())
	return pizza
}

func (c *ChicagoPizzaStore) CreatePizza(pizzaType string) Pizza {
	if pizzaType == "cheese" {
		return NewChicagoPizza()
	}
	return nil
}

func main() {
	nyStore := &NYPizzaStore{}
	chicagoStroe := &ChicagoPizzaStore{}

	nyStore.OrderPizza("cheese")
	chicagoStroe.OrderPizza("cheese")
}
