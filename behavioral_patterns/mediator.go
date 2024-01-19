package behavioral_patterns

import (
	"fmt"
	"sync"
)

type Mediator interface {
	canArrive(Train) bool
	notifyDeparture()
}

type Train interface {
	requestArrival()
	departure()
	permitArrival()
}

type PassengerTrain struct {
	mediator Mediator
}

func (t *PassengerTrain) requestArrival() {
	if t.mediator.canArrive(t) {
		fmt.Println("Passenger train is arriving on platform")
	} else {
		fmt.Println("Passenger train is waiting for the arrival permission")
	}
}

func (t *PassengerTrain) departure() {
	fmt.Println("Passenger train is leaving platform")
	t.mediator.notifyDeparture()
}

func (t *PassengerTrain) permitArrival() {
	fmt.Println("Passenger train permitted for arrival")
	t.requestArrival()
}

type GoodsTrain struct {
	mediator Mediator
}

func (t *GoodsTrain) requestArrival() {
	if t.mediator.canArrive(t) {
		fmt.Println("Goods train is arriving on platform")
	} else {
		fmt.Println("Goods train is waiting for the arrival permission")
	}
}

func (t *GoodsTrain) departure() {
	fmt.Println("Goods train is leaving platform")
	t.mediator.notifyDeparture()
}

func (t *GoodsTrain) permitArrival() {
	fmt.Println("Goods train permitted for arrival")
	t.requestArrival()
}

type StationMaster struct {
	isPlatformFree bool
	lock *sync.Mutex
	waitingTrains []Train
}

func (sm *StationMaster) canArrive(t Train) bool {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	if sm.isPlatformFree {
		sm.isPlatformFree = false
		return true
	}

	sm.waitingTrains = append(sm.waitingTrains, t)
	return false
}

func (sm *StationMaster) notifyDeparture() {
	if len(sm.waitingTrains) > 0 {
		firstWaitingTrain := sm.waitingTrains[0]
		sm.waitingTrains = sm.waitingTrains[1:]
		defer firstWaitingTrain.permitArrival()
	}
	
	sm.lock.Lock()
	defer sm.lock.Unlock()
	if !sm.isPlatformFree {
		sm.isPlatformFree = true
	}

}

func newStationMaster() *StationMaster {
	return &StationMaster{
		isPlatformFree: true,
		lock: &sync.Mutex{},
	}
}

func MediatorDesignPattern() {
	stationManager := newStationMaster()

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}
	
	goodsTrain := &GoodsTrain{
		mediator: stationManager,
	}
	
	passengerTrain.requestArrival()
	goodsTrain.requestArrival()
	passengerTrain.departure()
}