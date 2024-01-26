package creational_patterns

import (
	"fmt"
	"sync"
)

var once sync.Once

type single2 struct {}

var singleInstance2 *single2

func getInstance2() *single2 {
	if singleInstance2 == nil {
		once.Do(func () {
			fmt.Println("Creating single instance")
			singleInstance2 = &single2{}
		})
	} else {
		fmt.Println("Single instance already crated")
	}

	return singleInstance2
}

func Singleton2() {

	for i:=0; i<10; i++ {
		go getInstance2()
	}

	fmt.Scanln()

}