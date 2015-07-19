package main

import (
  "../lib"

  "flag"
  "fmt"
  "os/exec"
)

var (
  debug = flag.Bool("debug", false, "Whether or not bebug is on")
)

func debugPrintf(format string, a ...interface{}) {
  if (*debug) {
    fmt.Printf(format, a)
  }
}

func getInput() string {
    var response string
    fmt.Scanln(&response)
    return response
}

// tasker tar .
func hanndleTar() {
  if len(flag.Args()) < 2 {
    fmt.Printf("Not enough arguments to the tasker tar command\n")
  }
  taskPath := flag.Args()[1]
  taskName := ""
  taskVersion := ""

  manifest, err := tasker.GetOrCreateManifest(taskPath)
   else {
    tasker.ValidateManifest(manifest)

  }
  
  tarName := fmt.Sprintf("%s-%s.tar.gz", taskName, taskVersion)
  cmd := exec.Command("tar", "-czf", tarName, taskPath)
}

func main() {
  debugPrintf("*****Tasker CLI")
  flag.Parse()
  if len(flag.Args()) < 1 {
    fmt.Printf("You did not provide a command to execute.\n")
  }
  switch flag.Args()[0] {
    case "tar":
      handleTar()
  }
}
