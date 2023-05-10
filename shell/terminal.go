package shell

import (
    "github.com/d3code/pkg/errors"
    "os"
    "os/exec"
    "strings"
)

func RunOut(name string, args ...string) {
    command := exec.Command(name, args...)
    command.Stdout = os.Stdout
    command.Stderr = os.Stderr

    err := command.Run()
    errors.ExitIfError(err)
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

    errors.ExitIfError(err)
    return string(output)
}

func RunE(name string, args ...string) (string, error) {
    command := exec.Command(name, args...)
    output, err := command.Output()

    return string(output), err
}

func CurrentDirectory() string {
    pwd := Run("pwd")
    return strings.TrimSuffix(pwd, "\n")
}
