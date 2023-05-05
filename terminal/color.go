package terminal

import (
    "fmt"
    "regexp"
    "strings"
)

func ColorString(color string, text string) string {
    return color + text + COLOR_END
}

func PrintTemplate(input string) {
    fmt.Println(ColorStringTemplate(input))
}

func ColorStringTemplate(input string) string {
    re := regexp.MustCompile(`{{\s*(([^{}]+)*)\s*\|\s*(\w+)\s*}}`)
    matches := re.FindAllStringSubmatch(input, -1)

    for _, match := range matches {
        text := strings.TrimSpace(match[1])
        color := Color(match[3])

        output := ColorString(color, text)
        input = strings.ReplaceAll(input, match[0], output)
    }

    return input
}

func Color(color string) string {
    switch color {
    case "end":
        return COLOR_END
    case "bold":
        return COLOR_BOLD
    case "italic":
        return COLOR_ITALIC
    case "url":
        return COLOR_URL
    case "blink":
        return COLOR_BLINK
    case "blink2":
        return COLOR_BLINK2
    case "selected":
        return COLOR_SELECTED
    case "black":
        return COLOR_BLACK
    case "red":
        return COLOR_RED
    case "green":
        return COLOR_GREEN
    case "yellow":
        return COLOR_YELLOW
    case "blue":
        return COLOR_BLUE
    case "violet":
        return COLOR_VIOLET
    case "beige":
        return COLOR_BEIGE
    case "white":
        return COLOR_WHITE
    case "blackbg":
        return COLOR_BLACKBG
    case "redbg":
        return COLOR_REDBG
    case "greenbg":
        return COLOR_GREENBG
    case "yellowbg":
        return COLOR_YELLOWBG
    case "bluebg":
        return COLOR_BLUEBG
    case "violetbg":
        return COLOR_VIOLETBG
    case "beigebg":
        return COLOR_BEIGEBG
    case "whitebg":
        return COLOR_WHITEBG
    case "grey":
        return COLOR_GREY
    case "red2":
        return COLOR_RED2
    case "green2":
        return COLOR_GREEN2
    case "yellow2":
        return COLOR_YELLOW2
    case "blue2":
        return COLOR_BLUE2
    case "violet2":
        return COLOR_VIOLET2
    case "beige2":
        return COLOR_BEIGE2
    case "white2":
        return COLOR_WHITE2
    case "greybg":
        return COLOR_GREYBG
    case "redbg2":
        return COLOR_REDBG2
    case "greenbg2":
        return COLOR_GREENBG
    case "yellowbg2":
        return COLOR_YELLOWBG2
    case "bluebg2":
        return COLOR_BLUEBG2
    case "violetbg2":
        return COLOR_VIOLETBG2
    case "beigebg2":
        return COLOR_BEIGEBG2
    case "whitebg2":
        return COLOR_WHITEBG2
    default:
        return ""
    }
}

const (
    COLOR_END      = "\033[0m"
    COLOR_BOLD     = "\033[1m"
    COLOR_ITALIC   = "\033[3m"
    COLOR_URL      = "\033[4m"
    COLOR_BLINK    = "\033[5m"
    COLOR_BLINK2   = "\033[6m"
    COLOR_SELECTED = "\033[7m"

    COLOR_BLACK  = "\033[30m"
    COLOR_RED    = "\033[31m"
    COLOR_GREEN  = "\033[32m"
    COLOR_YELLOW = "\033[33m"
    COLOR_BLUE   = "\033[34m"
    COLOR_VIOLET = "\033[35m"
    COLOR_BEIGE  = "\033[36m"
    COLOR_WHITE  = "\033[37m"

    COLOR_BLACKBG  = "\033[40m"
    COLOR_REDBG    = "\033[41m"
    COLOR_GREENBG  = "\033[42m"
    COLOR_YELLOWBG = "\033[43m"
    COLOR_BLUEBG   = "\033[44m"
    COLOR_VIOLETBG = "\033[45m"
    COLOR_BEIGEBG  = "\033[46m"
    COLOR_WHITEBG  = "\033[47m"

    COLOR_GREY    = "\033[90m"
    COLOR_RED2    = "\033[91m"
    COLOR_GREEN2  = "\033[92m"
    COLOR_YELLOW2 = "\033[93m"
    COLOR_BLUE2   = "\033[94m"
    COLOR_VIOLET2 = "\033[95m"
    COLOR_BEIGE2  = "\033[96m"
    COLOR_WHITE2  = "\033[97m"

    COLOR_GREYBG    = "\033[100m"
    COLOR_REDBG2    = "\033[101m"
    COLOR_GREENBG2  = "\033[102m"
    COLOR_YELLOWBG2 = "\033[103m"
    COLOR_BLUEBG2   = "\033[104m"
    COLOR_VIOLETBG2 = "\033[105m"
    COLOR_BEIGEBG2  = "\033[106m"
    COLOR_WHITEBG2  = "\033[107m"
)
