package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updatemsg(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}
func main() {
	msg = "hello! "

	var mutex sync.Mutex

	wg.Add(2)
	go updatemsg("Hello Prity", &mutex)
	go updatemsg("ps", &mutex)
	wg.Wait()

	fmt.Println(msg)

}


//output:
//  go run syncMutex\usingmutex.go
// updating to  ps
// updating to  Hello Prity
// Hello Prity
