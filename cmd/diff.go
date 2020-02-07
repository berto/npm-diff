package cmd

import (
	"github.com/fatih/color"
)

func listDiff(deps, deps2 map[string]string) {
	for dep, version := range deps {
		if deps2[dep] != "" && deps2[dep] != version {
			color.White(dep)
			color.Green("\t" + version)
			color.Red("\t" + deps2[dep])
		}
	}
}
