package log

import (
    "testing"
)

func Test_LogInfo(t *testing.T) {
    t.Run("Log", func(t *testing.T) {
        Log.Info("Test")
    })
}
