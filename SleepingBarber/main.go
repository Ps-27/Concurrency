package main

import (
	"fmt"
	"math/rand"
	"time"
)

// variables
var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	//seed our random number generator
	rand.Seed(time.Now().UnixNano())

	//print  welcome msg
	fmt.Println("The Sleeping barber problem")
	fmt.Println("--------------------------")

	// create channels if we need any
	clientChan := make(chan string)
	doneChan := make(chan bool)

	//create the barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}
	fmt.Println("The shop is open for the day!")

	//add barbers
	shop.addBarber("Alex")
	shop.addBarber("jender")
	shop.addBarber("Milton ")
	shop.addBarber("Kelly")
	shop.addBarber("Sushma")
	//start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	//add clients
	i := 1

	go func(i int) {
		for {
			//get a random number with avg arrival rate
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}(i)

	//block until the barbershop is close
	<-closed
	// time.Sleep(1 * time.Second)
}

