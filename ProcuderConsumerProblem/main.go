package main

import (
	// "color"
	"fmt"
	"math/rand"
	"time"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}
type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order %d !\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza %d . It will take %d second ...\n", pizzaNumber, delay)

		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** we ran  out of ingrediant for pizza #%d", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}
func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 1
	// run forever or until we receive a quit notification

	//try to make pizzas
	for {
		currentPizza := makePizza(i)
		fmt.Println(currentPizza)
		//try to make pizzas
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried to make a pizza (we sent something to the data channel)
			case pizzaMaker.data <- **&currentPizza:

			case quitChan := <-pizzaMaker.quit:
				//close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
		//decision
	}
}

func main() {

	//seed the random nu,ber generator
	rand.Seed(time.Now().UnixNano())

	//print out a message
	// color.Cyan("The pizzeria is open for bussiness!")
	// color.Cyan("-------------------------")
	fmt.Println("The pizzeria is open for bussiness!")
	fmt.Println("-------------------------")

	//Create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//run the producer in the background
	go pizzeria(pizzaJob)

	//create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				fmt.Println(i.message)
				fmt.Printf("Order %d is out for delivery!", i.pizzaNumber)
			} else {
				fmt.Println(i.message)
				fmt.Println("The customer is really mad !")
			}
		} else {
			fmt.Println("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	//print out the ending

	fmt.Println("_________________")
	fmt.Println("Done for the day.")
	fmt.Printf("We made %d pizzas, but failed to make %d,with %d attempts in total.", pizzasMade, pizzasFailed, total)

}
//output:
// go run ProducerConsumer\main.go        
// The pizzeria is open for bussiness!
// -------------------------
// Received order 2 !
// Making pizza 2 . It will take 4 second ...     
// &{2 pizza order #2 is ready! true}
// Received order 3 !
// Making pizza 3 . It will take 2 second ...     
// pizza order #2 is ready!
// Order 2 is out for delivery!&{3 pizza order #3 
// is ready! true}
// pizza order #3 is ready!
// Order 3 is out for delivery!Received order 4 ! 
// Making pizza 4 . It will take 4 second ...     
// &{4 pizza order #4 is ready! true}
// Received order 5 !
// Making pizza 5 . It will take 3 second ...     
// pizza order #4 is ready!
// Order 4 is out for delivery!&{5 pizza order #5 
// is ready! true}
// Received order 6 !
// Making pizza 6 . It will take 3 second ...     
// pizza order #5 is ready!
// Order 5 is out for delivery!&{6 pizza order #6 
// is ready! true}
// pizza order #6 is ready!
// Order 6 is out for delivery!Received order 7 ! 
// Making pizza 7 . It will take 1 second ...     
// &{7 pizza order #7 is ready! true}
// Received order 8 !
// Making pizza 8 . It will take 3 second ...     
// pizza order #7 is ready!
// Order 7 is out for delivery!&{8 *** we ran  out of ingrediant for pizza #8 false}
// *** we ran  out of ingrediant for pizza #8
// The customer is really mad !
// Received order 9 !
// Making pizza 9 . It will take 4 second ...     
// &{9 pizza order #9 is ready! true}
// Received order 10 !
// Making pizza 10 . It will take 3 second ...    
// pizza order #9 is ready!
// Order 9 is out for delivery!&{10 pizza order #10 is ready! true}
// &{11  false}
// pizza order #10 is ready!
// Order 10 is out for delivery!Done making pizzas...
// &{12  false}
// _________________
// Done for the day.
// We made 8 pizzas, but failed to make 1,with 9 attempts in total.
