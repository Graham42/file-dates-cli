package cmd

import (
	"bytes"
	"testing"
)

func TestRunNormalize(t *testing.T) {
	in := bytes.NewBufferString("report_31-12-2022.txt\n")
	out := &bytes.Buffer{}
	if err := runNormalize(in, out); err != nil {
		t.Fatalf("runNormalize error: %v", err)
	}
	expect := "report_2022-12-31.txt\n"
	if out.String() != expect {
		t.Fatalf("expected %q, got %q", expect, out.String())
	}
}

func TestRunNormalizeAmbiguous(t *testing.T) {
	in := bytes.NewBufferString("report_03-04-22.txt\n")
	if err := runNormalize(in, &bytes.Buffer{}); err == nil {
		t.Fatalf("expected error for ambiguous date")
	}
}
