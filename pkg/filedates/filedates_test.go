package filedates

import (
	"testing"
)

func TestParseDate(t *testing.T) {
	cases := []struct {
		input string
		year  int
		month int
		day   int
	}{
		{"31-12-2022", 2022, 12, 31},
		{"2022-12-31", 2022, 12, 31},
		{"01-02-2022", 2022, 2, 1},
		{"02-01-2022", 2022, 1, 2},
	}

	for _, c := range cases {
		got, err := ParseDate(c.input)
		if err != nil {
			t.Fatalf("ParseDate(%s) returned error: %v", c.input, err)
		}
		if got.Year() != c.year || int(got.Month()) != c.month || got.Day() != c.day {
			t.Fatalf("ParseDate(%s) = %v, want %04d-%02d-%02d", c.input, got, c.year, c.month, c.day)
		}
	}
}

func TestParseDateAmbiguous(t *testing.T) {
	cases := []string{"03-04-22", "11-10-22"}
	for _, input := range cases {
		if _, err := ParseDate(input); err == nil {
			t.Fatalf("expected error for %s", input)
		}
	}
}

func TestFixDateInString(t *testing.T) {
	cases := []struct {
		input  string
		expect string
	}{
		{"report_31-12-2022.txt", "report_2022-12-31.txt"},
		{"report_2022-12-31.txt", "report_2022-12-31.txt"},
		{"notes 01-02-2022.md", "notes 2022-02-01.md"},
		{"notes 02-01-2022.md", "notes 2022-01-02.md"},
	}

	for _, c := range cases {
		got, err := FixDateInString(c.input)
		if err != nil {
			t.Fatalf("FixDateInString(%s) returned error: %v", c.input, err)
		}
		if got != c.expect {
			t.Fatalf("FixDateInString(%s) = %s, want %s", c.input, got, c.expect)
		}
	}
}

func TestFixDateInStringAmbiguous(t *testing.T) {
	cases := []string{"report_03-04-22.txt", "report_11-10-22.txt"}
	for _, input := range cases {
		if _, err := FixDateInString(input); err == nil {
			t.Fatalf("expected error for %s", input)
		}
	}
}
