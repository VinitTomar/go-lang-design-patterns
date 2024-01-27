package structural_patterns

import "fmt"

type printer interface {
	printFile()
}

type hpPrinter struct {}

func (hp hpPrinter) printFile() {
	fmt.Println("Printing file from hp printer.")
}

type epsonPrinter struct {}

func (ep epsonPrinter) printFile() {
	fmt.Println("Printing file from epson printer.")
}

type macComputer struct {
	prt printer
}

func (mc *macComputer) print() {
	fmt.Println("Printing from mac computer.")
	mc.prt.printFile()
}

func (mc *macComputer) setPrinter(p printer) {
	mc.prt = p
}

type windowsComputer struct {
	prt printer
}

func (mc *windowsComputer) print() {
	fmt.Println("Printing from windows computer.")
	mc.prt.printFile()
}

func (mc *windowsComputer) setPrinter(p printer) {
	mc.prt = p
}

func BridgePattern() {
	hp := hpPrinter{}
	epson := epsonPrinter{}

	mac := macComputer{}
	mac.setPrinter(hp)
	mac.print()

	mac.setPrinter(epson)
	mac.print()

	win := windowsComputer{}
	win.setPrinter(hp)
	win.print()

	win.setPrinter(epson)
	win.print()
}