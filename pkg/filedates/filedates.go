package filedates

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

var dateRegex = regexp.MustCompile(`\d{2}[\s-]+\d{2}[\s-]+\d{2,4}|\d{4}[\s-]+\d{2}[\s-]+\d{2}`)

// FixDateInString replaces the first date found in the input string with its ISO format.
func FixDateInString(input string) (string, error) {
	idx := dateRegex.FindStringIndex(input)
	if idx == nil {
		return "", errors.New("no date found in: '" + input + "'")
	}
	dateStr := input[idx[0]:idx[1]]
	dateStr = strings.ReplaceAll(dateStr, " ", "-")
	dateStr = strings.ReplaceAll(dateStr, "_", "-")
	dateStr = regexp.MustCompile(`-+`).ReplaceAllString(dateStr, "-")

	t, err := ParseDate(dateStr)
	if err != nil {
		return "", err
	}
	iso := t.Format("2006-01-02")
	return strings.Replace(input, dateStr, iso, 1), nil
}

func ParseDate(s string) (time.Time, error) {
	layouts := []string{"02-01-2006", "01-02-2006", "2006-01-02", "06-01-02", "02-01-06", "01-02-06"}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unrecognized date: %s", s)
}
