package main

type ITransport interface {
	setName(name string )
	setCost(cost int)
	getName() string
	getCost(kilometer int) int
}

