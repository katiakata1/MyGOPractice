package main

import "fmt"

type Command2 interface {
	execute()
	undo()
}

type Light2 struct{}

func (l *Light2) On() {
	fmt.Println("Light in ON")
}

func (l *Light2) Off() {
	fmt.Println("Light in OFF")
}

type CommandOnLight struct {
	light *Light2
}

func (c *CommandOnLight) execute() {
	c.light.On()
}

func (c *CommandOnLight) undo() {
	c.light.Off()
}

type CommandOffLight struct {
	light *Light2
}

func (c *CommandOffLight) execute() {
	c.light.Off()
}

func (c *CommandOffLight) undo() {
	c.light.On()
}

type RemoteControl2 struct {
	slot        Command2
	undoCommand Command2
}

func (r *RemoteControl2) setCommand(cmd Command2) {
	r.slot = cmd
}

func (r *RemoteControl2) PressButton() {
	r.slot.execute()
	r.undoCommand = r.slot
}

func (r *RemoteControl2) PressUndo() {
	if r.undoCommand != nil {
		r.undoCommand.undo()
	}
}

// func main() {
// 	light := &Light2{}

// 	onCommandLight := &CommandOnLight{light: light}
// 	// offCommandLight := &CommandOffLight{light: light}

// 	remote := &RemoteControl2{}
// 	remote.setCommand(onCommandLight)
// 	remote.PressButton()
// 	remote.PressUndo()

// }
