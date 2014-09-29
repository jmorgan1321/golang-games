package main

import (
	"fmt"
	"time"
)

//beg OMIT
func main() {
	done := make(chan bool)

	go func() {
		<-time.After(5 * time.Second)
		done <- true
	}()

	tc := time.Tick(1 * time.Second)
	for {
		select {
		case <-tc:
			fmt.Println("tick")
		case <-done:
			fmt.Println("done")
			return
		}
	}
}

//end OMIT
