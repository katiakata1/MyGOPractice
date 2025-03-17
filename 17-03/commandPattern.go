package main

import "fmt"

// Create interface
type Command interface {
	execute()
}

// Create Receiver (Light) and its methods
type Light struct{}

func (l *Light) On() {
	fmt.Println("Switching on the light")
}

func (l *Light) Off() {
	fmt.Println("Switchingoff the light")
}

// Implement a concrete command
type LigthOnCommand struct {
	light *Light
}

func (l *LigthOnCommand) execute() {
	l.light.On()
}

// Invoker asks command to execute
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) PressButton() {
	r.command.execute()
}

func main() {
	// Create receiver
	light := &Light{}

	// Create concrete command with receiver
	lightOn := &LigthOnCommand{light: light}

	// Create invoker and assign command
	remote := &RemoteControl{command: lightOn}

	// Press button to execute the command
	remote.PressButton()
}
