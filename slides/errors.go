package main

func funcThatReturnsAnError() (int, error) { return 0, nil }
func funcThatDoesntReturnAnError()         { return }
func main() {
	funcThatReturnsAnError()      // you can still ignore ALL
	x := funcThatReturnsAnError() // error, must explicitly ignore error
	x++
	funcThatDoesntReturnAnError()
	if _, err := funcThatReturnsAnError(); err != nil {
		//handle error
	}
}
