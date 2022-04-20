package main

import (
	"fmt"
	"time"
)

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go func() {
		for {
			time.Sleep(3 * time.Second)
			channel1 <- "Hey, channel 1 working - after  3 secs "
		}

	}()
	go func() {
		for {
			time.Sleep(7 * time.Second)
			channel2 <- "Hey, channel 2 working - after  7 secs "
		}
	}()

	for {
		select {
		case <-channel1:
			fmt.Println(<-channel1)
		case <-channel2:
			fmt.Println(<-channel2)
		case <-time.After(14 * time.Second):
			fmt.Println("Bye")
			return
		default:
			fmt.Println("Waiting...")
			time.Sleep(1 * time.Second)
		}
	}
}
