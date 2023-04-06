package cfg

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "io/ioutil"
    "sync"
)

var (
    configDatabaseEnvironment     DatabaseConfig
    onceConfigDatabaseEnvironment sync.Once
)

type DatabaseConfig struct {
    ConnectionType string `yaml:"connection_type"` // connector, tcp, unix
    ConnectionName string `yaml:"connection_name"`
    Host           string `yaml:"host"`
    Port           string `yaml:"port"`
    Private        bool   `yaml:"private"`
    Password       string `yaml:"password"`
    DatabaseName   string `yaml:"database_name"`
    User           string `yaml:"user"`
    CertPath       string `yaml:"cert_path"`
    RootCertPath   string `yaml:"root_cert_path"`
    KeyPath        string `yaml:"key_path"`
}

func GetDatabaseConfig(databaseName string) DatabaseConfig {
    onceConfigDatabaseEnvironment.Do(func() {
        environment := GetEnvironmentOrDefault("environment", "local")
        configLocation := GetEnvironmentOrDefault("config_location", "config")

        configPath := fmt.Sprintf("%s/database_%s_%s.yaml", configLocation, databaseName, environment)
        configFile, err := ioutil.ReadFile(configPath)
        if err != nil {
            panic(err)
        }

        err = yaml.Unmarshal(configFile, &configDatabaseEnvironment)
        if err != nil {
            panic(err)
        }

        var b interface{} = configDatabaseEnvironment

        x := NormalizeConfig(&b)
        configDatabaseEnvironment = x.Interface().(DatabaseConfig)
    })
    return configDatabaseEnvironment
}
