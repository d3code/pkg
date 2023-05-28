package shell

import (
    "github.com/d3code/clog"
    "os"
    "strings"
)

var home string

func init() {
    dir, err := os.UserHomeDir()
    if err != nil {
        clog.Info("Could not get user home directory")
        os.Exit(1)
    }

    home = dir
}

func FullPath(directory string) string {
    var path string

    if directory == "" || directory == "." {
        path = CurrentDirectory()
    } else if strings.HasPrefix(directory, "/") {
        path = directory
    } else {
        path = CurrentDirectory() + "/" + directory
    }

    return path
}

func UserHomeDirectory() string {
    return home
}

func CurrentDirectory() string {
    pwd := RunCmd(".", false, "pwd")
    return strings.TrimSuffix(pwd.Stdout, "\n")
}
