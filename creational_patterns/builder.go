package creational_patterns

import "fmt"

type iBuilder interface {
	reset()
	setWindowType()
	setDoorType()
	setNumFloor()
	makeHouseReady()
}

type house struct {
	windowType string
	doorType string
	floor int
}

type normalHouse struct {
	house
}

type normalBuilder struct {
	houseReady bool
	house *normalHouse
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{
		houseReady: false,
		house: &normalHouse{},
	}
}

func (nb *normalBuilder) reset() {
	nb.houseReady = false
	nb.house = &normalHouse{}
}

func (nb *normalBuilder) setWindowType() {
	nb.house.windowType = "Wooden windows"
}

func (nb *normalBuilder) setDoorType() {
	nb.house.doorType = "Wooden doors"
}

func (nb *normalBuilder) setNumFloor() {
	nb.house.floor = 4
}

func (nb *normalBuilder) makeHouseReady() {
	nb.houseReady = true
}

func (nb *normalBuilder) getHouse() (*normalHouse, error) {
	if !nb.houseReady {
		return nil, fmt.Errorf("normal house is not ready")
	}

	house := nb.house
	nb.reset()

	return house, nil
}

type iglooHouse struct {
	house
}

type iglooBuilder struct {
	houseReady bool
	house *iglooHouse
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{
		houseReady: false,
		house: &iglooHouse{},
	}
}

func (ib *iglooBuilder) reset() {
	ib.houseReady = false
	ib.house = &iglooHouse{}
}

func (ib *iglooBuilder) setWindowType() {
	ib.house.windowType = "Snow windows"
}

func (ib *iglooBuilder) setDoorType() {
	ib.house.doorType = "Snow doors"
}

func (ib *iglooBuilder) setNumFloor() {
	ib.house.floor = 4
}

func (ib *iglooBuilder) makeHouseReady() {
	ib.houseReady = true
}

func (ib *iglooBuilder) getHouse() (*iglooHouse, error) {
	if !ib.houseReady {
		return nil, fmt.Errorf("igloo house is not ready")
	}

	house := ib.house
	ib.reset()

	return house, nil
}

type director struct {}

func (d director) buildHouse(builder iBuilder) {
	builder.setDoorType()
	builder.setWindowType()
	builder.setNumFloor()
	builder.makeHouseReady()
}

func Builder_Pattern() {
	dtr := director{}

	nmrBuilder := newNormalBuilder()
	dtr.buildHouse(nmrBuilder)
	nmrHouse, _ := nmrBuilder.getHouse()

	fmt.Printf("House is ready = %v\n", nmrHouse)

	iglBuilder := newIglooBuilder()
	dtr.buildHouse(iglBuilder)
	iglHouse, _ := iglBuilder.getHouse()

	fmt.Printf("House is ready = %v\n", iglHouse)
}