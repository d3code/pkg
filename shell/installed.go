package shell

func Installed(binary string) bool {
    _, err := RunE("which", binary)
    if err != nil {
        return false
    }
    return true
}
