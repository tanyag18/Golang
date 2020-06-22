package main

import "fmt"


func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}


    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)


func main() {

    channel := make(chan string,2)            //create a channel //2 values buffered
	go func() { channel <- "ping1" }()
	go func() { channel <- "ping2" }()   //any order

	// channel <- "ping1"  // will print ping 1 then ping 2
	// channel <- "ping2"
	fmt.Println(<-channel) 
	fmt.Println(<-channel)           //  allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization
    // msg := <-messages
	// fmt.Println(msg)
	

	//channel direction
	pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)


}