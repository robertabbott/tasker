package main

import (
	tasker "../lib"

	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	dir = flag.Bool("dir", false, "Indicates that you wish to act on a directory")
)

// tasker tar .
func handleTar() {
	if len(flag.Args()) < 2 {
		fmt.Printf("Not enough arguments\n")
		os.Exit(1)
	}
	taskPath := flag.Args()[1]
	fmt.Printf("Creating tar from %s\n", taskPath)

	manifest := tasker.GetOrCreateManifest(taskPath)
	tarName := fmt.Sprintf("%s/%s-%s.tar.gz", tasker.TAR_PATH,
		manifest.Name, manifest.Version)
	cmd := exec.Command("tar", "-czf", tarName, taskPath)
	out, err := cmd.CombinedOutput()
	tasker.DebugPrintf("Output of tar command: %s\n", string(out))
	tasker.Fatalize(err)
	fmt.Println("Done")
}

func handleCreate() {
	if !*dir {
		// Need to tar it up first
		handleTar()
	}
	// TODO
}

func handleStart() {

}

func handleStop() {

}

func handleRestart() {
	handleStop()
	handleStart()
}

func handleStatus() {

}

func handleLog() {

}

func handleUptime() {

}

func handleTell() {

}

func handleDestroy() {

}

func handleDaemon() {
	if len(flag.Args()) < 2 {
		fmt.Printf("Not enough arguments\n")
		os.Exit(1)
	}
	command := flag.Args()[1]
	tasker.TellDaemon(fmt.Sprintf("%s\n", command))
}

func handleNoCommand() {
	fmt.Println("No command was provided...")
}

func handleArgs() {
	switch flag.Args()[0] {
	case "tar":
		handleTar()
	case "create":
		handleCreate()
	case "start":
		handleStart()
	case "stop":
		handleStop()
	case "restart":
		handleRestart()
	case "destroy":
		handleDestroy()
	case "status":
		handleStatus()
	case "log":
		handleLog()
	case "uptime":
		handleUptime()
	case "tell":
		handleTell()
	case "daemon":
		handleDaemon()
	default:
		handleNoCommand()
	}
}

func main() {
	tasker.Setup()
	tasker.DebugPrintf("*****Tasker CLI")
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Printf("You did not provide a command to execute.\n")
		return
	}
	handleArgs()
}
