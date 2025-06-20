package cmd

import (
	"bytes"
	"testing"
)

func TestRunNormalize(t *testing.T) {
	in := bytes.NewBufferString("report_12-31-2022.txt\n")
	out := &bytes.Buffer{}
	if err := runNormalize(in, out); err != nil {
		t.Fatalf("runNormalize error: %v", err)
	}
	expect := "report_2022-12-31.txt\n"
	if out.String() != expect {
		t.Fatalf("expected %q, got %q", expect, out.String())
	}
}
