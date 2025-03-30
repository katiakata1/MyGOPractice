package main

import "fmt"

type Duck interface {
	quack()
	fly()
}

type MallardDuck struct{}

func (m *MallardDuck) fly() {
	fmt.Println("I can fly")
}

func (m *MallardDuck) quack() {
	fmt.Println("I quack")
}

// Bird - Adaptee
type Bird struct{}

func (b *Bird) chirp() {
	fmt.Println("Chirp Chirp!")
}

func (b *Bird) flyHigh() {
	fmt.Println("Flying high in the sky!")
}

// BirdAdapter - Adapter to make Bird behave like a Duck
type BirdAdapter struct {
	bird *Bird
}

func (b *BirdAdapter) fly() {
	fmt.Print("Bird adapts: ")
	b.bird.flyHigh()
}

func (b *BirdAdapter) quack() {
	fmt.Print("Bird adapts: ")
	b.bird.chirp()
}

func main() {
	// create a duck
	duck := &MallardDuck{}
	fmt.Println("MallardDuck:")
	duck.fly()
	duck.quack()

	// create a bird
	bird := &Bird{}
	bird.chirp()
	bird.flyHigh()

	// Use BirdAdapter to make Bird act like a Duck
	adapterBird := &BirdAdapter{bird}
	adapterBird.fly()
	adapterBird.quack()
}
