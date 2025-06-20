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
	normalized, err := fixDateInString(filename)
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

// TODO improve the regex to ensure there's no extra digits on either side of the date
var dateRegex = regexp.MustCompile(`\d{2}[\s-]+\d{2}[\s-]+\d{2,4}|\d{4}[\s-]+\d{2}[\s-]+\d{2}`)

// fixDateInString replaces the first date found in the input string with its ISO format.
func fixDateInString(input string) (string, error) {
	dateMatchIndex := dateRegex.FindStringIndex(input)
	if dateMatchIndex == nil {
		return "", errors.New("no date found in: '" + input + "'")
	}
	dateStr := input[dateMatchIndex[0]:dateMatchIndex[1]]
	// normalize the separators to a single `-`
	dateStr = strings.ReplaceAll(dateStr, " ", "-")
	dateStr = strings.ReplaceAll(dateStr, "_", "-")
	// reduce multiple dashes to a single dash
	dateStr = regexp.MustCompile(`-+`).ReplaceAllString(dateStr, "-")

	t, err := parseDate(dateStr)
	if err != nil {
		return "", err
	}
	isoFormattedDate := t.Format("2006-01-02")
	return strings.Replace(input, dateStr, isoFormattedDate, 1), nil
}

func parseDate(s string) (time.Time, error) {
	layouts := []string{"02-01-2006", "01-02-2006", "2006-01-02",
		"06-01-02", "02-01-06", "01-02-06"}
	// TODO raise an error if the date could be ambiguous
	// e.g. 01-02-03 could be 2001-02-03 or 2003-01-02 or 2003-02-01
	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unrecognized date: %s", s)
}
