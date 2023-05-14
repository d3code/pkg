package clog

import (
    "fmt"
    "strings"
)

func Info(input ...string) {
    message := strings.Join(input, " ")
    message = colorMatchTemplate(message)

    fmt.Println(message)
}

func InfoF(format string, input ...any) {
    message := fmt.Sprintf(format, input...)
    message = colorMatchTemplate(message)

    fmt.Println(message)
}

func InfoL(inputLines ...string) {
    message := strings.Join(inputLines, "\n")
    message = colorMatchTemplate(message)

    fmt.Println(message)
}

func Debug(title string, message string) {
    message = removeColor(message)
    message = fmt.Sprintf("[ %s ] %s", title, message)
    message = ColorString(message, "grey")

    Info(message)
}

func Underline(message ...string) {
    title := strings.Join(message, " ")
    title = colorMatchTemplate(title)

    underline := removeColor(title)
    underline = strings.Repeat("-", len(underline))

    fmt.Println()
    InfoL(title, underline)
}

func UnderlineF(format string, input ...any) {
    title := fmt.Sprintf(format, input...)
    title = colorMatchTemplate(title)

    underline := removeColor(title)
    underline = strings.Repeat("-", len(underline))

    fmt.Println()
    InfoL(title, underline)
}
