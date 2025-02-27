package main

import "fmt"

// Create interfaces

type FlyBehavour interface {
	Fly()
}

type QuackBehavour interface {
	Quack()
}

// Implement different behavours

type FlyWithWings struct{}

func (f *FlyWithWings) Fly() {
	fmt.Println("I fly with wings")
}

type NoFly struct{}

func (f *NoFly) Fly() {
	fmt.Println("I cannot fly")
}

type Quack struct{}

func (q *Quack) Quack() {
	fmt.Println("I can quack")
}

type NoQuack struct{}

func (q *NoQuack) Quack() {
	fmt.Println("I cannot quack")
}

// Now create a Duck struct and modify it to use these behaviours

type Duck struct {
	flyBehaviour   FlyBehavour
	quackBehaviour QuackBehavour
}

func (d *Duck) PerformFly() {
	d.flyBehaviour.Fly()
}

func (d *Duck) PerformQuack() {
	d.quackBehaviour.Quack()
}

// create setters that can change the behaviour during runtime
func (d *Duck) SetFlyBehaviour(fb FlyBehavour) {
	d.flyBehaviour = fb
}

func (d *Duck) SetQuackBehaviour(qb QuackBehavour) {
	d.quackBehaviour = qb
}

// Now Create a specific duck type
type MallardDuck struct {
	Duck
}

func NewMallardDuck() *MallardDuck {
	return &MallardDuck{
		Duck{
			flyBehaviour:   &FlyWithWings{},
			quackBehaviour: &Quack{},
		},
	}
}

type RubberDuck struct {
	Duck
}

func NewRubberDuck() *RubberDuck {
	return &RubberDuck{
		Duck{
			flyBehaviour:   &NoFly{},
			quackBehaviour: &NoQuack{},
		},
	}
}

// Now test it

func main() {
	mallard := NewMallardDuck()
	fmt.Println("This is MallardDuck")
	mallard.PerformFly()
	mallard.PerformQuack()

	rubber := NewRubberDuck()
	fmt.Println("This is Rubber Duck")
	rubber.PerformFly()
	rubber.PerformQuack()

	// dynamically change behaviour
	fmt.Println("\nChanging Rubber Duck's Fly Behavior...")
	rubber.SetFlyBehaviour(&FlyWithWings{})
	rubber.PerformFly()
}
