package structural_patterns

import "fmt"


type computer interface {
	insertIntoLightningPort()
}

type mac struct {}

func (m mac) insertIntoLightningPort() {
	fmt.Println("Lighting connector inserted into mac machine.")
}

type windows struct{}

func (w windows) insertIntoUsbPort() {
	fmt.Println("Usb connector is connected into windows machine.")
}

type windowsAdapter struct {
	windowsMachine *windows
}

func (wa windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts lighting port signal to usb port.")
	wa.windowsMachine.insertIntoUsbPort()
}

type client struct {}

func (c client) insertLightningConnectorToComputer(com computer) {
	fmt.Println("Client inserts lightning port into computer")
	com.insertIntoLightningPort()
}

func AdapterPattern() {
	clt := client{}
	mac := mac{}

	clt.insertLightningConnectorToComputer(mac)

	adp := windowsAdapter {
		&windows{},
	}

	clt.insertLightningConnectorToComputer(adp)
}