package cfg

import (
    "os"
    "reflect"
    "strings"
)

func GetEnvironmentOrDefault(key string, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
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
