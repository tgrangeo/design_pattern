package main

import "fmt"

func main() {
    truck, _ := getTransport("truck")
    boat, _ := getTransport("boat")

    printDetails(truck)
    printDetails(boat)
}

func printDetails(g ITransport) {
    fmt.Printf("Name: %s", g.getName())
    fmt.Println()
    fmt.Printf("Cost for 100 kilometer: %d", g.getCost(100))
    fmt.Println()
}
