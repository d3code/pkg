package configuration

import (
    "fmt"
    "github.com/d3code/pkg/common_util"
    "github.com/d3code/pkg/log"
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

        configLocation := common_util.GetEnvironmentOrDefault("config_location", "config")
        configPath := fmt.Sprintf("%s/%s.yaml", configLocation, "_global")

        configFile, err := os.ReadFile(configPath)
        if err != nil {
            log.Log.Error(err)
            return
        }

        err = yaml.Unmarshal(configFile, &configGlobal)
        if err != nil {
            log.Log.Error(err)
            return
        }

        var b interface{} = configGlobal

        normalizeConfig := common_util.NormalizeConfig(&b)
        configGlobal = normalizeConfig.Interface().(GlobalConfig)

        for index, value := range configGlobal.Main {
            configGlobal.Main[index] = common_util.SubstituteEnvironmentProperty(value)
        }
    })

    return configGlobal
}
