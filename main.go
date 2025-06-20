package main

import "github.com/example/file-dates-cli/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		// Cobra already prints the error message
	}
}
