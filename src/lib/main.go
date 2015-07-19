package tasker

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
)

type Manifest struct {
  Name      string            `json:"name"`
  Version   string            `json:"version"`
  RunAs     string            `json:"run_as"`
  LogFile   string            `json:"log_file"`
  Readme    string            `json:"readme"`
  Env       map[string]string `json:"environment"`
  Config    map[string]interface{}   `json:"config"`
  Start     string    `json:"run_as"`
}

const (
  MANIFEST_NAME = ".tasker"
)

func getManifestBytes(taskPath string) ([]byte, error) {
  manifestPath := fmt.Sprintf("%s/%s", taskPath, MANIFEST_NAME)
  return ioutil.ReadFile(manifestPath)
}

func fatalize(err error) {
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func GetManifest(taskPath string) (*Manifest, error) {
  manifestBytes, err := getManifestBytes(taskPath)
  if err != nil {
    return nil, err
  }
  manifest := &Manifest{}
  err = json.Unmarshal(manifestBytes, manifest)
  return manifest, err
}

func GetOrCreateManifest(taskPath string) (*Manifest, error) {
  manifest, err := GetManifest(taskPath)
  if err != nil {
    fmt.Printf("The task you specified does not have a valid manifest\n")
    fmt.Printf("Do you want us to make one for you? [y/N]")
    switch getInput {
      case "y", "yes", "yea", "Y", "YES", "SURE", "YEAH", "YEA":
        break
      return
    }
    manifest, err = CreateDefaultManifest(taskPath)
  }
}

