package builder

import (
	"os"
	"testing"
	"time"

	"github.com/rogpeppe/go-internal/testscript"
)

// setup handles following setup errors
//
// Error:
//   - error: creating directory '/no-home/.cache/nix': Read-only file system
func setup(env *testscript.Env) error {
	// Grab specific env details from current system &
	// pass them to testscript
	env.Vars = append(
		env.Vars,
		"HOME="+os.Getenv("HOME"),
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
