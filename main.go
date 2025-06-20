package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: <command>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "normalize":
		if err := runNormalize(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func runNormalize() error {
	input, err := readStdin()
	if err != nil {
		return err
	}
	filename := strings.TrimSpace(input)
	if filename == "" {
		return errors.New("no input provided")
	}
	normalized, err := normalizeFilename(filename)
	if err != nil {
		return err
	}
	fmt.Println(normalized)
	return nil
}

func readStdin() (string, error) {
	// TODO: read one line at a time instead of all input
	b, err := io.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// TODO: refine this regex to support more date formats
var dateRE = regexp.MustCompile(`\d{2}-\d{2}-\d{2,4}`)

func normalizeFilename(name string) (string, error) {
	loc := dateRE.FindStringIndex(name)
	if loc == nil {
		return "", errors.New("no date found")
	}
	dateStr := name[loc[0]:loc[1]]
	t, err := parseDate(dateStr)
	if err != nil {
		return "", err
	}
	iso := t.Format("2006-01-02")
	return strings.Replace(name, dateStr, iso, 1), nil
}

func parseDate(s string) (time.Time, error) {
	layouts := []string{"02-01-2006", "01-02-2006", "06-01-02"}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unrecognized date: %s", s)
}
