package shell

import (
    "fmt"
    "regexp"
    "strings"
)

func PrintTemplate(input string) {
    re := regexp.MustCompile(`{{\s*(([^{}]+)*)\s*\|\s*(\w+)\s*}}`)
    matches := re.FindAllStringSubmatch(input, -1)

    for _, match := range matches {
        text := strings.TrimSpace(match[1])
        color := matchColor(match[3])

        output := ColorString(text, color)
        input = strings.ReplaceAll(input, match[0], output)
    }

    fmt.Println(input)
}

func ColorString(text string, color string) string {
    if strings.HasPrefix(color, "\\033[") {
        color = matchColor(color)
    }

    return color + text + color_END
}

func matchColor(color string) string {
    switch color {
    case "end":
        return color_END
    case "bold":
        return color_BOLD
    case "italic":
        return color_ITALIC
    case "url":
        return color_URL
    case "blink":
        return color_BLINK
    case "blink2":
        return color_BLINK2
    case "selected":
        return color_SELECTED
    case "black":
        return color_BLACK
    case "red":
        return color_RED
    case "green":
        return color_GREEN
    case "yellow":
        return color_YELLOW
    case "blue":
        return color_BLUE
    case "violet":
        return color_VIOLET
    case "beige":
        return color_BEIGE
    case "white":
        return color_WHITE
    case "blackbg":
        return color_BLACKBG
    case "redbg":
        return color_REDBG
    case "greenbg":
        return color_GREENBG
    case "yellowbg":
        return color_YELLOWBG
    case "bluebg":
        return color_BLUEBG
    case "violetbg":
        return color_VIOLETBG
    case "beigebg":
        return color_BEIGEBG
    case "whitebg":
        return color_WHITEBG
    case "grey":
        return color_GREY
    case "red2":
        return color_RED2
    case "green2":
        return color_GREEN2
    case "yellow2":
        return color_YELLOW2
    case "blue2":
        return color_BLUE2
    case "violet2":
        return color_VIOLET2
    case "beige2":
        return color_BEIGE2
    case "white2":
        return color_WHITE2
    case "greybg":
        return color_GREYBG
    case "redbg2":
        return color_REDBG2
    case "greenbg2":
        return color_GREENBG
    case "yellowbg2":
        return color_YELLOWBG2
    case "bluebg2":
        return color_BLUEBG2
    case "violetbg2":
        return color_VIOLETBG2
    case "beigebg2":
        return color_BEIGEBG2
    case "whitebg2":
        return color_WHITEBG2
    default:
        return ""
    }
}

const (
    color_END      = "\033[0m"
    color_BOLD     = "\033[1m"
    color_ITALIC   = "\033[3m"
    color_URL      = "\033[4m"
    color_BLINK    = "\033[5m"
    color_BLINK2   = "\033[6m"
    color_SELECTED = "\033[7m"

    color_BLACK  = "\033[30m"
    color_RED    = "\033[31m"
    color_GREEN  = "\033[32m"
    color_YELLOW = "\033[33m"
    color_BLUE   = "\033[34m"
    color_VIOLET = "\033[35m"
    color_BEIGE  = "\033[36m"
    color_WHITE  = "\033[37m"

    color_BLACKBG  = "\033[40m"
    color_REDBG    = "\033[41m"
    color_GREENBG  = "\033[42m"
    color_YELLOWBG = "\033[43m"
    color_BLUEBG   = "\033[44m"
    color_VIOLETBG = "\033[45m"
    color_BEIGEBG  = "\033[46m"
    color_WHITEBG  = "\033[47m"

    color_GREY    = "\033[90m"
    color_RED2    = "\033[91m"
    color_GREEN2  = "\033[92m"
    color_YELLOW2 = "\033[93m"
    color_BLUE2   = "\033[94m"
    color_VIOLET2 = "\033[95m"
    color_BEIGE2  = "\033[96m"
    color_WHITE2  = "\033[97m"

    color_GREYBG    = "\033[100m"
    color_REDBG2    = "\033[101m"
    color_GREENBG2  = "\033[102m"
    color_YELLOWBG2 = "\033[103m"
    color_BLUEBG2   = "\033[104m"
    color_VIOLETBG2 = "\033[105m"
    color_BEIGEBG2  = "\033[106m"
    color_WHITEBG2  = "\033[107m"
)
