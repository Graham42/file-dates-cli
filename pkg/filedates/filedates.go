package filedates

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/itlightning/dateparse"
)

var dayPattern = `(?:\d{1,2}(?:st|nd|rd|th)?)`
var monthPattern = `(?:Jan(?:uary)?|Feb(?:ruary)?|Mar(?:ch)?|Apr(?:il)?|May|Jun(?:e)?|Jul(?:y)?|Aug(?:ust)?|Sept?(?:ember)?|Oct(?:ober)?|Nov(?:ember)?|Dec(?:ember)?)`
var yearPattern = `(?:\d{2,4})`

// could be year month day, month day year, day month year, or numeric date
var datePattern = fmt.Sprintf("\b(?:%s.*%s.*%s|%s.*%s.*%s|%s.*%s.*%s|\\d{1,4}[-_ ,]+\\d{1,2}[-_ ,]+\\d{1,4})\b",
	dayPattern, monthPattern, yearPattern,
	monthPattern, dayPattern, yearPattern,
	yearPattern, monthPattern, dayPattern)
var dateRegex = regexp.MustCompile(datePattern)

// FixDateInString replaces the first date found in the input string with its ISO format.
func FixDateInString(input string) (string, error) {
	idx := dateRegex.FindStringIndex(input)
	if idx == nil {
		return "", fmt.Errorf("no date found in: '%s'", input)
	}
	dateStr := input[idx[0]:idx[1]]

	t, err := ParseDate(dateStr)
	if err != nil {
		return "", err
	}
	iso := t.Format("2006-01-02")
	return strings.Replace(input, dateStr, iso, 1), nil
}

func ParseDate(s string) (time.Time, error) {
	return dateparse.ParseStrict(s)
}
