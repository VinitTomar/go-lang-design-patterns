package behavioral_patterns

import (
	"fmt"
	"log"
)

type state interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

type noItemState struct {
	v *vendingMachine
}

func (s *noItemState) addItem(count int) error {
	s.v.incrementItemCount(count)
	s.v.setState(s.v.hasItem)
	return nil
}

func (s *noItemState) requestItem() error {
	return fmt.Errorf("out of stock")
}

func (s *noItemState) insertMoney(money int) error {
	return fmt.Errorf("out of stock")
}

func (s *noItemState) dispenseItem() error {
	return fmt.Errorf("out of stock")
}

type hasItemState struct {
	v *vendingMachine
}

func (s *hasItemState) addItem(count int) error {
	s.v.incrementItemCount(count)
	s.v.setState(s.v.hasItem)
	return nil
}

func (s *hasItemState) requestItem() error {
	if s.v.itemCount == 0 {
		s.v.setState(s.v.noItem)
		return fmt.Errorf("out of stock")
	}

	fmt.Println("Item requested")
	s.v.setState(s.v.itemRequested)
	return nil
}

func (s *hasItemState) insertMoney(money int) error {
	return fmt.Errorf("please request an item first")
}

func (s *hasItemState) dispenseItem() error {
	return fmt.Errorf("please select an Item")
}

type itemRequestedState struct {
	v *vendingMachine
}

func (s *itemRequestedState) addItem(int) error {
	return fmt.Errorf("item dispense in progress")
}

func (s *itemRequestedState) requestItem() error {
	return fmt.Errorf("item already requested")
}

func (s *itemRequestedState) insertMoney(money int) error {
	if money < s.v.itemPrice {
		return fmt.Errorf("insert money %v is less than item price %v", money, s.v.itemPrice)
	}

	fmt.Println("Money entered is OK")
	s.v.setState(s.v.hasMoney)
	return nil
}

func (s *itemRequestedState) dispenseItem() error {
	return fmt.Errorf("please insert money first")
}

type hasMoneyState struct {
	v *vendingMachine
}

func (s *hasMoneyState) addItem(int) error {
	return fmt.Errorf("item dispense in progress")
}

func (s *hasMoneyState) requestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (s *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("item dispense in progress")
}

func (s *hasMoneyState) dispenseItem() error {
	fmt.Println("dispensing item")

	s.v.itemCount -= 1

	if s.v.itemCount == 0 {
		s.v.setState(s.v.noItem)
	} else {
		s.v.setState(s.v.hasItem)
	}

	return nil
}

type vendingMachine struct {
	hasItem state
	itemRequested state
	noItem state
	hasMoney state

	currentState state
	
	itemCount int
	itemPrice int
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) addItem(n int) error {
	return v.currentState.addItem(n)
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) incrementItemCount(n int) {
	fmt.Printf("Adding %v more items.\n", n)
	v.itemCount += n
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func newVendingMachine(itmCount, itmPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itmCount,
		itemPrice: itmPrice,
	}

	hsItmState := &hasItemState{v: v}
	noItmState := &noItemState{v: v}
	itmRqtState := &itemRequestedState{v: v}
	hsMnyState := hasMoneyState{v: v}

	v.setState(hsItmState)
	v.hasItem = hsItmState
	v.noItem = noItmState
	v.itemRequested = itmRqtState
	v.hasMoney = &hsMnyState

	return v
}

func StatePattern() {
	vMachine := newVendingMachine(1, 1)

	 err := vMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

}