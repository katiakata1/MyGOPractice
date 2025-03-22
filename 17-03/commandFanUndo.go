package main

import "fmt"

// Receiver - the object that performs an actual action

const (
	Low    = 1
	Medium = 2
	High   = 3
	Off    = 0
)

type CeilingFan struct {
	speed int
}

func (c *CeilingFan) SetSpeed(speed int) {
	c.speed = speed
	fmt.Println("Speed is set to: ", c.GetSpeed())
}

func (c *CeilingFan) GetSpeed() int {
	switch c.speed {
	case High:
		return High
	case Medium:
		return Medium
	case Low:
		return Low
	default:
		return Off
	}
}

// Command interface - mediator between Receiver and Invoker
type Command3 struct {
	execute   func()
	undo      func() int
	prevSpeed int
}

// Invoker

type Remote struct {
	command Command3
}

func (r *Remote) SetCommand(command Command3) {
	r.command = command
}

func (r *Remote) PressButton() {
	if r.command.execute != nil {
		r.command.prevSpeed = r.command.undo()
		r.command.execute()
	}
}

func (r *Remote) PressUndo() {
	if r.command.undo != nil {
		fanSpeed := r.command.prevSpeed
		r.command.execute = func() {
			fmt.Println("Undoing previous speed")
			r.command.prevSpeed = fanSpeed
		}
	}
}

func main() {
	fan := &CeilingFan{}
	remote := &Remote{}

	highSpeed := Command3{
		execute: func() { fan.SetSpeed(High) },
		undo: func() int {
			prev := fan.GetSpeed()
			fan.SetSpeed(0)
			return prev
		},
	}

	mediumSpeed := Command3{
		execute: func() { fan.SetSpeed(Medium) },
		undo: func() int {
			prev := fan.GetSpeed()
			fan.SetSpeed(0)
			return prev
		},
	}

	remote.SetCommand(highSpeed)
	remote.PressButton()
	remote.PressUndo()

	remote.SetCommand(mediumSpeed)
	remote.PressButton()
	remote.PressUndo()
}
