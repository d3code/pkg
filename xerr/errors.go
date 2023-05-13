package xerr

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func ExitIfError(err error) {
    if err != nil {
        if exitErr, ok := err.(*exec.ExitError); ok {
            errorString := string(exitErr.Stderr)
            singleLineError := strings.TrimSuffix(errorString, "\n")
            fmt.Println(singleLineError)
        }
        fmt.Println(err)
        os.Exit(1)
    }
}
