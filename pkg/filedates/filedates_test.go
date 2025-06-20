package filedates

import (
	"testing"
)

func TestParseDate(t *testing.T) {
	cases := []struct{ input string }{
		{"31-12-2022"},
		{"12-31-2022"},
		{"22-12-31"},
	}
	for _, c := range cases {
		got, err := ParseDate(c.input)
		if err != nil {
			t.Fatalf("ParseDate(%s) returned error: %v", c.input, err)
		}
		if got.Year() != 2022 || got.Month() != 12 || got.Day() != 31 {
			t.Fatalf("ParseDate(%s) = %v", c.input, got)
		}
	}
}

func TestFixDateInString(t *testing.T) {
	cases := []struct {
		input  string
		expect string
	}{
		{"report_31-12-2022.txt", "report_2022-12-31.txt"},
		{"report_12-31-2022.txt", "report_2022-12-31.txt"},
		{"report_2022-12-31.txt", "report_2022-12-31.txt"},
		{"report_22-12-31.txt", "report_2022-12-31.txt"},
		{"report_12-31-22.txt", "report_2022-12-31.txt"},
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
