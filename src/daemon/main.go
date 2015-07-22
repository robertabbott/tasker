package main

import (
	"fmt"
	"io/ioutil"
	"os"

	tasker "../lib"
)

func listenStdIn() {
	var line string
	for {
		fmt.Scanln(&line)
		fmt.Println(line)
		switch line {
		case "kill":
			fmt.Println("Daemon exiting")
			os.Exit(0)
		}
	}
}

func writePID() {
	pid := fmt.Sprintf("%d", os.Getpid())
	ioutil.WriteFile(tasker.DAEMON_PID_FILE, []byte(pid), 0777)
}

func main() {
	tasker.Setup()
	fmt.Println("Daemon is starting")
	writePID()
	listenStdIn()
}
