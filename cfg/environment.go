package cfg

import (
    "os"
    "reflect"
    "regexp"
    "strings"
)

func GetEnvironmentOrDefault(key string, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func EnvironmentTemplate(input []byte) []byte {
    config := string(input)

    re := regexp.MustCompile(`[{]{2}([^{|}]*)[}]{2}`)
    matches := re.FindAllStringSubmatch(config, -1)

    for _, match := range matches {
        text := strings.TrimSpace(match[1])
        if value, found := os.LookupEnv(text); found {
            config = strings.ReplaceAll(config, match[0], value)
        } else {
            config = strings.ReplaceAll(config, match[0], text)
        }
    }

    return []byte(config)
}

func SubstituteEnvironmentProperty(value string) string {
    if strings.HasPrefix(value, "env[") && strings.HasSuffix(value, "]") {
        substring := value[4 : len(value)-1]
        env := os.Getenv(substring)
        return env
    }
    return value
}

func NormalizeConfig(x interface{}) reflect.Value {
    values := reflect.ValueOf(x).Elem()

    newValue := reflect.New(values.Elem().Type()).Elem()
    newValue.Set(values.Elem())

    numField := newValue.NumField()
    types := newValue.Type()

    for i := 0; i < numField; i++ {
        if types.Field(i).Type.Kind() == reflect.Struct {

            newField := newValue.FieldByName(types.Field(i).Name)

            var b interface{}
            b = newField.Interface()

            y := NormalizeConfig(&b)
            newField.Set(y)
        } else if types.Field(i).Type.Kind() == reflect.String {

            newField := newValue.FieldByName(types.Field(i).Name)
            newField.SetString(SubstituteEnvironmentProperty(newField.String()))
        }
    }

    return newValue
}
