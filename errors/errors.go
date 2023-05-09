package errors

import (
    "fmt"
    "os"
)

func ExitIfError(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func ExitIfErrorWithMessage(err error, message string, printError bool) {
    if err != nil {
        fmt.Println(message)
        if printError {
            fmt.Println(err)
        }
        os.Exit(1)
    }
}

func Exit(message string) {
    fmt.Println(message)
    os.Exit(1)
}
