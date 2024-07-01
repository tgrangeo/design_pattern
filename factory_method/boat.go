package main

type Boat struct{
	Transport
	year int
}

func newBoat() ITransport{
	return &Boat{
		Transport: Transport{
			name:"Boat",
			cost: 10,
		},
	}
}