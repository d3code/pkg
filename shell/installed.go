package shell

func Installed(binary string) bool {
    _, err := RunCmdE(".", false, "which", binary)
    if err != nil {
        return false
    }
    return true
}
