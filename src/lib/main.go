package tasker

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type Manifest struct {
	Name    string                 `json:"name"`
	Version string                 `json:"version"`
	RunAs   string                 `json:"run_as"`
	LogFile string                 `json:"log_file"`
	Readme  string                 `json:"readme"`
	Env     map[string]string      `json:"environment"`
	Config  map[string]interface{} `json:"config"`
}

const (
	MANIFEST_NAME = ".tasker"
	TAR_PATH      = "/data/tasker/tars/"
	INTERNAL_PATH = "/data/tasker/.internal"

	DAEMON_STDOUT_FILE = "/data/tasker/.internal/DAEMON_STDOUT"
	DAEMON_STDIN_FILE  = "/data/tasker/.internal/DAEMON_STDIN"
	DAEMON_PID_FILE    = "/data/tasker/.internal/PID"
)

var (
	debug = flag.Bool("debug", false, "Whether or not bebug is on")
)

func getManifestBytes(taskPath string) ([]byte, error) {
	manifestPath := fmt.Sprintf("%s/%s", taskPath, MANIFEST_NAME)
	return ioutil.ReadFile(manifestPath)
}

func getInput() string {
	var response string
	fmt.Scanln(&response)
	return response
}

func Setup() {
	DebugPrintf("Setting up Tasker...")
	Fatalize(exec.Command("mkdir", "-p", TAR_PATH).Run())
	Fatalize(exec.Command("mkdir", "-p", INTERNAL_PATH).Run())
}

func Fatalize(err error) {
	if err != nil {
		fmt.Printf("Encountered an error: %s\n", err)
		os.Exit(1)
	}
}

func DebugPrintf(format string, a ...interface{}) {
	if *debug {
		fmt.Printf(format, a)
	}
}

func TellDaemon(command string) {
	// stdinPath := fmt.Sprintf("/proc/%s/fd/0")
	// if _, err := os.Stat(filename); os.IsNotExist(err) {
	//   fmt.Printf("no such file or directory: %s", filename)
	//   return
	// }
	// }
	// ioutil.WriteFile(stdinPath, command, 0777)
}

func GetManifest(taskPath string) *Manifest {
	manifestBytes, err := getManifestBytes(taskPath)
	if err != nil {
		DebugPrintf("Error while getting manifest bytes: %s\n", err)
		return nil
	}
	manifest := &Manifest{}
	err = json.Unmarshal(manifestBytes, manifest)
	if err != nil {
		DebugPrintf("Error while unmarshaling manifest: %s\n", err)
		return nil
	}
	return manifest
}

func GetOrCreateManifest(taskPath string) *Manifest {
	manifest := GetManifest(taskPath)
	if manifest == nil {
		fmt.Printf("The task you specified does not have a valid manifest\n")
		fmt.Printf("Do you want us to make one for you? [y/N]\n")
		switch getInput() {
		case "y", "yes", "yea", "Y", "YES", "SURE", "YEAH", "YEA":
			break
			os.Exit(0)
		}
		// TODO Create manifest
	}
	return manifest
}
