package cmd

import (
	"strings"

	"github.com/fatih/color"
)

const (
	// Patch releases: 1.0 or 1.0.x or ~1.0.4
	patch = iota
	// Minor releases: 1 or 1.x or ^1.0.4
	minor
	// Major releases: * or x
	major
)

func parseVersion(version string) (string, int) {
	if version == "" {
		return version, major
	}
	switch version[0] {
	case '~':
		return version[1:], patch
	case '^':
		return version[1:], minor
	default:
		return version, major
	}
}

func fillVersion(version string) (newVersion []string) {
	splitVersion := strings.Split(version, ".")
	switch len(splitVersion) {
	case 1:
		newVersion = append(newVersion, splitVersion[0])
		newVersion = append(newVersion, "*")
		newVersion = append(newVersion, "*")
	case 2:
		newVersion = append(newVersion, splitVersion[0])
		newVersion = append(newVersion, splitVersion[1])
		newVersion = append(newVersion, "*")
	default:
		newVersion = append(newVersion, splitVersion[0])
		newVersion = append(newVersion, splitVersion[1])
		newVersion = append(newVersion, splitVersion[2])
	}
	return
}

func simplify(version string, updateType int) string {
	newVersion := fillVersion(version)
	switch updateType {
	case patch:
		newVersion[2] = "*"
	case minor:
		newVersion[1] = "*"
		newVersion[2] = "*"
	case major:
		newVersion[0] = "*"
	}
	return strings.Join(newVersion, ".")
}

func diff(v1, v2 string) bool {
	version1, type1 := parseVersion(v1)
	version2, type2 := parseVersion(v2)
	if type1 != type2 {
		return true
	}
	c1, c2 := simplify(version1, type1), simplify(version2, type2)
	if c1 != c2 {
		return true
	}
	return false
}

func listDiff(strict bool, deps1, deps2 map[string]string) {
	noChange := true
	for dep, version := range deps1 {
		version2 := deps2[dep]
		if version2 != "" && ((strict && version != version2) || diff(version, version2)) {
			noChange = false
			color.White(dep)
			color.Green("\t" + version)
			color.Red("\t" + version2)
		}
	}
	if noChange {
		color.White("No Differences Found!")
	}
}
