package cfg

import (
    "gopkg.in/yaml.v3"
    "os"
    "path"
)

func LoadConfiguration(name string, config interface{}) error {
    environment := GetEnvironmentOrDefault("ENVIRONMENT", "local")
    configLocation := GetEnvironmentOrDefault("CONFIG_LOCATION", "config")

    var configFilename string
    if name != "" {
        configFilename = environment + "_" + name + ".yaml"
    } else {
        configFilename = environment + ".yaml"
    }

    configPath := path.Join(configLocation, configFilename)
    configFile, err := os.ReadFile(configPath)
    if err != nil {
        return err
    }

    configFile = EnvironmentTemplate(configFile)
    err = yaml.Unmarshal(configFile, config)
    return err
}
