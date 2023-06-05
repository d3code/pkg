package cfg

import (
    "github.com/d3code/cryptoledger-task/config/config"
    "os"
    "testing"
)

func TestTemplate(t *testing.T) {

    err := os.Setenv("DATABASE_SCHEMA", "cryptoledger")
    if err != nil {
        t.Errorf("Setenv() = %v", err)
    }

    tests := []struct {
        name    string
        message string
        want    string
    }{
        {name: "replace", message: "{{message}}", want: "message"},
        {name: "replace2", message: config.testConfig, want: "message"},
    }
    for _, tt := range tests {

        t.Run(tt.name, func(t *testing.T) {
            if got := EnvironmentTemplate([]byte(tt.message)); string(got) != tt.want {
                t.Errorf("Template() = %v, want %v", string(got), tt.want)
            }
        })
    }
}
