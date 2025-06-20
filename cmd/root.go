package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "file-dates",
	Short: "CLI tools for working with filenames containing dates",
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}
