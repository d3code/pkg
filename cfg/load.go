package cfg

import (
    "github.com/d3code/zlog"
    "gopkg.in/yaml.v3"
    "os"
    "path"
)

func LoadConfiguration(name string, config interface{}) {
    environment := GetEnvironmentOrDefault("ENVIRONMENT", "local")
    configLocation := GetEnvironmentOrDefault("CONFIG_LOCATION", "config")

    var configFilename string
    if name != "" {
        configFilename = environment + "_" + name + ".yaml"
    } else {
        configFilename = environment + ".yaml"
    }

    configPath := path.Join(configLocation, configFilename)
    zlog.Log.Debugf("Database configuration [ %s ]", configPath)

    configFile, err := os.ReadFile(configPath)
    if err != nil {
        zlog.Log.Fatal(err)
    }

    configFile = EnvironmentTemplate(configFile)
    err = yaml.Unmarshal(configFile, config)
    if err != nil {
        zlog.Log.Fatal(err)
    }
}
