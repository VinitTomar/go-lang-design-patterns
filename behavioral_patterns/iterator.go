package behavioral_patterns

import (
	"fmt"
)

type iterator[T user] interface {
	next() *T
	hasNext() bool
}

type collection[T any] interface {
	createIterator() iterator[user]
}

type user struct {
	name string
	age int
}

type userIterator struct {
	pos int
	users []*user
}

/**
TODO remove below unused warning
*/
func (uItr *userIterator) next() *user {
	if uItr.pos < len(uItr.users) {
		usr := uItr.users[uItr.pos]
		uItr.pos += 1
		return usr
	}

	return nil
}

/**
TODO remove below unused warning
*/
func (uItr *userIterator) hasNext() bool {
	return uItr.pos < len(uItr.users)
}

type userCollection struct {
	users []*user
}

func (usrClc *userCollection) createIterator() iterator[user] {
	return &userIterator{
		pos: 0,
		users: usrClc.users,
	}
}


func IteratorPattern() {
	usr1 := &user{
		name: "John",
		age: 34,
	}

	usr2 := &user{
		name: "Ben",
		age: 21,
	}

	var usrCollection collection[user] = &userCollection{
		users: []*user{usr1, usr2},
	}

	var itr iterator[user] = usrCollection.createIterator()

	for itr.hasNext() {
		fmt.Printf("User is %v\n", itr.next())
	}
	
}
