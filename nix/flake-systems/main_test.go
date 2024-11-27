package flake_systems

import (
	"os"
	"testing"
	"time"

	"github.com/rogpeppe/go-internal/testscript"
)

// setup handles following setup errors
//
// Error:
//   - path '$WORK' does not contain a 'flake.nix'
//   - error: experimental Nix feature 'nix-command' is disabled; add '--extra-experimental-features nix-command' to enable it
func setup(env *testscript.Env) error {
	// Grab specific env details from current system &
	// pass them to testscript
	env.Vars = append(
		env.Vars,
		"HOME="+os.Getenv("HOME"),
		"PWD="+os.Getenv("PWD"),
	)
	return nil
}

func Test(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir:                  "testdata",
		Files:                nil,
		Setup:                setup,
		Condition:            nil,
		Cmds:                 nil,
		TestWork:             false,
		WorkdirRoot:          "",
		IgnoreMissedCoverage: false,
		UpdateScripts:        false,
		RequireExplicitExec:  false,
		RequireUniqueNames:   false,
		ContinueOnError:      false,
		Deadline:             time.Time{},
	})
}
