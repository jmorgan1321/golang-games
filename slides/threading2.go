package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("hello")
	defer fmt.Println("goodbye")

	//beg OMIT
	const actors int = 5
	r := rand.New(rand.NewSource(42))
	doneCh := make(chan bool)
	reportCh := make(chan bool)

	for i := 0; i < actors; i++ {
		go func(id int) {
			fmt.Println("actor", id, "reporting for duty!")
			reportCh <- true
			<-time.After(1 * time.Second)

			ch := time.After(time.Duration(r.Intn(5)+5) * time.Second)
			for {
				select {
				case <-time.Tick(time.Duration(r.Intn(500)+10) * time.Millisecond):
					fmt.Println("actor", id, "update")
				case <-ch:
					fmt.Println("actor", id, "finished")
					doneCh <- true
					return
				}
			}
		}(i)
		<-time.After(1 * time.Second)
	}

	for i := 0; i < actors; i++ {
		<-reportCh
	}
	fmt.Println("all reported!")

	for i := 0; i < actors; i++ {
		<-doneCh
	}

	fmt.Println("all finished!")
	//end OMIT
}
