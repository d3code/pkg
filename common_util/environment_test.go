package common_util

import (
    "fmt"
    "os"
    "testing"
)

func Test_SubstituteEnvironmentProperty(t *testing.T) {
    err := os.Setenv("test", "config")
    if err != nil {
        t.Error(err)
    }

    var tests = []struct {
        input string
        want  string
    }{
        {"env[test]", "config"},
        {"test", "test"},
        {"env[environment]", ""},
        {"", ""},
    }

    for _, test := range tests {
        name := fmt.Sprintf("%s", test.input)
        t.Run(name, func(t *testing.T) {
            ans := SubstituteEnvironmentProperty(test.input)
            if ans != test.want {
                t.Errorf("got %s, want %s", ans, test.want)
            }
        })
    }
}
