package main

import (
	"fmt"
	"time"
)

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for {
			channel1 <- "Hey channel 1 working "
			time.Sleep(2 * time.Second)
		}

	}()
	go func() {
		for {
			channel2 <- "Hey channel 2 working "
			time.Sleep(3 * time.Second)
		}

	}()

	for {

		fmt.Println(<-channel1)
		fmt.Println(<-channel2)

	}
}
