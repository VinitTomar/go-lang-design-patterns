package behavioral_patterns

import "fmt"

type observer interface {
	update(string)
	getID() string
}

type customer struct {
	id string
}

func (c *customer) update(event string) {
	fmt.Printf("Updating customer %v about the event %v\n", c.getID(), event)
}

func (c *customer) getID() string {
	return c.id
}

type item struct {
	observerList []observer
	name string
	inStock bool
}

func (i *item) updateAvailability() {
	fmt.Printf("Item %v is in stock now.\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o observer) {
	i.observerList = removeFromList(i.observerList, o)
}

func (i *item) notifyAll() {
	for _, obr := range i.observerList {
		obr.update(i.name)
	}
}

func newItem(name string) item {
	return item{
		name: name,
	}
}

func removeFromList(observerList []observer, o observer) []observer {
	toBeRemovedId := o.getID()

	for i, currentObserver := range observerList {
		if currentObserver.getID() == toBeRemovedId {
			observerList[0], observerList[i] = observerList[i], observerList[0]
			return observerList[1:]
		}
	}

	return observerList
}

func ObserverPattern() {
	item := newItem("Nike Shirt")

	customer1 := customer{
		id: "One",
	}

	customer2 := customer{
		id: "Two",
	}

	customer3 := customer{
		id: "Three",
	}

	item.register(&customer1)
	item.register(&customer2)
	item.register(&customer3)

	item.updateAvailability()

	fmt.Printf("Removing customer %v from list\n", customer2.getID())

	item.deregister(&customer2)

	item.updateAvailability()
}