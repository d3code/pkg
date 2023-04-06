package cfg

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "os"
    "sync"
)

var (
    configGlobal     GlobalConfig
    onceGlobalConfig sync.Once
)

type GlobalConfig struct {
    ApplicationName string            `yaml:"application_name"`
    Main            map[string]string `yaml:"main"`
}

func GetGlobalConfig() GlobalConfig {
    onceGlobalConfig.Do(func() {

        configLocation := GetEnvironmentOrDefault("config_location", "config")
        configPath := fmt.Sprintf("%s/%s.yaml", configLocation, "_global")

        configFile, err := os.ReadFile(configPath)
        if err != nil {
            panic(err)
            return
        }

        err = yaml.Unmarshal(configFile, &configGlobal)
        if err != nil {
            return
        }

        var b interface{} = configGlobal

        normalizeConfig := NormalizeConfig(&b)
        configGlobal = normalizeConfig.Interface().(GlobalConfig)

        for index, value := range configGlobal.Main {
            configGlobal.Main[index] = SubstituteEnvironmentProperty(value)
        }
    })

    return configGlobal
}
