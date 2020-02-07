package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type PackageFile struct {
	DevDependencies map[string]string `json:"devDependencies"`
	Dependencies    map[string]string `json:"dependencies"`
}

func readPackageFile(path string) (map[string]string, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "read json file")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var packageFile PackageFile
	json.Unmarshal(byteValue, &packageFile)
	for dep, version := range packageFile.DevDependencies {
		packageFile.Dependencies[dep] = version
	}
	return packageFile.Dependencies, nil
}
