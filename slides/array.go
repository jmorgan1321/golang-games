package main

func Foo(arr [3]int) {}

func main() {
	a3, a4 := [3]int{}, [4]int{}

	a3 = a4
	a3[3] = 1
	Foo(a4)

	Foo(a3)
}
