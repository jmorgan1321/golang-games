package main

import (
	"bufio"
	"fmt"
	"github.com/jmorgan1321/golang-games/core/utils"
	"os"
	"strings"
)

func main() {
	printStartMsg()
	defer printExitMsg()

	processCmdLine()

	ec := make(chan int)
	// listen to stdin and exit/restart if asked to
	go func() {
		bio := bufio.NewReader(os.Stdin)
		for {
			line, _, _ := bio.ReadLine()
			s := string(line)
			if strings.HasPrefix(s, "exit") {
				printExitMsg()
				ec <- utils.ES_Success
			} else if strings.HasPrefix(s, "restart") {
				printRestartMsg()
				ec <- utils.ES_Restart
			}
			// else {
			//   // send command to rest of game
			// }
		}
	}()

	os.Exit(<-ec)
}

func processCmdLine() {
	fmt.Println("args: refs", os.Args)

	// // search for help argument
	// for _, cmd := range os.Args {
	//  if cmd == "help" {
	//      printHelp()
	//      return
	//  }
	// }

	for _, cmd := range os.Args[1:] {
		switch cmd {
		// case "launch_browser":
		// 	openBrowser("http://127.0.0.1:0")
		default:
			fmt.Printf("Invalid argument to game: %s\n", cmd)
			printHelp()
			return
		}
	}
}

func printStartMsg() {
	fmt.Println("starting game...")
}

func printExitMsg() {
	fmt.Println("exiting game...")
}

func printRestartMsg() {
	fmt.Println("restarting the game...")
}
func printHelp() {
	fmt.Println("help!")
}
