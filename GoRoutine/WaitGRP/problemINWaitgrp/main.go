package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMsg(s string) {
	defer wg.Done()
	msg = s
}
func main() {
	msg = "hello! "
	wg.Add(2)
	
	go updateMsg("Hello Prity")
	go updateMsg("ps")
	wg.Wait()

	fmt.Println(msg)
	fmt.Println()

}

//problem in WaitGRP
// output:
// go run syncMutex\main.go
// Hello Prity
