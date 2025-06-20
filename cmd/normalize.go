package cmd

import (
	"bufio"
	"errors"
	"github.com/example/file-dates-cli/pkg/filedates"
	"github.com/spf13/cobra"
	"io"
	"strings"
)

func init() {
	rootCmd.AddCommand(normalizeCmd)
}

var normalizeCmd = &cobra.Command{
	Use:   "normalize",
	Short: "Normalize dates in filenames to ISO 8601",
	RunE: func(cmd *cobra.Command, args []string) error {
		in := cmd.InOrStdin()
		out := cmd.OutOrStdout()
		return runNormalize(in, out)
	},
}

// runNormalize reads from in, normalizes the filename, and writes to out.
func runNormalize(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		return errors.New("no input provided")
	}
	filename := strings.TrimSpace(scanner.Text())

	if scanner.Scan() {
		return errors.New("expected single line of input")
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	normalized, err := filedates.FixDateInString(filename)
	if err != nil {
		return err
	}
	_, err = io.WriteString(out, normalized+"\n")
	return err
}
