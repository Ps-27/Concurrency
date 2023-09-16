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


// output:
//          go run SleepingBarber\main.go SleepingBarber\barbershop.go
// The Sleeping barber problem
// --------------------------
// The shop is open for the day!
// Alex goes to  the waiting room to check for clients.
// jender goes to  the waiting room to check for clients.
// Kelly goes to  the waiting room to check 
// for clients.
// There is nothing to do,so Kelly takes a nap.
// Milton  goes to  the waiting room to check for clients.
// There is nothing to do,so Alex takes a nap.
// There is nothing to do,so jender takes a 
// nap.
// Sushma goes to  the waiting room to check for clients.
// There is nothing to do,so Sushma takes a 
// nap.
// There is nothing to do,so Milton  takes a nap.
// **** Client #1 arrives!Client #1 takes a 
// seat in the waiting room.
// Client #1 wakes Kelly up.
// Kelly is cutting Client #1's hair.       
// **** Client #2 arrives!Client #2 takes a 
// seat in the waiting room.
// Client #2 wakes Alex up.
// Alex is cutting Client #2's hair.        
// **** Client #3 arrives!Client #3 takes a 
// seat in the waiting room.
// Client #3 wakes jender up.
// jender is cutting Client #3's hair.      
// **** Client #4 arrives!Client #4 takes a 
// seat in the waiting room.
// Client #4 wakes Sushma up.
// Sushma is cutting Client #4's hair.      
// **** Client #5 arrives!Client #5 takes a 
// seat in the waiting room.
// Client #5 wakes Milton  up.
// Milton  is cutting Client #5's hair.     
// **** Client #6 arrives!The waiting room is full, so Client #6 leaves.
// **** Client #7 arrives!The waiting room is full, so Client #7 leaves.
// **** Client #8 arrives!The waiting room is full, so Client #8 leaves.
// **** Client #9 arrives!The waiting room is full, so Client #9 leaves.
// **** Client #10 arrives!The waiting room 
// is full, so Client #10 leaves.
// Kelly is finished cutting Client #1's hair.
// There is nothing to do,so Kelly takes a nap.
// **** Client #11 arrives!Client #11 takes 
// a seat in the waiting room.
// Client #11 wakes Kelly up.
// Kelly is cutting Client #11's hair.      
// Alex is finished cutting Client #2's hair.
// There is nothing to do,so Alex takes a nap.
// **** Client #12 arrives!Client #12 takes 
// a seat in the waiting room.
// Client #12 wakes Alex up.
// Alex is cutting Client #12's hair.       
// jender is finished cutting Client #3's hair.
// There is nothing to do,so jender takes a 
// nap.
// **** Client #13 arrives!Client #13 takes 
// a seat in the waiting room.
// Client #13 wakes jender up.
// jender is cutting Client #13's hair.     
// Sushma is finished cutting Client #4's hair.
// There is nothing to do,so Sushma takes a 
// nap.
// **** Client #14 arrives!Client #14 takes 
// a seat in the waiting room.
// Client #14 wakes Sushma up.
// Sushma is cutting Client #14's hair.     
// Milton  is finished cutting Client #5's hair.
// There is nothing to do,so Milton  takes a nap.
// **** Client #15 arrives!Client #15 takes 
// a seat in the waiting room.
// Client #15 wakes Milton  up.
// Milton  is cutting Client #15's hair.    
// **** Client #16 arrives!The waiting room 
// is full, so Client #16 leaves.
// Kelly is finished cutting Client #11's hair.
// There is nothing to do,so Kelly takes a nap.
// **** Client #17 arrives!Client #17 takes 
// a seat in the waiting room.
// Client #17 wakes Kelly up.
// Kelly is cutting Client #17's hair.      
// Alex is finished cutting Client #12's hair.
// There is nothing to do,so Alex takes a nap.
// **** Client #18 arrives!Client #18 takes 
// a seat in the waiting room.
// Client #18 wakes Alex up.
// Alex is cutting Client #18's hair.       
// **** Client #19 arrives!The waiting room 
// is full, so Client #19 leaves.
// **** Client #20 arrives!The waiting room 
// is full, so Client #20 leaves.
// **** Client #21 arrives!The waiting room 
// is full, so Client #21 leaves.
// jender is finished cutting Client #13's hair.
// There is nothing to do,so jender takes a 
// nap.
// **** Client #22 arrives!Client #22 takes 
// a seat in the waiting room.
// Client #22 wakes jender up.
// jender is cutting Client #22's hair.     
// Sushma is finished cutting Client #14's hair.
// There is nothing to do,so Sushma takes a 
// nap.
// **** Client #23 arrives!Client #23 takes 
// a seat in the waiting room.
// Client #23 wakes Sushma up.
// Sushma is cutting Client #23's hair.     
// **** Client #24 arrives!The waiting room 
// is full, so Client #24 leaves.
// Milton  is finished cutting Client #15's 
// hair.
// There is nothing to do,so Milton  takes a nap.
// **** Client #25 arrives!Client #25 takes 
// a seat in the waiting room.
// Client #25 wakes Milton  up.
// Milton  is cutting Client #25's hair.    
// **** Client #26 arrives!The waiting room 
// is full, so Client #26 leaves.
// **** Client #27 arrives!The waiting room 
// is full, so Client #27 leaves.
// **** Client #28 arrives!The waiting room 
// is full, so Client #28 leaves.
// **** Client #29 arrives!The waiting room 
// is full, so Client #29 leaves.
// Kelly is finished cutting Client #17's hair.
// There is nothing to do,so Kelly takes a nap.
// **** Client #30 arrives!Client #30 takes 
// a seat in the waiting room.
// Alex is finished cutting Client #18's hair.
// Client #30 wakes Kelly up.
// Kelly is cutting Client #30's hair.      
// There is nothing to do,so Alex takes a nap.
// jender is finished cutting Client #22's hair.
// There is nothing to do,so jender takes a 
// nap.
// **** Client #31 arrives!Client #31 takes 
// a seat in the waiting room.
// Client #31 wakes Alex up.
// Alex is cutting Client #31's hair.       
// Sushma is finished cutting Client #23's hair.
// There is nothing to do,so Sushma takes a 
// nap.
// **** Client #32 arrives!Client #32 takes 
// a seat in the waiting room.
// Client #32 wakes jender up.
// jender is cutting Client #32's hair.     
// **** Client #33 arrives!Client #33 takes 
// a seat in the waiting room.
// Client #33 wakes Sushma up.
// Sushma is cutting Client #33's hair.     
// Milton  is finished cutting Client #25's 
// hair.
// There is nothing to do,so Milton  takes a nap.
// **** Client #34 arrives!Client #34 takes 
// a seat in the waiting room.
// Client #34 wakes Milton  up.
// Milton  is cutting Client #34's hair.    
// **** Client #35 arrives!The waiting room 
// is full, so Client #35 leaves.
// **** Client #36 arrives!The waiting room 
// is full, so Client #36 leaves.
// **** Client #37 arrives!The waiting room 
// is full, so Client #37 leaves.
// **** Client #38 arrives!The waiting room 
// is full, so Client #38 leaves.
// Kelly is finished cutting Client #30's hair.
// There is nothing to do,so Kelly takes a nap.
// **** Client #39 arrives!Client #39 wakes 
// Kelly up.
// Kelly is cutting Client #39's hair.      
// Client #39 takes a seat in the waiting room.
// Alex is finished cutting Client #31's hair.
// There is nothing to do,so Alex takes a nap.
// **** Client #40 arrives!Client #40 takes 
// a seat in the waiting room.
// Client #40 wakes Alex up.
// Alex is cutting Client #40's hair.       
// **** Client #41 arrives!The waiting room 
// is full, so Client #41 leaves.
// jender is finished cutting Client #32's hair.
// There is nothing to do,so jender takes a 
// nap.
// **** Client #42 arrives!Client #42 takes 
// a seat in the waiting room.
// Client #42 wakes jender up.
// jender is cutting Client #42's hair.     
// **** Client #43 arrives!The waiting room 
// is full, so Client #43 leaves.
// Sushma is finished cutting Client #33's hair.
// There is nothing to do,so Sushma takes a 
// nap.
// **** Client #44 arrives!Client #44 takes 
// a seat in the waiting room.
// Client #44 wakes Sushma up.
// Sushma is cutting Client #44's hair.     
// Milton  is finished cutting Client #34's 
// hair.
// There is nothing to do,so Milton  takes a nap.
// **** Client #45 arrives!Client #45 takes 
// a seat in the waiting room.
// Client #45 wakes Milton  up.
// Milton  is cutting Client #45's hair.    
// **** Client #46 arrives!The waiting room 
// is full, so Client #46 leaves.
// **** Client #47 arrives!The waiting room 
// is full, so Client #47 leaves.
// **** Client #48 arrives!The waiting room 
// is full, so Client #48 leaves.
// Kelly is finished cutting Client #39's hair.
// There is nothing to do,so Kelly takes a nap.
// **** Client #49 arrives!Client #49 takes 
// a seat in the waiting room.
// Client #49 wakes Kelly up.
// Kelly is cutting Client #49's hair.      
// **** Client #50 arrives!The waiting room 
// is full, so Client #50 leaves.
// Alex is finished cutting Client #40's hair.
// There is nothing to do,so Alex takes a nap.
// **** Client #51 arrives!Client #51 takes 
// a seat in the waiting room.
// Client #51 wakes Alex up.
// Alex is cutting Client #51's hair.       
// jender is finished cutting Client #42's hair.
// There is nothing to do,so jender takes a 
// nap.
// **** Client #52 arrives!Client #52 takes 
// a seat in the waiting room.
// Client #52 wakes jender up.
// jender is cutting Client #52's hair.     
// **** Client #53 arrives!The waiting room 
// is full, so Client #53 leaves.
// Sushma is finished cutting Client #44's hair.
// There is nothing to do,so Sushma takes a 
// nap.
// Milton  is finished cutting Client #45's 
// hair.
// There is nothing to do,so Milton  takes a nap.
// **** Client #54 arrives!Client #54 takes 
// a seat in the waiting room.
// Client #54 wakes Sushma up.
// Sushma is cutting Client #54's hair.     
// **** Client #55 arrives!Client #55 takes 
// a seat in the waiting room.
// Client #55 wakes Milton  up.
// Milton  is cutting Client #55's hair.    
// **** Client #56 arrives!The waiting room 
// is full, so Client #56 leaves.
// **** Client #57 arrives!The waiting room 
// is full, so Client #57 leaves.
// **** Client #58 arrives!The waiting room 
// is full, so Client #58 leaves.
// Kelly is finished cutting Client #49's hair.
// There is nothing to do,so Kelly takes a nap.
// **** Client #59 arrives!Client #59 takes 
// a seat in the waiting room.
// Client #59 wakes Kelly up.
// Kelly is cutting Client #59's hair.      
// Alex is finished cutting Client #51's hair.
// There is nothing to do,so Alex takes a nap.
// **** Client #60 arrives!Client #60 takes 
// a seat in the waiting room.
// Client #60 wakes Alex up.
// Alex is cutting Client #60's hair.       
// jender is finished cutting Client #52's hair.
// There is nothing to do,so jender takes a 
// nap.
// **** Client #61 arrives!Client #61 takes 
// a seat in the waiting room.
// Client #61 wakes jender up.
// jender is cutting Client #61's hair.     
// **** Client #62 arrives!The waiting room 
// is full, so Client #62 leaves.
// Sushma is finished cutting Client #54's hair.
// There is nothing to do,so Sushma takes a 
// nap.
// **** Client #63 arrives!Client #63 takes 
// a seat in the waiting room.
// Client #63 wakes Sushma up.
// Sushma is cutting Client #63's hair.     
// **** Client #64 arrives!The waiting room 
// is full, so Client #64 leaves.
// Milton  is finished cutting Client #55's 
// hair.
// There is nothing to do,so Milton  takes a nap.
// **** Client #65 arrives!Client #65 takes 
// a seat in the waiting room.
// Client #65 wakes Milton  up.
// Milton  is cutting Client #65's hair.    
// **** Client #66 arrives!The waiting room 
// is full, so Client #66 leaves.
// **** Client #67 arrives!The waiting room 
// is full, so Client #67 leaves.
// **** Client #68 arrives!The waiting room 
// is full, so Client #68 leaves.
// Kelly is finished cutting Client #59's hair.
// There is nothing to do,so Kelly takes a nap.
// **** Client #69 arrives!Client #69 takes 
// a seat in the waiting room.
// Client #69 wakes Kelly up.
// Kelly is cutting Client #69's hair.      
// Alex is finished cutting Client #60's hair.
// There is nothing to do,so Alex takes a nap.
// **** Client #70 arrives!Client #70 takes 
// a seat in the waiting room.
// Client #70 wakes Alex up.
// Alex is cutting Client #70's hair.       
// **** Client #71 arrives!The waiting room 
// is full, so Client #71 leaves.
// jender is finished cutting Client #61's hair.
// There is nothing to do,so jender takes a 
// nap.
// **** Client #72 arrives!Client #72 takes 
// a seat in the waiting room.
// Client #72 wakes jender up.
// jender is cutting Client #72's hair.     
// **** Client #73 arrives!The waiting room 
// is full, so Client #73 leaves.
// Sushma is finished cutting Client #63's hair.
// There is nothing to do,so Sushma takes a 
// nap.
// **** Client #74 arrives!Client #74 takes 
// a seat in the waiting room.
// Client #74 wakes Sushma up.
// Sushma is cutting Client #74's hair.     
// **** Client #75 arrives!The waiting room 
// is full, so Client #75 leaves.
// Milton  is finished cutting Client #65's 
// hair.
// There is nothing to do,so Milton  takes a nap.
// **** Client #76 arrives!Client #76 takes 
// a seat in the waiting room.
// Client #76 wakes Milton  up.
// Milton  is cutting Client #76's hair.    
// **** Client #77 arrives!The waiting room 
// is full, so Client #77 leaves.
// Kelly is finished cutting Client #69's hair.
// There is nothing to do,so Kelly takes a nap.
// Alex is finished cutting Client #70's hair.
// There is nothing to do,so Alex takes a nap.
// **** Client #78 arrives!Client #78 takes 
// a seat in the waiting room.
// Client #78 wakes Kelly up.
// Kelly is cutting Client #78's hair.      
// **** Client #79 arrives!Client #79 takes 
// a seat in the waiting room.
// Client #79 wakes Alex up.
// Alex is cutting Client #79's hair.       
// jender is finished cutting Client #72's hair.
// There is nothing to do,so jender takes a 
// nap.
// **** Client #80 arrives!Client #80 takes 
// a seat in the waiting room.
// Client #80 wakes jender up.
// jender is cutting Client #80's hair.     
// **** Client #81 arrives!The waiting room 
// is full, so Client #81 leaves.
// Sushma is finished cutting Client #74's hair.
// There is nothing to do,so Sushma takes a 
// nap.
// **** Client #82 arrives!Client #82 takes 
// a seat in the waiting room.
// Client #82 wakes Sushma up.
// Sushma is cutting Client #82's hair.     
// **** Client #83 arrives!The waiting room 
// is full, so Client #83 leaves.
// Milton  is finished cutting Client #76's 
// hair.
// There is nothing to do,so Milton  takes a nap.
// **** Client #84 arrives!Client #84 takes 
// a seat in the waiting room.
// Client #84 wakes Milton  up.
// Milton  is cutting Client #84's hair.    
// **** Client #85 arrives!The waiting room 
// is full, so Client #85 leaves.
// **** Client #86 arrives!The waiting room 
// is full, so Client #86 leaves.
// **** Client #87 arrives!The waiting room 
// is full, so Client #87 leaves.
// Kelly is finished cutting Client #78's hair.
// There is nothing to do,so Kelly takes a nap.
// Closing shop for the day!
// Kelly is going home.
// Alex is finished cutting Client #79's hair.
// There is nothing to do,so Alex takes a nap.
// Alex is going home.
// jender is finished cutting Client #80's hair.
// There is nothing to do,so jender takes a 
// nap.
// jender is going home.
// Sushma is finished cutting Client #82's hair.
// There is nothing to do,so Sushma takes a 
// nap.
// Sushma is going home.
// Milton  is finished cutting Client #84's hair.
// There is nothing to do,so Milton  takes a nap.
// Milton  is going home.
// ------------------------------------------------
// The barbershop is now closed for the day!, and everyone has gone home.


