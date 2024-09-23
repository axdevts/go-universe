package main

import (
	"fmt"

	fetches "github.com/axdevts/go-universe/fetchs"
)

func main() {
	fmt.Println("Hello from main!")
	fetches.SayHello()

	fetches.LoadData()

	ch := make(chan string)

	// Asynchronously fetch data
	go fetches.FetchData(ch)

	// Receive data from the channel
	result := <-ch

	fmt.Println("Result of fetched data >>>> ")
	fmt.Println(result)

}
