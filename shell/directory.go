package shell

import (
    "strings"
)

func FullPath(directory string) string {
    var path string
    if strings.HasPrefix(directory, "/") {
        path = directory
    } else {
        path = CurrentDirectory() + "/" + directory
    }
    return path
}
