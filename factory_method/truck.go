package main

type Truck struct{
	Transport
}

func newTruck() ITransport{
	return &Truck{
		Transport: Transport{
			name:"Truck",
			cost: 5,
		},
	}
}