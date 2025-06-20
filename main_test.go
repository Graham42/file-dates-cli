package main

import "testing"

func TestParseDate(t *testing.T) {
	cases := []struct {
		input string
	}{
		{"31-12-2022"},
		{"12-31-2022"},
		{"22-12-31"},
	}
	for _, c := range cases {
		got, err := parseDate(c.input)
		if err != nil {
			t.Fatalf("parseDate(%s) returned error: %v", c.input, err)
		}
		if got.Year() != 2022 || got.Month() != 12 || got.Day() != 31 {
			t.Fatalf("parseDate(%s) = %v", c.input, got)
		}
	}
}

func TestNormalizeFilename(t *testing.T) {
	got, err := normalizeFilename("report_12-31-2022.txt")
	if err != nil {
		t.Fatalf("normalizeFilename returned error: %v", err)
	}
	expect := "report_2022-12-31.txt"
	if got != expect {
		t.Fatalf("expected %s, got %s", expect, got)
	}
}
