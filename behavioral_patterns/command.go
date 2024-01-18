package behavioral_patterns

import "fmt"

type Device interface {
	on()
	off()
}

type TV struct {
	isRunning bool
}

func (tv *TV) on() {
	fmt.Printf("is Tv running? %v\n", tv.isRunning)
	tv.isRunning = true
	fmt.Println("Tv is on now.")
}

func (tv *TV) off() {
	fmt.Printf("is Tv running? %v\n", tv.isRunning)
	tv.isRunning = false
	fmt.Println("Tv is off now.")
}

type Command interface {
	execute()
}

type OnCommand struct {
	device Device
}

func (cmd *OnCommand) execute() {
	cmd.device.on()
}

type OffCommand struct {
	device Device
}

func (cmd *OffCommand) execute() {
	cmd.device.off()
}

type Button struct {
	name string
	cmd Command
}

func (btn *Button) press() {
	btn.cmd.execute()
}

func CommandPattern() {
	myTv := &TV{
		isRunning: false,
	}

	onCmd := &OnCommand{
		device: myTv,
	}

	offCmd := &OffCommand{
		device: myTv,
	}

	onBtn := &Button{
		name: "On",
		cmd: onCmd,
	}

	offBtn := &Button{
		name: "Off",
		cmd: offCmd,
	}

	fmt.Printf("Pressing button %v\n", onBtn.name)
	onBtn.press()

	fmt.Printf("Pressing button %v\n", offBtn.name)
	offBtn.press()

	fmt.Printf("Pressing button %v\n", onBtn.name)
	onBtn.press()
}