package main

import (
	"fmt"
	"sync"
)

func PritText(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	fmt.Println("First Example")
	var wg sync.WaitGroup
	words := []string{
		"alpha",
		"beta",
		"gamma",
	}

	wg.Add(3)

	for i, v := range words {
		go PritText(fmt.Sprintf("%d: %s", i, v), &wg)
	}
	wg.Wait()
	fmt.Println("END!")

}
// output: 
// go run SimpleWaitgrpExample\main.go    
// First Example
// 2: gamma
// 0: alpha
// 1: beta
// END!


// no order is following
