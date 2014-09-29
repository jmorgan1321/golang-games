package main

import (
	"fmt"
	"strings"
)

// Adding functions to a slice of string
type SmartSlice []string

func (ss SmartSlice) Filter(prefix string) SmartSlice {
	out := []string{}
	for _, s := range ss {
		if strings.HasPrefix(s, prefix) {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	s := SmartSlice{"fee", "fie", "foe", "fum"}
	fmt.Println(s.Filter("fi"))
}
