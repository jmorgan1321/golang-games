package main

import (
	"flag"
	"fmt"
	"github.com/jmorgan1321/golang-games/core/utils"
	"os"
	"os/exec"
	"syscall"
)

var (
	gameFlag = flag.String("game", "refs", "which game to launch")
	urlFlag  = flag.String("url", "http://127.0.0.1", "which url to serve the game from")
	portFlag = flag.String("port", "8080", "which port to communicate with the game on")
	rootFlag = flag.String("root", ".", "Define the root filesystem path")
)

func init() {
	flag.Usage = printHelp
}

func main() {
	flag.Parse()

	if len(os.Args) == 1 {
		printHelp()
		return
	}

	// Find the game the user specified.
	path, err := exec.LookPath(*gameFlag)
	if err != nil {
		fmt.Println("path err:", err)
		return
	}

	// If this is the first time launching the game (ie, not a restart), then
	// we want to tell the game to open a browser.
	// Otherwise, we'll just keep looping and restarting the game to similate
	// hot-code reloading.
	first, launch := true, true
	for launch {
		launch = false // reset the launch flag

		// Collect all of the non flag args and pass them into the game,
		// passing in the launch_browser flag if this is the first time we
		// started the game.
		cmds := flag.Args()
		if first {
			first = false
			openBrowser(url)
			cmds = append(cmds, "url="+url)
		}

		// Connect the input and output streams to the game's sub-process.
		cmd := exec.Command(path, cmds...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Start()
		if err != nil {
			fmt.Println("err", err)
		}

		err = cmd.Wait()
		// Check the error code of the process, to see if the game needs to
		// be restarted.
		if err != nil {
			ec := err.(*exec.ExitError).ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
			if ec == utils.ES_Restart {
				launch = true
			} else {
				fmt.Printf("\tgame err: %s\n", err)
			}
		}
	}
}

func printHelp() {
	desc := `Launches the game; restarting it if necessary.`

	fmt.Printf("\nDesc :\t%s\n", desc)
	fmt.Printf("\nUsage:\t%s [game to run flag] <commands...>\n\n", os.Args[0])
	fmt.Println("flags: (-flag=default: desc)")
	flag.PrintDefaults()
}
