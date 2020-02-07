package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Execute runs diff command
func Execute() {
	diffCmd.Execute()
}

var diffCmd = &cobra.Command{
	Use:     "npm-diff",
	Short:   "npm package.json file dependency difference",
	Example: "npm-diff ./package.json ../path/to/other/package.json",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("invalid args: must supply two files")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		deps, err := readPackageFile(args[0])
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("read first file: %s", args[0]))
		}
		deps2, err := readPackageFile(args[1])
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("read second file: %s", args[1]))
		}
		listDiff(deps, deps2)
		return nil
	},
}
