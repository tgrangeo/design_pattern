package main

type Transport struct {
	name string
	cost int
}

func (t *Transport) setName(name string) {
    t.name = name
}

func (t *Transport) getName() string {
    return t.name
}

func (t *Transport) setCost(cost int) {
    t.cost = cost
}

func (t *Transport) getCost(kilometer int) int {
    return t.cost * kilometer
}