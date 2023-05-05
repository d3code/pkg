package terminal

import (
    "fmt"
    "strings"
)

func PrintList(list []string) {
    for _, item := range list {
        fmt.Println("[", item, "]")
    }
}

func PrintHeader(header string) {
    num := 50 - len(header)
    padding := strings.Repeat("-", num)

    mod := fmt.Sprintf("-- %s %s", header, padding)
    fmt.Println(mod)
}
