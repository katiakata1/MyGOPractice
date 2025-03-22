package main

import "fmt"

type Command1 interface {
	execute()
}

// Implementing a placeholder
type NoCommand struct{}

func (n *NoCommand) execute() {
	// Does nothing
}

type Ligth1 struct {
	location string
}

func (l *Ligth1) On1() {
	fmt.Println("Light on")
}

func (l *Ligth1) Off1() {
	fmt.Println("Light Off")
}

type Fan struct {
	location string
}

func (f *Fan) On() {
	fmt.Println("Fan is on")
}
func (f *Fan) Off() {
	fmt.Println("Fan is off")
}

type LightOnCommand1 struct {
	light1 *Ligth1
}

func (l *LightOnCommand1) execute() {
	l.light1.On1()
}

type LightOffCommand1 struct {
	light1 *Ligth1
}

func (l *LightOffCommand1) execute() {
	l.light1.Off1()
}

type FanOnCommand struct {
	fan *Fan
}

func (f *FanOnCommand) execute() {
	f.fan.On()
}

type FanOffCommand struct {
	fan *Fan
}

func (f *FanOffCommand) execute() {
	f.fan.Off()
}

// We need to extend remote control to allow multiple commands
type RemoteControl1 struct {
	onCommands  [7]Command1
	offCommands [7]Command1
}

func NewRemoteControl() *RemoteControl1 {
	remote := &RemoteControl1{}
	noCommand := &NoCommand{}

	for i := 0; i < 7; i++ {
		remote.onCommands[i] = noCommand
		remote.offCommands[i] = noCommand
	}

	return remote
}

// SetCommand assigns commands to a slot
func (r *RemoteControl1) SetCommand(slot int, onCommand, offCommand Command1) {
	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
}

func (r *RemoteControl1) PressOnButtom(slot int) {
	fmt.Println("Pressing ON button for slot %d", slot)
	r.onCommands[slot].execute()
}

func (r *RemoteControl1) PressOffButton(slot int) {
	fmt.Println("Pressing OFF Button")
	r.offCommands[slot].execute()
}

// func main() {
// 	remote := NewRemoteControl()

// 	// Create light and fan receivers
// 	livingRoomLight := &Ligth1{location: "Living Room"}
// 	kitchenLight := &Ligth1{location: "Kitchen"}
// 	bedroomFan := &Fan{location: "Bedroom"}

// 	// Creating corresponding commands
// 	livingRoomLightOn := &LightOnCommand1{light1: livingRoomLight}
// 	livingRoomLightOff := &LightOffCommand1{light1: livingRoomLight}

// 	kitchenLightOn := &LightOnCommand1{light1: kitchenLight}
// 	kitchenLightOff := &LightOffCommand1{light1: kitchenLight}

// 	bedroomFanOn := &FanOnCommand{fan: bedroomFan}
// 	bedroomFanOff := &FanOffCommand{fan: bedroomFan}

// 	// Assign Commands to remote slot
// 	remote.SetCommand(0, livingRoomLightOn, livingRoomLightOff)
// 	remote.SetCommand(1, kitchenLightOn, kitchenLightOff)
// 	remote.SetCommand(1, bedroomFanOn, bedroomFanOff)

// 	// Simulate button presses
// 	remote.PressOnButtom(0)
// 	remote.PressOffButton(0)
// 	remote.PressOnButtom(1)
// 	remote.PressOffButton(1)
// 	remote.PressOnButtom(2)
// 	remote.PressOffButton(2)

// 	// Pressing unassigned slot (should do nothing)
// 	remote.PressOnButtom(3)
// 	remote.PressOffButton(3)

// }
