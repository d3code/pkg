package xerr

import (
    "github.com/d3code/clog"
    "os"
    "os/exec"
    "strings"
)

func ExitIfError(err error) {
    if err != nil {
        if exitErr, ok := err.(*exec.ExitError); ok {
            stderr := string(exitErr.Stderr)
            clog.Error(strings.TrimSuffix(stderr, "\n"))
        } else {
            s := err.Error()
            if s == "^C" {
                os.Exit(0)
            }
            clog.Error(s)
        }
        os.Exit(1)
    }
}
