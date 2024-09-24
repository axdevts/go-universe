package main

import (
	"fmt"

	"github.com/axdevts/go-universe/fetches"
	"github.com/axdevts/go-universe/scrap"
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

	scrap.ScrapDataBySoup()

	scrap.ScrapByColly()

	scrap.ScrapByRod()

}
