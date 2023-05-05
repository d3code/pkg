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

func FatalErrorMessage(err error, message string, printError bool) {
    if err != nil {
        fmt.Println(message)
        if printError {
            fmt.Println(err)
        }
        os.Exit(1)
    }
}

func Fatal(err string) {
    fmt.Println(err)
    os.Exit(1)
}
