package main

import (
	tasker "../lib"

	"archive/tar"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	dir = flag.Bool("dir", false, "Indicates that you wish to act on a directory")
)

func handleTar() {
	if len(flag.Args()) < 2 {
		fmt.Printf("Not enough arguments\n")
		os.Exit(1)
	}
	taskPath := flag.Args()[1]
	fmt.Printf("Creating tar from %s\n", taskPath)

	manifest := tasker.GetOrCreateManifest(taskPath)
	tarName := fmt.Sprintf("%s/%s-%s.tar.gz", tasker.TAR_PATH, manifest.Name, manifest.Version)

	err := TarFile(tarName, taskPath)
	if err != nil {
		log.Fatalln(err)
	}
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

func TarFile(tarName, filePath string) error {
	buf := new(bytes.Buffer)
	tarWriter := tar.NewWriter(buf)

	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	parts := strings.Split(filePath, "/")
	header := &tar.Header{
		Name: parts[len(parts)-1],
		Size: int64(len(body)),
	}
	if err := tarWriter.WriteHeader(header); err != nil {
		return err
	}
	if _, err := tarWriter.Write(body); err != nil {
		return err
	}
	if err := tarWriter.Close(); err != nil {
		return err
	}

	if err := ioutil.WriteFile(tarName, buf.Bytes(), 0644); err != nil {
		return err
	}
	fmt.Println("Done")
	return nil
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
