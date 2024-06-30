package main

import (
    "fmt"
)

func main() {
	ptr_instance := getInstance()
	ptr_instance.name = "alice"
    fmt.Println(ptr_instance.name)

	other_instance := getInstance()
	other_instance.name = "bob"
	
	fmt.Println(ptr_instance.name)
}