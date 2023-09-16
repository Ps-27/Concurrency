package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {

	//create two channel
	ping := make(chan string)
	pong := make(chan string)

	//running in background bcause it's infinite loop
	go shout(ping, pong)

	// time.Sleep(10 * time.Second)

	fmt.Println("Type of something and press Enter(enter Q to quit)")

	for {
		//print a prompt
		fmt.Print("->")

		//get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput
		//wait for a response

		response := <-pong
		fmt.Println("Response:", response)

	}

	fmt.Println("All done. closing channels.")
	close(ping)
	close(pong)

}
// output:
// go run channel.go
// Type of something and press Enter(enter Q to quit)
// ->prity
// Response: PRITY!!!
// ->Pgtyh
// Response: PGTYH!!!
// ->q
// All done. closing channels.
