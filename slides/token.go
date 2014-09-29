package main

import (
	"fmt"
	"github.com/jmorgan1321/golang-games/slides/tok"
)

func main() {
	if tk := tok.For; tok.IsKeyword(tk) {
		fmt.Printf("'%s' is a keyword.\n", tk)
	} else {
		fmt.Printf("'%s' is NOT a keyword.\n", tk)
	}

	if tk := tok.Mul; tok.IsKeyword(tk) {
		fmt.Printf("'%s' is a keyword.\n", tk)
	} else {
		fmt.Printf("'%s' is NOT a keyword.\n", tk)
	}
}
