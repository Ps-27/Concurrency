package main

import (
	"fmt"
	"time"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		fmt.Printf("%s goes to  the waiting room to check for clients.\n", barber)

		for {
			if len(shop.ClientsChan) == 0 {
				fmt.Printf("There is nothing to do,so %s takes a nap.\n", barber)
				isSleeping = true
			}
			client, shopOpen := <-shop.ClientsChan

			if shopOpen {
				if isSleeping {
					fmt.Printf("%s wakes %s up.\n", client, barber)
				}
				//cut hair
				shop.cutHair(barber, client)
			} else {
				shop.SendBarberHome(barber)
				return
			}

		}

	}()
}
func (shop *BarberShop) cutHair(barber, client string) {
	fmt.Printf("%s is cutting %s's hair.\n", barber, client)
	time.Sleep(shop.HairCutDuration)
	fmt.Printf("%s is finished cutting %s's hair.\n", barber, client)
}
func (shop *BarberShop) SendBarberHome(barber string) {
	fmt.Printf("%s is going home.\n", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	fmt.Printf("Closing shop for the day!\n")
	close(shop.ClientsChan)
	shop.Open = false

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}
	close(shop.BarbersDoneChan)
	fmt.Println("------------------------------------------------")
	fmt.Println("The barbershop is now closed for the day!, and everyone has gone home.\n")

}

func (shop *BarberShop) addClient(client string) {
	fmt.Printf("**** %s arrives!", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			fmt.Printf("%s takes a seat in the waiting room.\n", client)

		default:
			fmt.Printf("The waiting room is full, so %s leaves.\n", client)

		}
	} else {
		fmt.Printf("The shop is already closed, so %s leaves!\n", client)
	}
}
