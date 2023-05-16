package shell

import (
    "bytes"
    "github.com/d3code/clog"
    "io"
    "os"
    "os/exec"
    "strings"
)

type CommandResponse struct {
    Stdout string
    Stderr string
}

func RunCmdE(path string, stdout bool, program string, args ...string) (CommandResponse, error) {
    debugLogging(program, args)
    command := exec.Command(program, args...)
    if path != "" && path != "." {
        command.Dir = path
    }

    outBytes, errBytes := writer(command, stdout)
    err := command.Run()

    return CommandResponse{
        Stdout: strings.TrimSuffix(outBytes.String(), "\n"),
        Stderr: strings.TrimSuffix(errBytes.String(), "\n"),
    }, err
}

func RunCmd(path string, stdout bool, program string, args ...string) CommandResponse {
    commandResponse, err := RunCmdE(path, stdout, program, args...)
    if err != nil {
        if !stdout {
            if commandResponse.Stderr != "" {
                clog.Error(commandResponse.Stderr)
            } else if commandResponse.Stdout != "" {
                clog.Error(commandResponse.Stdout)
            } else {
                clog.Error(err.Error())
            }
        }
        os.Exit(1)
    }

    return commandResponse
}

func RunShell(stdout bool, args ...string) CommandResponse {
    osShell := os.Getenv("SHELL")
    args = append([]string{"-c"}, args...)

    return RunCmd(".", stdout, osShell, args...)
}

func writer(cmd *exec.Cmd, stdout bool) (*bytes.Buffer, *bytes.Buffer) {
    outBytes := new(bytes.Buffer)
    errBytes := new(bytes.Buffer)

    var outWriter io.Writer
    var errWriter io.Writer

    if stdout {
        outWriter = io.MultiWriter(os.Stdout, outBytes)
        errWriter = io.MultiWriter(os.Stderr, errBytes)
    } else {
        outWriter = io.Writer(outBytes)
        errWriter = io.Writer(errBytes)
    }

    cmd.Stdin = os.Stdin
    cmd.Stdout = outWriter
    cmd.Stderr = errWriter

    return outBytes, errBytes
}

func debugLogging(program string, args []string) {
    cm := append([]string{"[ exec ]", program})
    for _, arg := range args {
        if strings.Contains(arg, " ") {
            arg = "\"" + arg + "\""
        }
        cm = append(cm, arg)
    }
    clog.Debug(cm...)
}
