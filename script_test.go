package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestScripts(t *testing.T) {
	root, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
		Setup: func(env *testscript.Env) error {
			exe := filepath.Join(env.WorkDir, "file-dates")
			cmd := exec.Command("go", "build", "-o", exe, root)
			if out, err := cmd.CombinedOutput(); err != nil {
				return fmt.Errorf("build error: %v\n%s", err, out)
			}
			env.Setenv("CLI", exe)
			return nil
		},
	})
}
