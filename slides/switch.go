package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _ = range time.Tick(500 * time.Millisecond) {
		switch x := r.Intn(3); x { // HLs23
		case 0, 1: // HLs23
			y := 5 // HLs23
			fmt.Println(y + x)
		default:
			// there's no automatic fall through  // HLs23
			fmt.Println("default")
		}
	}
}
