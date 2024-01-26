package creational_patterns

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			fmt.Println("creating single instance")
			singleInstance = &single{}
		} else {
			fmt.Println("single instance already crated-1")
		}
	} else {
		fmt.Println("single instance already crated-2")
	}

	return singleInstance
}

func Singleton1() {
	for i:=0; i<30; i++ {
		go getInstance()
	}

	fmt.Scanln()
}