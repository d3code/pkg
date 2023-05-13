package shell

import (
    "fmt"
    "github.com/d3code/pkg/xerr"
    "os"
    "os/exec"
    "strings"
)

func RunOut(name string, args ...string) {
    command := exec.Command(name, args...)
    command.Stdout = os.Stdout
    command.Stderr = os.Stderr

    err := command.Run()
    xerr.ExitIfError(err)
}

func RunOutE(name string, args ...string) error {
    command := exec.Command(name, args...)
    command.Stdout = os.Stdout
    command.Stderr = os.Stderr

    return command.Run()
}

func Run(name string, args ...string) string {
    command := exec.Command(name, args...)
    output, err := command.Output()

    xerr.ExitIfError(err)
    out := string(output)
    return strings.TrimSuffix(out, "\n")
}

func RunE(name string, args ...string) (string, error) {
    command := exec.Command(name, args...)
    output, err := command.Output()

    out := string(output)
    return strings.TrimSuffix(out, "\n"), err
}

func RunDirE(path string, name string, args ...string) (string, error) {
    command := exec.Command(name, args...)
    command.Dir = path

    output, err := command.Output()
    if err != nil {
        fmt.Println(err)
        return "", err
    }

    out := string(output)
    return strings.TrimSuffix(out, "\n"), nil
}

func RunDir(path string, name string, args ...string) string {
    command := exec.Command(name, args...)
    command.Dir = path

    output, err := command.Output()

    xerr.ExitIfError(err)
    out := string(output)
    return strings.TrimSuffix(out, "\n")
}

func RunOutDir(path string, name string, args ...string) {
    command := exec.Command(name, args...)
    command.Stdout = os.Stdout
    command.Stderr = os.Stderr
    command.Dir = path

    err := command.Run()
    xerr.ExitIfError(err)
}

func RunShell(args ...string) string {
    osShell := os.Getenv("SHELL")
    args = append([]string{"-c"}, args...)
    command := exec.Command(osShell, args...)

    output, err := command.Output()

    xerr.ExitIfError(err)
    out := string(output)
    return strings.TrimSuffix(out, "\n")
}
