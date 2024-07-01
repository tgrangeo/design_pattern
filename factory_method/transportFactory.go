package main

import "fmt"


//To get a new transport use this function
func getTransport(transportType string)(ITransport,error){
	if transportType == "truck"{
		return newTruck(),nil
	}
	if transportType == "boat"{
		return newBoat(),nil
	}
	return nil, fmt.Errorf("wrong type")
}