package zlog

import (
    "fmt"
    "os"
    "testing"
)

func Test_LogInfo(t *testing.T) {
    t.Run("Log", func(t *testing.T) {
        fmt.Println()

        os.Setenv("environment", "production")

        Log.Info("Test")
        Log.Debug("Test")
        Log.Warn("Test")
        Log.Error("Test")

        fmt.Println()
    })
}
