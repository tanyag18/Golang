package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(1 * time.Second)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {                      // just like select case , in this two using less times and is recieved first but not terminated 
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
			fmt.Println("received", msg2)
			// we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.
        }
    }
}