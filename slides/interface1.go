package main

import "fmt"

//beg show A OMIT
type Fighter interface { // HLxxx
	Fight() // HLxxx
} // HLxxx

type Boxer int

func (b *Boxer) Fight() {
	for i := *b; i >= 0; i-- {
		fmt.Print("\tpunch")
	}
	fmt.Println("")
}

type MmaFighter struct{ Fighter } // HLxxx

type KravMagaMaster struct {
	styles []Fighter
}

func (km *KravMagaMaster) Fight() {
	for _, s := range km.styles {
		s.Fight()
	}
}

//end show A OMIT

type KickBoxer struct {
	BlackBeltDegree int
}

func (k *KickBoxer) Fight() {
	for i := k.BlackBeltDegree; i >= 0; i-- {
		fmt.Print("\tkick")
	}
	fmt.Println("")
}

type Judoka struct{}

func (*Judoka) Fight() {
	fmt.Println("\tthrow")
}

type Wrestler struct{}

func (*Wrestler) Fight() {
	fmt.Println("\tgrapple")
}

//beg show B OMIT

func main() {
	b3, b1 := Boxer(3), Boxer(1)

	fighters := []Fighter{
		&b1,
		&KickBoxer{BlackBeltDegree: 3},
		&Wrestler{},
		&MmaFighter{Fighter: &Judoka{}},
		&KravMagaMaster{
			styles: []Fighter{
				&KickBoxer{2},
				&Judoka{},
				&Wrestler{},
				&b3,
			},
		},
	}

	for _, f := range fighters {
		fmt.Println("=============================")
		f.Fight()
		fmt.Println("=============================")
	}
}

//end show B OMIT
