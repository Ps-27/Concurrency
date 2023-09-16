package main

import (
	"fmt"
	"sync"
	"time"
)

//The dining Philosophers problem
//five Philosophers
//eating at same time
//sheghetti which require 2 fork
//may be eating simultaneously , since there are five philosphers and five fork

type Philosophers struct {
	name      string
	rightFork int
	leftFork  int
}

// list of all philosophers
var philosophers = []Philosophers{
	{name: "Prity", leftFork: 4, rightFork: 0},
	{name: "Prity1", leftFork: 0, rightFork: 1},
	{name: "Prity2", leftFork: 1, rightFork: 2},
	{name: "Prity3", leftFork: 2, rightFork: 3},
	{name: "Prity4", leftFork: 3, rightFork: 4},
}

// define some variable
var hunger = 3 //how many times does aperson eat?
var eattime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

// added this

var orderMutex sync.Mutex
var orderFinished []string

func main() {
	// print out a welcome msg
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("--------------------------------")
	fmt.Println("The table is empty.")

	time.Sleep(sleepTime)

	// start the meal
	dine()

	// print out finished msg
	time.Sleep(sleepTime)
	fmt.Println("The table is empty.")

	fmt.Println("order finished : ", orderFinished)
}
func dine() {
	eattime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second
	wg := &sync.WaitGroup{}

	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	//forks is a map of all 5 forks.
	var forks = make(map[int]*sync.Mutex)

	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	//start the meal
	for i := 0; i < len(philosophers); i++ {
		//fire off a goroutine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}
	wg.Wait()

}
func diningProblem(philosopher Philosophers, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	//seat the philosopher at athe table

	fmt.Printf("%s is seated at the table.\n", philosopher.name)
	seated.Done()

	//eat three times

	for i := hunger; i > 0; i-- {
		//get a lock on both forks
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eattime)

		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", philosopher.name)

	}
	fmt.Println(philosopher.name, "is satisfied.")
	fmt.Println(philosopher.name, "left the table")

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()
}
