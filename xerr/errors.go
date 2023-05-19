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
            clog.Error(strings.TrimSuffix(string(exitErr.Stderr), "\n"))
        } else {
            clog.Error(err.Error())
        }
        os.Exit(1)
    }
}
